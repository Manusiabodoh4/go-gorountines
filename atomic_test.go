package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {

	group := sync.WaitGroup{}
	var counter int64 = 0

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println(counter)

}
