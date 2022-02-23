package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var group = sync.WaitGroup{}
var cond = sync.NewCond(&sync.Mutex{})

func WaitWithCondition(value int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {

	for i := 0; i < 10; i++ {
		go WaitWithCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		time.Sleep(5 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()

}
