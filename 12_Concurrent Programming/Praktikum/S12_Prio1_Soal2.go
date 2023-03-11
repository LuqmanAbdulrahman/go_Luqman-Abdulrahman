package main

import (
	"fmt"
	"time"
)

func printMultiples(ch chan<- int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			ch <- i
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int)

	go printMultiples(ch)

	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout!")
			return
		}
	}
}
