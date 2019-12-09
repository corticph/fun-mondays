package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timed out!")
				return
			default:
				time.Sleep(time.Millisecond * 250)
				doSomething(ctx)
				fmt.Println("doing something!")
			}
		}

	}()

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	doSomething(ctx)

	time.Sleep(time.Second * 5)

}
