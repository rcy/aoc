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
	

	process(program)

	
}

func process(program []int) {
//	fmt.Printf("program: %d", program)

	pc := 0

	program[1] = 12
	program[2] = 2

	for program[pc] != 99 {
		//fmt.Printf("program = %d, pc=%d\n", program, pc)
		if (program[pc] == 1) {
			//fmt.Printf("addition %d %d", program[program[pc+1]], program[program[pc+1]])
			// addition
			program[program[pc + 3]] = program[program[pc + 1]] + program[program[pc + 2]]
		}
		if (program[pc] == 2) {
			//fmt.Printf("addition %d %d", program[program[pc+1]], program[program[pc+1]])
			// addition
			program[program[pc + 3]] = program[program[pc + 1]] * program[program[pc + 2]]
		}
		pc = pc + 4
	}
	fmt.Println(program)
}
