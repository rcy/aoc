package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"strconv"
)

var input = flag.String("input", "./input", "The input to the problem.")

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")[0]

	s := strings.Split(lines, ",")
	program := []int{}
	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		program = append(program, j)
	}
	
	// part 1
	run(program, 10)
}

func output(value int) {
	fmt.Printf("output %d\n", value)
}

func run(program []int, input int) int {
	ip := 0

	p := make([]int, len(program))
	copy(p, program)

	// p[1] = noun
	// p[2] = verb

	for p[ip] != 99 {
		fmt.Println("opcode", p[ip])
		if (p[ip] == 1) {
			// addition
			p[p[ip + 3]] = p[p[ip + 1]] + p[p[ip + 2]]
			ip += 4
		} else if (p[ip] == 2) {
			// multiplication
			p[p[ip + 3]] = p[p[ip + 1]] * p[p[ip + 2]]
			ip += 4
		} else if (p[ip] == 3) {
			fmt.Println("storeinput", p[ip + 1], input)
			p[p[ip + 1]] = input
			ip += 2
		} else if (p[ip] == 4) {
			output(p[p[ip + 1]])
			ip += 2
		} else {
			fmt.Println(p[ip])
			panic("unknown opcode")
		}
		fmt.Println(p)
	}

	return p[0]
}
