package main

import (
	"fmt"
	"sync"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	freqs := make(map[rune]int)
	var wg sync.WaitGroup
	var lock sync.Mutex

	ch := make(chan rune)

	// Memunculkan goroutine untuk menghitung frekuensi setiap karakter
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range ch {
				lock.Lock()
				freqs[c]++
				lock.Unlock()
			}
		}()
	}

	// Kirim setiap karakter ke channel untuk dihitung
	for _, c := range text {
		ch <- c
	}
	close(ch)

	// Tunggu hingga semua goroutine selesai
	wg.Wait()

	// Frekuensi cetak setiap karakter
	for c, freq := range freqs {
		fmt.Printf("%c: %d\n", c, freq)
	}
}
