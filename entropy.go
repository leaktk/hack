package main

import (
	"fmt"
	"math"
	"os"
)

func shannonEntropy(data string) (entropy float64) {
	if data == "" {
		return 0
	}

	charCounts := make(map[rune]int)
	for _, char := range data {
		charCounts[char]++
	}

	invLength := 1.0 / float64(len(data))
	for _, count := range charCounts {
		freq := float64(count) * invLength
		entropy -= freq * math.Log2(freq)
	}

	return entropy
}

func main() {
	for _, data := range os.Args[1:] {
		fmt.Printf("%f %s\n", shannonEntropy(data), data)
	}
}
