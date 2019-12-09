package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doSomething(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timed out!")
				return
			default:
				time.Sleep(time.Millisecond * 500)
				fmt.Println("doing something!")
			}
		}

	}()

}

func doSomethingElse(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timed out!")
				wg.Done()
				return
			default:
				time.Sleep(time.Millisecond * 500)
				fmt.Println("doing something else!")
			}
		}

	}()

}

func main() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	doSomething(ctx)
	doSomethingElse(ctx)
	wg.Wait()

}
