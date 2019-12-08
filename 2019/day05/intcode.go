package main
import (
//	"fmt"
)

func code(value int) []int {
	params := []int{0,0,0,0}
	params[0] = value % 100 // opcode
	params[1] = value / 100 % 10
	params[2] = value / 100 / 10 % 10
	params[3] = value / 100 / 10 / 10 % 10

	return params
}

func lookup(program []int, param int, mode int) int {
	if (mode == 1) {
		return param
	} else {
		return program[param]
	}
}

type Program []int

func run(program Program, input []int) int {
	//fmt.Println("run", program, input)
	ip := 0
	ic := 0
	var out int

	p := make(Program, len(program))
	copy(p, program)

	// p[1] = noun
	// p[2] = verb

	for p[ip] != 99 {
		//fmt.Println(p)

		inst := code(p[ip])
		//fmt.Println("-- INSTRUCTION", p[ip], inst)
		if (inst[0] == 1) {
			// addition
//			fmt.Println("MULT", p[ip:ip+4])
			p1 := lookup(p, p[ip+1], inst[1])
			p2 := lookup(p, p[ip+2], inst[2])
			p[p[ip + 3]] = p1 + p2
			ip += 4
		} else if (inst[0] == 2) {
			// multiplication
//			fmt.Println("MULT", p[ip:ip+4])
			p1 := lookup(p, p[ip+1], inst[1])
			p2 := lookup(p, p[ip+2], inst[2])
			p[p[ip + 3]] = p1 * p2
			ip += 4
		} else if (inst[0] == 3) {
			// storeinput
//			fmt.Println("INPUT", p[ip:ip+2], input[ic])
			p[p[ip + 1]] = input[ic]
			ic += 1
			ip += 2
		} else if (inst[0] == 4) {
			// writeoutput
//			fmt.Println("EMIT", p[ip:ip+2])
			out = lookup(p, p[ip + 1], inst[1])
			// if out == 0 {
			// 	fmt.Println(".")
			// } else {
			// 	fmt.Println(".output", out)
			// }
			//printOutput(out)
			ip += 2
		} else if (inst[0] == 5) {
			// jump-if-true
//			fmt.Println("JUMP-IF-TRUE", p[ip:ip+3])
			if (lookup(p, p[ip+1], inst[1]) > 0) {
				ip = lookup(p, p[ip+2], inst[2])
			} else {
				ip += 3
			}
		} else if (inst[0] == 6) {
			// jump-if-false
//			fmt.Println("JUMP-IF-FALSE", p[ip:ip+3])
			if (lookup(p, p[ip+1], inst[1]) == 0) {
				ip = lookup(p, p[ip+2], inst[2])
			} else {
				ip += 3
			}
		} else if (inst[0] == 7) {
			// less-than
//			fmt.Println("LESS-THAN", p[ip:ip+4])
			if (lookup(p, p[ip+1], inst[1]) < lookup(p, p[ip+2], inst[2])) {
				p[p[ip+3]] = 1
			} else {
				p[p[ip+3]] = 0
			}
			ip += 4
		} else if (inst[0] == 8) {
			// equals
//			fmt.Println("EQUALS", p[ip:ip+4])
			if (lookup(p, p[ip+1], inst[1]) == lookup(p, p[ip+2], inst[2])) {
				p[p[ip+3]] = 1
			} else {
				p[p[ip+3]] = 0
			}
			ip += 4
		} else {
			panic("unknown opcode")
		}
	}

	return out
}
