package qdb

import (
	"context"
	"github.com/pg-sharding/spqr/pkg/spqrlog"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// todo delete it

// LeaseWatcher watches for lease expiration events in etcd
type LeaseWatcher struct {
	client *clientv3.Client
	done   chan struct{}
}

// NewLeaseWatcher creates a new lease watcher instance
// func NewLeaseWatcher(addr string, coord coordinator.Coordinator) (*LeaseWatcher, error) {
func NewLeaseWatcher(addr string) (*LeaseWatcher, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{addr},
		DialOptions: []grpc.DialOption{ // TODO remove WithInsecure
			grpc.WithInsecure(), //nolint:all
		},
	})
	if err != nil {
		return nil, err
	}

	return &LeaseWatcher{
		client: client,
		done:   make(chan struct{}),
	}, nil
}

// Start begins watching for lease expiration events
func (w *LeaseWatcher) Start(ctx context.Context) {
	go func() {
		watcher := clientv3.NewWatcher(w.client)
		defer watcher.Close()

		// use WithPrevKV to get the value, after delete value is null
		watchChan := watcher.Watch(ctx, TwoPhaseCommitsLease, clientv3.WithPrefix(), clientv3.WithPrevKV())

		for {
			select {
			case <-w.done:
				return
			case <-ctx.Done():
				return
			case watchResp := <-watchChan:
				if watchResp.Err() != nil {
					spqrlog.Zero.Error().
						Err(watchResp.Err()).
						Msg("Failed to watch")
					continue
				}

				for _, event := range watchResp.Events {
					if event.Type == clientv3.EventTypeDelete {
						// every event need
						spqrlog.Zero.Debug().
							Str("key", string(event.Kv.Key)).
							Str("value", string(event.PrevKv.Value)).
							Msg("Lease expired")
					}
				}
			}
		}
	}()
}

// Stop stops the lease watcher
func (w *LeaseWatcher) Stop() {
	close(w.done)
}
