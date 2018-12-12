package main

import (
	"fmt"
)

var memo(map[int]*int)

type SquarePower struct {
	x int
	y int
}

var smemo(map[SquarePower]*int)

func main() {
	memo = make(map[int]*int)
	smemo = make(map[SquarePower]*int)

	tests()

	//fmt.Println(largestBlock(6548))
}

func tests() {
	// fmt.Println(cellPower(3, 5, 18))
	// fmt.Println(cellPower(122, 79, 57))
	// fmt.Println(cellPower(217, 196, 39))
	// fmt.Println(cellPower(101, 153, 71))

	//fmt.Println(blockPower(21, 61, 2, 42))
	//fmt.Println(smemo)
	//fmt.Println(blockPower(21, 61, 3, 42))
	//fmt.Println(blockPower(21, 61, 4, 42))
	//fmt.Println(largestBlock(42))
	fmt.Println(largestBlock(6548))
}

func cellPower(x, y, serial int) int {
	var power int
	m := memo[x * 301 + y]
	if m != nil {
		power = *m
	} else {
		rackId := x + 10
		m := (((((rackId * y) + serial) * rackId) / 100) % 10) - 5
		memo[x * 301 + y] = &m
		power = m
	}
	//fmt.Println("cellPower", x, y, power)
	return power
}

func blockPower(startx, starty, size, serial int) int {
	total := 0
	smallerBlockPower := smemo[SquarePower{startx, starty}]
	if smallerBlockPower != nil {
		//fmt.Println("smallerBlockPower", smallerBlockPower)
		// iio
		// iio
		// ooo
		// we know power of area covered by i, compute power of area i+o
		total = *smallerBlockPower
		for x := startx; x < startx + size; x++ {
			total += cellPower(x, starty + size - 1, serial)
		}
		for y := starty; y < starty + size - 1; y++ {
			total += cellPower(startx + size - 1, y, serial)
		}
	} else {
		for y := starty; y < starty + size; y++ {
			for x := startx; x < startx + size; x++ {
				total += cellPower(x, y, serial)
			}
		}
	}
	smemo[SquarePower{x: startx, y: starty}] = &total
	return total
}


func largestBlock(serial int) int {
	max := 0
	for size := 1; size <= 300; size++ {
		fmt.Print(".")
		for y := 1; y <= 300 - size + 1; y++ {
			for x := 1; x <= 300 - size + 1; x++ {
				power := blockPower(x, y, size, serial)
				if power > max {
				 	max = power
				 	fmt.Println(max, x, y, size)
				}
			}
		}
	}
	return max
}
