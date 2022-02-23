package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Reza Andriansyah"
	fmt.Println("Berhasil mengirim data ke channel")
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Reza Andriansyah"
	fmt.Println("Berhasil mengirim data ke channel")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestCreateChannelDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}

func TestCreateChannelSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestCreateChannelRange(t *testing.T) {
	channel := make(chan string)

	go func() {

		for i := 0; i < 10; i++ {
			channel <- "Reza Andriansyah" + strconv.Itoa(i)
		}

		close(channel)

	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

}

func TestCreateChannelBuffered(t *testing.T) {
	channel := make(chan string, 5)
	defer close(channel)

	channel <- "Reza Andriansyah"
	channel <- "Amelia Rosa"
	channel <- "Rizky Prawira"

	fmt.Println(len(channel))
	fmt.Println(cap(channel))

	fmt.Println("Selesai")

}

func TestCreateChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestCreateChannelParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Reza Andriansyah"
		fmt.Println("Berhasil mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}
