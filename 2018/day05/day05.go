package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

	fmt.Println("partA:", len(reduce(bytes)) - 1)

	trimmed := []byte{}

	min := len(bytes)

	for i := 'a'; i <= 'z'; i++ {
		trimmed = trim(bytes, byte(i))
		length := len(reduce(trimmed)) - 1
		if length < min {
			min = length
		}
	}
	fmt.Println("partB:", min)
}

func trim(bytes []byte, remove byte) []byte {
	result := []byte{}
	for _, b := range bytes {
		if b == remove || b == remove - 32 {
			continue
		}
		result = append(result, b)
	}
	return result
}

func reduce(bytes []byte) []byte {
	result := make([]byte, len(bytes))
	copy(result, bytes)

	lastByte := byte(0)

	for {
	top:
		lastByte = byte(0)
		for i, b := range result {
			if lastByte - b == 32 || b - lastByte == 32 {
				result = append(result[0:i-1], result[i+1:]...)
				goto top
			}
			lastByte = b
		}
		return result
	}
}
