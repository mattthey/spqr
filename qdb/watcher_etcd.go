package qdb

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// EtcdWatcher реализация WatcherInterface для etcd
type EtcdWatcher struct {
	client  *clientv3.Client
	watcher clientv3.Watcher
}

func NewEtcdWatcher(client *clientv3.Client) *EtcdWatcher {
	return &EtcdWatcher{
		client:  client,
		watcher: clientv3.NewWatcher(client),
	}
}

func (w *EtcdWatcher) Watch(ctx context.Context, key string) WatchChan {
	etcdChan := w.watcher.Watch(ctx, key, clientv3.WithPrefix(), clientv3.WithPrevKV())
	ch := make(chan WatchResponse)

	go func() {
		//defer close(ch)
		for watchResp := range etcdChan {
			if watchResp.Err() != nil {
				ch <- WatchResponse{Err: watchResp.Err()}
				continue
			}

			events := make([]*WatchEvent, 0, len(watchResp.Events))
			for _, e := range watchResp.Events {
				event := &WatchEvent{
					Key:   string(e.Kv.Key),
					Value: string(e.Kv.Value),
				}

				if e.Type == clientv3.EventTypeDelete {
					event.Type = EventTypeDelete
				} else {
					event.Type = EventTypePut
				}

				if e.PrevKv != nil {
					event.PrevVal = string(e.PrevKv.Value)
				}

				events = append(events, event)
			}

			ch <- WatchResponse{Events: events}
		}
	}()

	return ch
}

func (w *EtcdWatcher) Close() error {
	return w.watcher.Close()
}
