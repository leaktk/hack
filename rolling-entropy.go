package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filepath1> [filepath2 ...]\n", os.Args[0])
		os.Exit(1)
	}

	const chunkSize = 24
	const stride = 8

	buf := make([]byte, 4096)

	for _, filepath := range os.Args[1:] {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening file %s: %v\n", filepath, err)
			continue
		}

		reader := bufio.NewReader(file)
		results := make([][3]interface{}, 0)
		for n := 0; err == nil; {
			n, err = reader.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Fprintf(os.Stderr, "error reading initial chunk from %s: %v\n", filepath, err)
				file.Close()
				continue
			}

			for start := 0; start < n; start += stride {
				end := min(start+chunkSize, n)
				entropy := shannonEntropy(buf[start:end])
				results = append(results, [3]interface{}{start, end, entropy})
			}
		}

		file.Close()

		// After processing all chunks for a file, construct the JSON object.
		result := map[string]interface{}{
			"path":      filepath,
			"entropies": results,
		}

		jsonBytes, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshalling JSON for %s: %v\n", filepath, err)
			continue
		}

		fmt.Println(string(jsonBytes))
	}
}
