package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
)

var input = flag.String("input", "../inputs/day07.txt", "The input to the problem.")

type Node struct {
	letter string
	time int
	deps []string
}

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

var nodemap map[string]*Node

func main() {
	flag.Parse()

	lines := linesFromInput(*input)

	nodemap = make(map[string]*Node)

	baseTime := 61

	for _, line := range lines {
		if len(line) > 0 {
			tokens := strings.Split(line, " ")
			n := findOrMakeNode(tokens[1], baseTime)
			m := findOrMakeNode(tokens[7], baseTime)
			m.deps = append(m.deps, n.letter)
		}
	}

	workers := [5]*Node{}

	done := false

	str := ""

	for s := 0; !done; s++ {
		for i, n := range workers {
			if n == nil {
				workers[i] = popNode()
			} else {
				n.time--
				if n.time == 0 {
					completeNode(n)
					str += n.letter
					workers[i] = popNode()
				}
			}
		}

		fmt.Printf("% 3d", s)
		done = true
		for _, w := range workers {
			if w == nil {
				fmt.Print(" .")
			} else {
				fmt.Printf(" %s", w.letter)
			}
			if w != nil {
				done = false
			}
		}
		fmt.Printf("\t%s\n", str)
	}
}

func popNode() *Node {
	var next *Node

	// remove lowest lettered, no dependency node
	for _, node := range nodemap {
		//fmt.Println("looking at", node)
		if len(node.deps) == 0 {
			if next == nil || node.letter < next.letter {
				//fmt.Println("teeing up", node)
				next = node
			}
		}
	}
	//fmt.Println("deleting", next)


	if next != nil {
		delete(nodemap, next.letter)
	}
	
	return next
}

func completeNode(done *Node) {
	// remove node from dependency lists
	for _, node := range nodemap {
		i := indexOf(node.deps, done.letter)
		if i >= 0 {
			node.deps = append(node.deps[:i], node.deps[i+1:]...)
		}
	}
}

func indexOf(a []string, e string) int {
	for i, v := range a {
		if v == e {
			return i
		}
	}
	return -1
}

func printNodeMap() {
	for k, v := range nodemap {
		fmt.Println(k, v.time, v.deps)
	}
}

func findOrMakeNode(letter string, baseTime int) *Node {
	node := nodemap[letter]
	if node == nil {
		node = &Node{letter: letter, time: baseTime + int(letter[0]) - 65}
		nodemap[letter] = node
	}
	return node
}
