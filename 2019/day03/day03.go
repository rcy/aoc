package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var input = flag.String("input", "./input", "The input to the problem.")

type Grid struct {
	size int
	data [][]byte
}

func (g *Grid) print() {
	for y := range g.data {
		for x := range g.data[y] {
			if (g.size/2 == y && g.size/2 == x) {
				fmt.Print("o")
			} else if (g.data[y][x] == 0) {
				fmt.Print(" ")
			} else if (g.data[y][x] == 1) {
				fmt.Print("1")
			} else if (g.data[y][x] == 2) {
				fmt.Print("2")
			} else {
				fmt.Print("x")
			}
		}
		fmt.Println()
	}
}

func (g *Grid) findNearestCrossDistance() float64 {
	var min float64 = 0
	for y := range g.data {
		for x := range g.data[y] {
			if (g.data[y][x] == 3) {
				distance := math.Abs(float64(y - g.size/2)) + math.Abs(float64(x - g.size/2))
				fmt.Println(distance, ":", x, y)
				if min == 0 || distance < min {
					min = distance
				}
			}
		}
	}
	return min
}

func (g *Grid) init() {
	fmt.Println( g.size, "grid size")
	g.data = make([][]byte, g.size)
	for i := range g.data {
		g.data[i] = make([]byte, g.size)
	}
}

func (g *Grid) run(path []string, code byte) {
	var x = 0
	var y = 0
	for _, leg := range path {
		dir := leg[0]
		len, _ := strconv.Atoi(leg[1:])
		//fmt.Printf("%c %d\n", dir, len)
		for i := 0; i < len; i++ {
			if (dir == 'R') {
				x += 1
			} else if (dir == 'L') {
				x -= 1
			} else if (dir == 'U') {
				y -= 1
			} else if (dir == 'D') {
				y += 1
			} else {
				panic(0)
			}
			g.set(x, y, code)
		}
	}
}

func (g *Grid) set(x int, y int, code byte) {
	g.data[g.size / 2 + y][g.size / 2 + x] |= code
}

func pathFromLine(line string) []string {
	return strings.Split(line, ",")
}

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

	grid := Grid{size: 211111}
	grid.init()

	var paths [2][]string
	paths[0] = pathFromLine(lines[0])
	paths[1] = pathFromLine(lines[1])

	grid.run(paths[0], 1)
	grid.run(paths[1], 2)

//	grid.print()

	fmt.Println(grid.findNearestCrossDistance())
}


// ...........
// ...........
// ...........
// ....+----+.
// ....|....|.
// ....|....|.
// ....|....|.
// .........|.
// .o-------+.
// ...........

// ...........
// .+-----+...
// .|.....|...
// .|..+--X-+.
// .|..|..|.|.
// .|.-X--+.|.
// .|..|....|.
// .|.......|.
// .o-------+.
// ...........
