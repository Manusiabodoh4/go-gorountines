package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "NEW"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Reza")
	pool.Put("Andriansyah")
	pool.Put("Amelia")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	group.Wait()

	fmt.Println("Complete")

}
