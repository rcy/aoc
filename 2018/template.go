package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
)

var input = flag.String("input", "../inputs/day03.txt", "The input to the problem.")

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

func main() {
	flag.Parse()

	lines := linesFromInput(*input)

	for i, line := range lines {
		if len(line) > 0 {
			fmt.Printf("%d %s\n", i, line)
		}
	}
}
