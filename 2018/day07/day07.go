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

// func main() {
// 	nodemap = make(map[string]*Node)
// 	n := findOrMakeNode("A")
// 	n.deps = append(n.deps, "A")
// 	fmt.Println(n)
// 	for k, v := range nodemap {
// 		fmt.Println(k, *v, n)
// 	}
// }

func main() {
	flag.Parse()

	lines := linesFromInput(*input)

	nodemap = make(map[string]*Node)

	for _, line := range lines {
		if len(line) > 0 {
			tokens := strings.Split(line, " ")
			n := findOrMakeNode(tokens[1])
			m := findOrMakeNode(tokens[7])
			m.deps = append(m.deps, n.letter)
		}
	}

	var node *Node

	for {
		node = popNode()
		if node == nil {
			break
		}
		fmt.Print(node.letter)
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

		// remove node from dependency lists
		for _, node := range nodemap {
			i := indexOf(node.deps, next.letter)
			if i >= 0 {
				node.deps = append(node.deps[:i], node.deps[i+1:]...)
			}
		}
	}
	
	return next
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
		fmt.Println(k, v.deps)
	}
}

func findOrMakeNode(letter string) *Node {
	node := nodemap[letter]
	if node == nil {
		node = &Node{letter: letter}
		nodemap[letter] = node
	}
	return node
}
