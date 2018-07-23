package cache

import "testing"

func TestNewLRU(t *testing.T) {
	Expected := "capacity must be gt 0"
	var actual error

	_, actual = NewLRU(0)

	if actual.Error() != Expected {
		t.Failed()
	}

	_, actual = NewLRU(-1)

	if actual.Error() != Expected {
		t.Failed()
	}

	_, actual = NewLRU(1)

	if actual != nil {
		t.Failed()
	}
}

func TestLRU(t *testing.T) {
	lru, _ := NewLRU(2)

	lru.Put(1, 1)
	lru.Put(2, 2)

	if lru.Get(1) != 1 {
		t.Error("1 expected")
	}

	lru.Put(3, 3)

	if lru.Get(2) != -1 {
		t.Error("-1 expected")
	}

	lru.Put(4, 4)

	if lru.Get(3) != 3 {
		t.Error("3 expected")
	}

	if lru.Get(4) != 4 {
		t.Error("4 expected")
	}
}

func TestLRU_StorageBug(t *testing.T) {
	lru, _ := NewLRU(2)

	lru.Put(1, 111)
	lru.Put(2, 222)
	lru.Put(3, 333)

	expected := 333
	actual := lru.Get(3)

	if actual != expected {
		t.Error("333 expected")
	}
}

// При попытке конкурентной записи в map без синхронизации поймаем фатал
func TestLRU_ConcurrentWrite(t *testing.T) {
	goCount := 100
	size := 10

	lru, _ := NewLRU(size)

	put := func(val int) {
		lru.Put(val, val)
	}

	for i := 0; i < goCount; i++ {
		go put(i)
	}
}
