package twopc

import (
	"context"
	"fmt"
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

const (
	COMMIT_STRATEGY_BEST_EFFORT = "best-effort"
	/* same af above */
	COMMIT_STRATEGY_1PC = "1pc"
	COMMIT_STRATEGY_2PC = "2pc"
)

func ExecuteTwoPhaseCommit(clid uint, s server.Server, qdb qdb.QDB) error {
	/*
	* go along first phase
	 */

	txid := uuidv7.New().String()

	// 1. create tx and get lease

	for _, dsh := range s.Datashards() {
		st, err := shard.DeployTxOnShard(dsh, &pgproto3.Query{
			String: fmt.Sprintf(`PREPARE TRANSACTION '%s'`, txid),
		}, txstatus.TXIDLE)

		// если приложение аварийно завершится после того как частично транзакции будут препарены, тогда
		// строки частично зависнут и мы не сможем с ними работать. При поднятии приложения необходимо откатить транзакции
		//err = qdb.Create2phaseCommit(context.Background(), txid, []string{dsh.ShardKeyName()})
		err = create2PCCommit(qdb, txid, dsh.ShardKeyName())

		if err != nil {
			/* assert st == txtstatus.TXERR? */
			s.SetTxStatus(txstatus.TXStatus(txstatus.TXERR))
			return err
		}

		s.SetTxStatus(txstatus.TXStatus(st))
	}

	// update tx set status commit

	err := qdb.Create2phaseCommit(context.Background(), txid, []string{"prepared_all"})
	if err != nil {
		s.SetTxStatus(txstatus.TXStatus(txstatus.TXERR))
		return err
	}

	spqrlog.Zero.Info().Uint("client", clid).Str("txid", txid).Msg("first phase succeeded")

	for _, dsh := range s.Datashards() {
		st, err := shard.DeployTxOnShard(dsh, &pgproto3.Query{
			String: fmt.Sprintf(`COMMIT PREPARED '%s'`, txid),
		}, txstatus.TXIDLE)

		// если уже выполнили commit prepared, то транзакция должна выполниться окончательно на всех шардах

		//shardsArray := []string{dsh.ShardKeyName()}
		//err = qdb.GetAll2phaseCommits(context.Background(), txid, shardsArray)
		//if err != nil {
		//	s.SetTxStatus(txstatus.TXStatus(txstatus.TXERR))
		//	return err
		//}

		if err != nil {
			/* assert st == txtstatus.TXERR? */
			s.SetTxStatus(txstatus.TXStatus(txstatus.TXERR))
			return err
		}

		s.SetTxStatus(txstatus.TXStatus(st))
	}

	return nil
}

func create2PCCommit(qdb qdb.QDB, txid string, shard string) error {
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

	err = adapter.Create2PhaseCommit(txid, shard)
	if err != nil {
		return err
	}

	return nil
}
