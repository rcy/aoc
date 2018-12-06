package main

import (
	"flag"
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
	"log"
)

var input = flag.String("input", "../inputs/day06.txt", "The input to the problem.")

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

type coord struct {
	infinite bool
	count int
	x int
	y int
}

var maxX int
var minX int
var maxY int
var minY int
var coords []coord

func main() {
	flag.Parse()

	lines := linesFromInput(*input)

	minX = 1000
	maxX = 0
	minY = 1000
	maxY = 0
	
	for _, line := range lines {
		if len(line) > 0 {
			s := strings.Split(line, ", ")
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])

			c := coord{x: x, y: y, count: 0}
			coords = append(coords, c)

			if x < minX { minX = x }
			if x > maxX { maxX = x }
			if y < minY { minY = y }
			if y > maxY { maxY = y }
		}
	}

	grid := make([]string, (maxX + 1) * (maxY + 1))

	safeArea := 0

	// for each point on the grid mark the closest coord index
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			closestValue := 10000
			closestIndex := -1
			conflict := false

			totalDistance := 0

			for ci, c := range coords {
				distance := Abs(x - c.x) + Abs(y - c.y)

				if distance == closestValue {
					conflict = true
				}
				if distance < closestValue {
					closestValue = distance
					closestIndex = ci
					conflict = false
				}

				totalDistance += distance
			}
			if !conflict {
				set(grid, x, y, closestIndex)
			}
			if totalDistance < 10000 {
				safeArea += 1
			}
		}
	}

	max := 0
	for _, c := range coords {
		if !c.infinite {
			if c.count > max {
				max = c.count
			}
		}
	}
	fmt.Println("partA: ", max)
	fmt.Println("partB: ", safeArea)

	drawGrid(grid)
}


// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func set(grid []string, x int, y int, v int) {
	grid[x + (maxX + 1)*y] = string(65 + v)
	coords[v].count++
	if x == 0 || y == 0 || x == maxX || y == maxY {
		coords[v].infinite = true
	}
}

func get(grid []string, x int, y int) string {
	return grid[x + (maxX + 1)*y]
}

func drawGrid(grid []string) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			v := get(grid, x, y)
			if (v == "") {
				fmt.Print(".")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
}
