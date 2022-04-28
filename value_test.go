package safe

import (
	"sync"
	"testing"
)

func TestValueRace(t *testing.T) {
	var wg sync.WaitGroup
	var val Value[int]
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				val.Set(i)
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				val.Get()
			}
		}()
	}
	wg.Wait()
}
