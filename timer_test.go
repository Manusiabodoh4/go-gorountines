package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	time := <-timer.C
	fmt.Println(time)
}
func TestAfterTimer(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())
	time := <-channel
	fmt.Println(time)
}

func TestAfterFuncTimer(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()
	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	for tick := range channel {
		fmt.Println(tick)
	}
}
