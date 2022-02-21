package main

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestCreateGoro(t *testing.T) {

	go HelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)

}

func TestCreateManyGoro(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(1 * time.Second)
}
