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
	ch := make(chan int, 10)

	go printMultiples(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
