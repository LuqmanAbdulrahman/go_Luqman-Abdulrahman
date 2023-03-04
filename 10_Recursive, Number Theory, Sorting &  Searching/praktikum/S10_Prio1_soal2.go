package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(item []string) []pair {
	// hitung jumlah kemunculan setiap item
	counts := make(map[string]int)
	for _, v := range item {
		counts[v]++
	}

	// konversi hasil hitung menjadi slice dari pair
	pairs := make([]pair, 0, len(counts))
	for k, v := range counts {
		pairs = append(pairs, pair{k, v})
	}

	// urutkan slice pair berdasarkan jumlah kemunculan (count) secara menurun
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "js", "js"}))
	// [{js 4} {ruby 1} {golang 1}]
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// [{A 4} {B 3} {D 2} {C 1}]
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// [{football 1} {basketball 1} {tenis 1}]
}
