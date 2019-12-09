package main

import (
	"context"
	"fmt"
	"time"
)

func exampleTimeout(ctx context.Context) {

	select {

	case <-time.After(time.Second * 1):
		fmt.Println("did some stuff")
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go exampleTimeout(ctx)
	cancel()
	time.Sleep(time.Second * 5)

}
