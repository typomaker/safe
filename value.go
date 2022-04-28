package safe

import (
	"sync"
)

type Value[T any] struct {
	m sync.RWMutex
	v T
}

func (it *Value[T]) Set(v T) {
	it.m.Lock()
	defer it.m.Unlock()
	it.v = v
}
func (it *Value[T]) Get() T {
	it.m.RLock()
	defer it.m.RUnlock()
	return it.v
}
