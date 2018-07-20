package cache

import (
	"container/list"
	"errors"
)

const elementNotFound = -1

func NewLRU(capacity int) (LRU, error) {
	if capacity <= 0 {
		return LRU{}, errors.New("capacity must be gt 0")
	}

	return LRU{
		capacity: capacity,
		list:     list.New(),
		elements: make(map[int]*list.Element),
	}, nil
}

type LRU struct {
	capacity int
	list     *list.List
	elements map[int]*list.Element
}

func (l *LRU) Put(key int, value int) {
	element := l.elements[key]

	if element == nil {
		element = l.list.PushFront(key)
		l.elements[key] = element
	}

	if l.list.Len() > l.capacity {
		l.removeOldest()
	}
}

func (l *LRU) Get(key int) int {
	if element, ok := l.elements[key]; ok {
		l.list.MoveToFront(element)

		return element.Value.(int)
	}

	return elementNotFound
}

func (l *LRU) removeOldest() {
	back := l.list.Back()

	if back != nil {
		l.list.Remove(back)
		delete(l.elements, back.Value.(int))
	}
}
