// -*- compile-command: "go run day07.go ./intcode.go" -*-

package main

import (
	"fmt"
	"flag"
	"strings"
	"io/ioutil"
	"log"
	"strconv"
)

var input = flag.String("input", "./input", "The input to the problem.")

func load(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	lines := strings.Split(str, "\n")

	program := Program{}

	for _, i := range strings.Split(lines[0], ",") {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		program = append(program, j)
	}

	return program
}

type Sequence []int

func runSequence(program Program, sequence Sequence) int {
	out := 0
	for i := 0; i < 5; i++ {
		out = run(program, []int{sequence[i], out})
	}
	return out
}

func swap(arr Sequence, i int, j int) Sequence {
	arr2 := make(Sequence, len(arr))
	copy(arr2, arr)
	arr2[i] = arr[j]
	arr2[j] = arr[i]
	return arr2
}

func generate(k int, arr Sequence, acc []Sequence) ([]Sequence, Sequence) {
	if k == 1 {
		acc = append(acc, arr)
	} else {
		for i := 0; i < k; i++ {
			acc, arr = generate(k - 1, arr, acc)
			if k % 2 == 0 {
				arr = swap(arr, i, k-1)
			} else {
				arr = swap(arr, 0, k-1)
			}
		}
	}
	return acc, arr
}

func findmax(program Program) int {
	sequences, _ := generate(5, []int{0,1,2,3,4}, []Sequence{})
	var out int
	var max int
	for _, s := range sequences {
		out = runSequence(program, s)
		if out > max {
			max = out
		}
	}
	return max
}

func main() {
	flag.Parse()

	program := load(*input)
	fmt.Println("part 1", findmax(program))
}
