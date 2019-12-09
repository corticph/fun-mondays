package main

import (
	"fmt"
	"time"
)

func main() {

	gen := func() <-chan int {
		dst := make(chan int, 100)
		n := 1
		go func() {
			for {
				fmt.Println(n)
				n++
				dst <- n
			}
		}()
		return dst
	}

	for n := range gen() {
		if n == 5 {
			break
		}
	}
	time.Sleep(time.Second * 1)
}
