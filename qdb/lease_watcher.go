package qdb

import (
	"context"
	"github.com/pg-sharding/spqr/pkg/spqrlog"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// LeaseWatcher watches for lease expiration events in etcd
type LeaseWatcher struct {
	client *clientv3.Client
	done   chan struct{}
}

// NewLeaseWatcher creates a new lease watcher instance
func NewLeaseWatcher(client *clientv3.Client) *LeaseWatcher {
	return &LeaseWatcher{
		client: client,
		done:   make(chan struct{}),
	}
}

// Start begins watching for lease expiration events
func (w *LeaseWatcher) Start(ctx context.Context) {
	go func() {
		watcher := clientv3.NewWatcher(w.client)
		defer watcher.Close()

		// use WithPrevKV to get the value, after delete value is null
		watchChan := watcher.Watch(ctx, twoPhaseCommitsLease, clientv3.WithPrefix(), clientv3.WithPrevKV())

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
