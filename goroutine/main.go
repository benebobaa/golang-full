package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var balance int
var balanceSafe atomic.Int32

func main() {

	balance = 10000
	balanceSafe.Add(10000)

	fmt.Println("init balance: ", balance)

	// blast transfer with goroutine
	for i := 1; i <= 1000; i++ {
		go transfer(10)
		go transferSafe(10)
	}

	// the current balance after transfer
	// should be 0, because 1.000 * 10 = 10.000
	// then 10.000 - 10.000 = 0
	time.Sleep(1 * time.Second)
	fmt.Println("current balance not safe: ", balance)
	fmt.Println("current balance safe: ", balanceSafe.Load())
}

// this func is not thread safe
func transfer(value int) {
	balance -= value
}

// this func thread safe
func transferSafe(value int32) {
	balanceSafe.Add(-value)
}
