package main

import (
	"strconv"
	"fmt"
	"bufio"
	"os"
	"log"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var freq int64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		num, _ := strconv.ParseInt(text, 10, 64)
		freq += num
		fmt.Println(num, freq)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
