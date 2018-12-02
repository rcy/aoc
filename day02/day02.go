package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
)

var input = flag.String("inputFile", "input.txt", "The input to the problem.")

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	lines := strings.Split(str, "\n")

	total2 := 0
	total3 := 0

	for i, line := range lines {
		if line == "" {
			break
		}
		if hasTwoChars(line) {
			total2++
		}
		if hasThreeChars(line) {
			total3++
		}
		for _, line2 := range lines[i+1:len(lines)-1] {
			match := nearMatch(line, line2)
			if (len(match) >= 1) {
				fmt.Printf("partB: %s\n", match)
			}
		}
	}

	fmt.Printf("partA: %d * %d = %d\n", total2, total3, total2 * total3)
}

func nearMatch(s1 string, s2 string) string {
	differences := 0
	similar := ""
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			differences++
			if differences > 1 {
				return ""
			}
		} else {
			similar = similar + string(s1[i])
		}
	}
	return similar
}

func hasTwoChars(str string) bool {
	for _, ch := range str {
		if strings.Count(str, string(ch)) == 2 {
			return true
		}
	}
	return false
}

func hasThreeChars(str string) bool {
	for _, ch := range str {
		if strings.Count(str, string(ch)) == 3 {
			return true
		}
	}
	return false
}
