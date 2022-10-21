package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := l.items[key]; ok {
		// передвинем вперед так как в этом смысл кэша
		// часто используемое располагать сверху
		l.queue.MoveToFront(item)
		(item.Value.(*cacheItem)).value = value
		return true
	}

	// создадим новый ключ-значение
	item := &cacheItem{
		key:   key,
		value: value,
	}

	// В связи с таким тестом:
	// например: n = 3, добавили 4 элемента - 1й из кэша вытолкнулся
	// переделаем
	queueLen := l.queue.Len() >= l.capacity
	if queueLen {
		l.queue.Remove(l.queue.Back())
		delete(l.items, (l.queue.Back().Value.(*cacheItem)).key)
	}
	l.items[key] = l.queue.PushFront(item)

	return queueLen
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := l.items[key]; ok {
		// тут просто возьмем  кусок из Set
		l.queue.MoveToFront(item)
		return (item.Value.(*cacheItem)).value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	// просто создадим новый список так как ресивер у нас по указателю
	l.queue = NewList()
	l.items = make(map[Key]*ListItem)
}
