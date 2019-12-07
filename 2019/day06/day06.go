package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
)

var input = flag.String("input", "./input", "The input to the problem.")

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

type Node struct {
	name string
	orbits []*Node
	center *Node
	transfers int
}

var nodeMap map[string]*Node

func main() {
	flag.Parse()

	lines := linesFromInput(*input)

	nodeMap = make(map[string]*Node)

	for _, line := range lines {
		if len(line) > 0 {
			objects := strings.Split(line, ")")
			makeOrbit(objects[0], objects[1])			
		}
	}
	fmt.Println("Part One:", countOrbits(lookup("COM"), 0))
	fmt.Println("Part Two:", countTransfers())
}

func countTransfers() int {
	for you := lookup("YOU").center; you.name != "COM"; you = you.center {
		you.center.transfers += 1 + you.transfers
	}

	transfers := 0
	for san := lookup("SAN").center; san.name != "COM"; san = san.center {
		transfers += 1 + san.center.transfers
		if (san.center.transfers > 0) {
			return transfers
		}
	}
	return 0
}

func countOrbits(node *Node, depth int) int {
	if len(node.orbits) == 0 {
		return depth
	} else {
		count := depth
		for _, onode := range(node.orbits) {
			count += countOrbits(onode, depth + 1)
		}
		return count
	}
}

func lookup(name string) *Node {
	return nodeMap[name]
}

func create(name string) *Node {
	node := Node{name: name}
	nodeMap[name] = &node
	return &node
}

// o orbits around c
func makeOrbit(c string, o string) {
	center := lookup(c)
	orbit := lookup(o)
	if center == nil {
		center = create(c)
	}
	if orbit == nil {
		orbit = create(o)
	}
	center.orbits = append(center.orbits, orbit)
	orbit.center = center
}
