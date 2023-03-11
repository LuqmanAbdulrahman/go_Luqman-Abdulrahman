package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex = &sync.Mutex{}

func main() {
	for i := 0; i < 5; i++ {
		go incrementCounter(i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("Counter value:", counter)
}

func incrementCounter(id int) {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		counter++
		fmt.Printf("Goroutine %d incremented counter: %d\n", id, counter)
		mutex.Unlock()
		time.Sleep(time.Millisecond * 500)
	}
}
