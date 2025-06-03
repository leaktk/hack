package main

import (
	"github.com/zricethezav/gitleaks/v8/detect/codec"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	maxDecodeDepth := 8
	var err error

	if len(os.Args) == 2 || len(os.Args) > 3 {
		log.Fatal("Usage:\n\tomnidecode [--depth n] < file")
	}

	if len(os.Args) == 3 && os.Args[2] == "--depth" {
		maxDecodeDepth, err = strconv.Atoi(os.Args[3])
	}

	if err != nil {
		log.Fatal(err)
	}

	rawData, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	currentDecodeDepth := 0
	data := string(rawData)
	decoder := codec.NewDecoder()
	encodedSegments := []*codec.EncodedSegment{}

	for {
		// increment the depth by 1 as we start our decoding pass
		currentDecodeDepth++

		// stop the loop if we've hit our max decoding depth
		if currentDecodeDepth > maxDecodeDepth {
			break
		}

		data, encodedSegments = decoder.Decode(data, encodedSegments)

		if len(encodedSegments) == 0 {
			break
		}
	}

	fmt.Print(data)
}
