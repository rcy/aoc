package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
//	"math"
)

var input = flag.String("input", "../inputs/day05.txt", "The input to the problem.")

func bytesFromInput(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func main() {
	flag.Parse()

	bytes := bytesFromInput(*input)

	bytes = reduce(bytes)

	fmt.Printf("%d %s", len(bytes) - 1, bytes)
}

func reduce(bytes []byte) []byte {
	lastByte := byte(0)

	for {
	top:
		lastByte = byte(0)
		for i, b := range bytes {
			if lastByte - b == 32 || b - lastByte == 32 {
				bytes = append(bytes[0:i-1], bytes[i+1:]...)
				goto top
			}
			lastByte = b
		}
		return bytes
	}
}
