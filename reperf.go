package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Usage:\n\treperf <iterations> <regex> <file>")
	}

	iterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile(os.Args[2])
	input, err := os.ReadFile(os.Args[3])

	if err != nil {
		log.Fatal(err)
	}

	times := make([]time.Duration, iterations)
	var matches [][]int

	for i := 0; i < iterations; i++ {
		start := time.Now()
		matches = pattern.FindAllIndex(input, -1)
		times[i] = time.Since(start)
	}

	var totalTime time.Duration
	for _, time := range times {
		totalTime += time
	}
	avgTime := totalTime / time.Duration(iterations)

	fmt.Printf("matches=%v avgtime=%v\n", len(matches), avgTime)
}
