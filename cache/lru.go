package cache

import (
	"container/list"
	"errors"
	"sync"
)

const elementNotFound = -1

func NewLRU(capacity int) (LRU, error) {
	if capacity <= 0 {
		return LRU{}, errors.New("capacity must be gt 0")
	}

	return LRU{
		capacity: capacity,
		list:     list.New(),
		items:    make(map[int]*list.Element),
		mu:       sync.Mutex{},
	}, nil
}

type Item struct {
	key   int
	value int
}

type LRU struct {
	mu       sync.Mutex
	capacity int
	list     *list.List
	items    map[int]*list.Element
}

func (l *LRU) Put(key int, value int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	el := l.items[key]

	if el == nil {
		item := &Item{key, value}
		el = l.list.PushFront(item)
		l.items[key] = el
	}

	if l.list.Len() > l.capacity {
		l.removeOldest()
	}
}

func (l *LRU) Get(key int) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	if element, ok := l.items[key]; ok {
		l.list.MoveToFront(element)
		item := element.Value.(*Item)

		return item.value
	}

	return elementNotFound
}

func (l *LRU) removeOldest() {
	back := l.list.Back()

	if back != nil {
		elem := l.list.Remove(back)
		item := elem.(*Item)
		delete(l.items, item.key)
	}
}
