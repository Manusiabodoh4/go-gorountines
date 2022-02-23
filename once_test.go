package main

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	// var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			OnlyOnce()
		}()
	}

	group.Wait()
	fmt.Println(counter)

}
