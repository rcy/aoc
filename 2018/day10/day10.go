package main

import (
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"io/ioutil"
	"log"
	"time"
)

var input = flag.String("input", "../inputs/day10.txt", "The input to the problem.")

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

type Point struct {
	i int
	x int
	y int
	vx int
	vy int
}

func main() {
	flag.Parse()
	fmt.Println("DAY10\n-----\n\n")

	lines := linesFromInput(*input)

	points := []Point{}

	re := regexp.MustCompile("^position=< *(.+), *(.+)> velocity=< *(.+), *(.+)>$")
	for i, line := range lines {
		if len(line) > 0 {
			match := re.FindStringSubmatch(line)
			x,_ := strconv.Atoi(match[1])
			y,_ := strconv.Atoi(match[2])
			vx,_ := strconv.Atoi(match[3])
			vy,_ := strconv.Atoi(match[4])
			point := Point{i:i, x:x, y:y, vx:vx, vy:vy}
			points = append(points, point)
		}
	}



	lastPoints := make([]Point, len(points))
	lastArea := 0
	//copy(lastPoints, points)

	for i := 0; ; i++ {
		area := area(points)
		if i > 0 {
			if lastArea < area {
				draw(lastPoints)
				break
			}
		}
		lastArea = area
		copy(lastPoints, points)
		move(points)

		fmt.Println(i)
	}
	return
	move(points)
	fmt.Println(area(points))
	move(points)
	fmt.Println(area(points))
	move(points)

	fmt.Println(points)
	fmt.Println(lastPoints)

	return

	animate(points)
}

func area(points []Point) int {
	var minx int
	var miny int
	var maxx int
	var maxy int

	for i, p := range points {
		if i > 0 {
			if p.x > maxx {
				maxx = p.x
			}
			if p.y > maxy {
				maxy = p.y
			}
			if p.x < minx {
				minx = p.x
			}
			if p.y < miny {
				miny= p.x
			}
		}
	}

	return (minx - maxx) * (miny - maxy)
}

func animate(points []Point) {
	for {
		draw(points)
		fmt.Println("---")
		move(points)
		time.Sleep(100 * time.Millisecond)
	}
}

func move(points []Point) {
	for i := 0; i < len(points); i++ {
		points[i].x += points[i].vx
		points[i].y += points[i].vy
	}
}

func display(points []Point) {
	//minx := findMinx(points)
	for _, p := range points {
		//fmt.Printf("p.x: %d\n", p.x)
		//plot(p.x - minx)
		fmt.Printf("%d -- %d,%d (%d, %d)\n", p.i, p.x,p.y, p.vx, p.vy)
	}
}

func draw(points []Point) {
	minx := findMinx(points)
	//fmt.Printf("minx: %d\n", minx)
	x := minx
	for i, p := range sorty(points) {
		if i > 0 {
			if points[i].y != points[i-1].y {
				for j := points[i-1].y; j < points[i].y; j++ {
					fmt.Println()
				}
				x = minx
			}
			if points[i].x == points[i-1].x {
				continue
			}
		}
		plot(p.x - x)
		x += p.x - x + 1
	}
	fmt.Println()
}

func plot(x int) {
	for i := 0; i < x; i++ {
		fmt.Print(" ")
	}
	fmt.Print("#")
}

func findMinx(points []Point) int {
	min := 10000
	for _, p := range points {
		if p.x < min {
			min = p.x
		}
	}
	return min
}

func sorty(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].y < points[j].y {
			return true
		}
		if points[i].y > points[j].y {
			return false
		}
		return points[i].x < points[j].x
	})
	return points
}
