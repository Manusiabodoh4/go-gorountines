package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) SetBalance(amount int) {
	user.Balance = amount
}

func (user *UserBalance) AddBalance(amount int) {
	user.SetBalance(user.Balance + amount)
}

func (user *UserBalance) MinusBalance(amount int) {
	user.SetBalance(user.Balance - amount)
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("User 1 Lock ", user1.Name)
	user1.MinusBalance(amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("User 2 Lock ", user2.Name)
	user2.AddBalance(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()

}

func TestDeadLockCase(t *testing.T) {

	user1 := UserBalance{
		Name:    "Reza",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Amelia",
		Balance: 500000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, " Balance = ", user1.Balance)
	fmt.Println("User ", user2.Name, " Balance = ", user2.Balance)

}
