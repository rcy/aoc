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
	fmt.Println(process(program, 12, 2))


	// part 2
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			result := process(program, noun, verb)
			if (result == 19690720) {
				fmt.Println(100 * noun + verb)
				goto done
			}
		}
	}
	done:
}

func process(prog []int, noun int, verb int) int {
	pc := 0

	program := make([]int, len(prog))
	copy(program, prog)

	program[1] = noun
	program[2] = verb

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

	return program[0]
}
