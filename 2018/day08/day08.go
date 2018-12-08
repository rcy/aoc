package main

import (
	"flag"
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
	"log"
)

var input = flag.String("input", "../inputs/day08.txt", "The input to the problem.")

func arrayFromInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := strings.Replace(string(bytes), "\n", "", -1)
	var nums []int
	for _, nstr := range strings.Split(str, " ") {
		num, _ := strconv.Atoi(nstr)
		nums = append(nums, num)
	}
	return nums
}

type Node struct {
	cc int
	mc int
	children []*Node
	metadata []int
	mv int
}

func main() {
	flag.Parse()

	numbers := arrayFromInput(*input)

	node, _ := createNode(numbers)

	total := 0
	for _, metadata := range collectMetadata(node) {
		total += metadata
	}
	fmt.Println("partA:", total)

	fmt.Println("partB:", value(node))
}

func collectMetadata(node *Node) []int {
	metadata := []int{}
	for _, n := range node.children {
		metadata = append(metadata, collectMetadata(n)...)
	}
	metadata = append(metadata, node.metadata...)
	return metadata
}

func value(node *Node) int {
	v := 0
	if len(node.children) == 0 {
		v = node.mv
	} else {
		for _, i := range node.metadata {
			if i-1 < len(node.children) {
				v += value(node.children[i-1])
			}
		}
	}
	return v
}

func createNode(numbers []int) (*Node, int) {
	cc := numbers[0]
	mc := numbers[1]

	p := 2

	node := Node{cc: cc, mc: mc}

	for i := 0; i < cc; i++ {
		childNode, consumed := createNode(numbers[p:])
		node.children = append(node.children, childNode)
		p += consumed
	}

	node.metadata = numbers[p:p+mc]
	for _, i := range node.metadata {
		node.mv += i
	}
	p += mc

	return &node, p
}
