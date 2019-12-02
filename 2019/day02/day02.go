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

func process(program []int, noun int, verb int) int {
	ip := 0

	p := make([]int, len(program))
	copy(p, program)

	p[1] = noun
	p[2] = verb

	for p[ip] != 99 {
		if (p[ip] == 1) {
			// addition
			p[p[ip + 3]] = p[p[ip + 1]] + p[p[ip + 2]]
		} else {
			// multiplication
			p[p[ip + 3]] = p[p[ip + 1]] * p[p[ip + 2]]
		}
		ip += 4
	}

	return p[0]
}
