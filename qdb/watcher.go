package qdb

import "context"

// Watcher определяет интерфейс для отслеживания изменений
type Watcher interface {
	// Watch начинает отслеживание изменений по ключу
	Watch(ctx context.Context, key string) WatchChan
	// Close закрывает watcher
	Close() error
}

// WatchEvent представляет событие изменения
type WatchEvent struct {
	Type    EventType // тип события (создание, удаление, изменение)
	Key     string    // ключ
	Value   string    // текущее значение
	PrevVal string    // предыдущее значение
}

// EventType определяет тип события
type EventType int

const (
	EventTypePut    EventType = iota // создание/изменение
	EventTypeDelete                  // удаление
)

// WatchChan канал для получения событий
type WatchChan <-chan WatchResponse

// WatchResponse содержит события или ошибку
type WatchResponse struct {
	Events []*WatchEvent
	Err    error
}
