package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("context")

	ctx := context.Background()
	fmt.Println("background :: ", ctx)

	ctxValue := context.WithValue(ctx, "1", "satu")
	fmt.Println("value :: ", ctxValue)

	ctxCancel, _ := context.WithCancel(ctx)
	fmt.Println("cancel :: ", ctxCancel)

	ctxTimeout, _ := context.WithTimeout(ctx, 1*time.Second)
	fmt.Println("timeout :: ", ctxTimeout)

	ctxDeadline, _ := context.WithDeadline(ctx, time.Now().Add(1*time.Second))

	fmt.Println("deadline :: ", ctxDeadline)
}
