package twopc

import (
	"context"
	"errors"
	"fmt"
	"github.com/pg-sharding/spqr/pkg/models/topology"

	"github.com/jackc/pgx/v5/pgproto3"
	"github.com/pg-sharding/spqr/pkg/coord"
	"github.com/pg-sharding/spqr/pkg/shard"
	"github.com/pg-sharding/spqr/pkg/spqrlog"
	"github.com/pg-sharding/spqr/pkg/txstatus"
	"github.com/pg-sharding/spqr/qdb"
	"github.com/pg-sharding/spqr/router/server"
	"github.com/samborkent/uuidv7"
	"google.golang.org/grpc"
)

func ExecuteTwoPhaseCommit(clid uint, s server.Server, qdb qdb.QDB) error {
	/*
	* go along first phase
	 */

	txid := uuidv7.New().String()

	// 1. create tx and get lease
	// если приложение аварийно завершится после того как частично транзакции будут препарены, тогда
	// строки частично зависнут и мы не сможем с ними работать. При поднятии приложения необходимо откатить транзакции
	//err = qdb.Create2phaseCommit(context.Background(), txid, []string{dsh.ShardKeyName()})
	lease, err := create2PCCommit(qdb, txid)
	if err != nil {
		spqrlog.Zero.Error().Err(err).Msg("failed to create 2PC commit")
		return err
	}

	spqrlog.Zero.Debug().
		Str("txId", txid).
		Int64("lease", lease).
		Msg("Create 2PC commit")

	for _, dsh := range s.Datashards() {
		st, err := shard.DeployTxOnShard(dsh, &pgproto3.Query{
			String: fmt.Sprintf(`PREPARE TRANSACTION '%s'`, txid),
		}, txstatus.TXIDLE)

		if err != nil {
			/* assert st == txtstatus.TXERR? */
			s.SetTxStatus(txstatus.TXStatus(txstatus.TXERR))
			return err
		}

		s.SetTxStatus(txstatus.TXStatus(st))
	}

	spqrlog.Zero.Info().Uint("client", clid).Str("txid", txid).Msg("first phase succeeded")

	err = update2PCCommit(qdb, txid, "prepared")
	if err != nil {
		spqrlog.Zero.Error().Err(err).Msg("failed to update 2PC commit")
		return err
	}

	for _, dsh := range s.Datashards() {
		st, err := shard.DeployTxOnShard(dsh, &pgproto3.Query{
			String: fmt.Sprintf(`COMMIT PREPARED '%s'`, txid),
		}, txstatus.TXIDLE)

		if err != nil {
			/* assert st == txtstatus.TXERR? */
			s.SetTxStatus(txstatus.TXStatus(txstatus.TXERR))
			return err
		}

		s.SetTxStatus(txstatus.TXStatus(st))
	}

	err = update2PCCommit(qdb, txid, "committed")
	if err != nil {
		spqrlog.Zero.Error().Err(err).Msg("failed to update 2PC commit")
		return err
	}

	spqrlog.Zero.Info().Uint("client", clid).Str("txid", txid).Msg("second phase succeeded")

	return nil
}

func FinishTwoPhaseCommit(txid string, status string, shards []*topology.DataShard, qdb qdb.QDB) error {
	return errors.New("not implemented yet")
}

// todo mattthey extract adapter to main method
func create2PCCommit(qdb qdb.QDB, txid string) (int64, error) {
	coordinator, err := qdb.GetCoordinator(context.Background())
	if err != nil {
		return -1, err
	}
	conn, err := grpc.NewClient(coordinator, grpc.WithInsecure())
	if err != nil {
		return -1, err
	}
	defer conn.Close()
	adapter := coord.NewAdapter(conn)
	if adapter == nil {
		return -1, fmt.Errorf("failed to create coordinator adapter")
	}

	lease, err := adapter.Create2PhaseCommitWithLease(context.Background(), txid)

	return lease, err
}

func update2PCCommit(qdb qdb.QDB, txid string, status string) error {
	coordinator, err := qdb.GetCoordinator(context.Background())
	if err != nil {
		return err
	}
	conn, err := grpc.NewClient(coordinator, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	adapter := coord.NewAdapter(conn)
	if adapter == nil {
		return fmt.Errorf("failed to create coordinator adapter")
	}

	err = adapter.Update2PhaseCommit(context.Background(), txid, status)

	return err
}
