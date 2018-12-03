package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var input = flag.String("input", "../inputs/day03.txt", "The input to the problem.")

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

func squares(line string) (int, []string) {
	re := regexp.MustCompile("^#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)$")
	match := re.FindStringSubmatch(line)
	id, _ := strconv.Atoi(match[1])
	x, _ := strconv.Atoi(match[2])
	y, _ := strconv.Atoi(match[3])
	w, _ := strconv.Atoi(match[4])
	h, _ := strconv.Atoi(match[5])

	var a []string

	for i := x; i < x + w; i++ {
		for j := y; j < y + h; j++ {
			a = append(a, fmt.Sprintf("%d:%d", i, j))
		}
	}
	return id, a
}

func main() {
	flag.Parse()

	lines := linesFromInput(*input)

	m := make(map[string][]int)

	conflictIds := make(map[int]bool)

	for _, line := range lines {
		if len(line) > 0 {
			id, coords := squares(line)
			conflictIds[id] = false
			for _, coord := range coords {
				m[coord] = append(m[coord], id)
			}
		}
	}


	taken := 0
	for _, ids := range m {
		if len(ids) > 1 {
			taken++
			for _, id := range ids {
				conflictIds[id] = true
			}
		}
	}
	fmt.Println("partA:", taken)

	for id, conflict := range conflictIds {
		if !conflict {
			fmt.Println("partB:", id)
		}
	}
}
