package main

import (
	"fmt"
	"sync"
	"testing"
)

var data sync.Map

func AddToMap(group *sync.WaitGroup, value int) {
	defer group.Done()
	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go AddToMap(&group, i)
	}
	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
