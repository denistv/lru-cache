package cache

import "testing"

func TestLRU(t *testing.T) {
	lru, _ := NewLRU(2)

	lru.Put(1,1)
	lru.Put(2,2)

	if lru.Get(1) != 1 {
		t.Error("1 expected")
	}

	lru.Put(3,3)

	if lru.Get(2) != -1 {
		t.Error("-1 expected")
	}

	lru.Put(4,4)

	if lru.Get(3) != 3 {
		t.Error("3 expected")
	}

	if lru.Get(4) != 4 {
		t.Error("4 expected")
	}
}

func TestNewLRULtZero(t *testing.T) {
	_, actual := NewLRU(0)

	Expected := "capacity must be gt 0"

	if actual.Error() != Expected {
		t.Error("")
	}
}
