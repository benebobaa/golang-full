package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type User struct {
	Name string
}

var balance int
var balanceSafe atomic.Int32
var balanceSafe2 atomic.Pointer[User]

func main() {
	wg := &sync.WaitGroup{}
	balance = 10000
	balanceSafe.Add(10000)

	fmt.Println("init balance: ", balance)

	// blast transfer with goroutine
	for i := 1; i <= 1000; i++ {
		wg.Add(2)
		go transfer(10, wg)
		go transferSafe(10, wg)
	}

	// the current balance after transfer
	// should be 0, because 1.000 * 10 = 10.000
	// then 10.000 - 10.000 = 0
	wg.Wait()
	fmt.Println("current balance not safe: ", balance)
	fmt.Println("current balance safe: ", balanceSafe.Load())
}

// this func is not thread safe
func transfer(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	balance -= value
}

// this func thread safe
func transferSafe(value int32, wg *sync.WaitGroup) {
	defer wg.Done()
	balanceSafe.Add(-value)
}
