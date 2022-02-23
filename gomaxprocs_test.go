package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)
	totalGoro := runtime.NumGoroutine()
	fmt.Println(totalGoro)
}
