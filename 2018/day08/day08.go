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
}

func main() {
	flag.Parse()

	numbers := arrayFromInput(*input)

	node, _ := createNode(numbers)

	total := 0
	for _, metadata := range collectMetadata(node) {
		total += metadata
	}
	fmt.Println(total)
}

func collectMetadata(node *Node) []int {
	metadata := []int{}
	fmt.Println(node.metadata)
	for _, n := range node.children {
		metadata = append(metadata, collectMetadata(n)...)
	}
	metadata = append(metadata, node.metadata...)
	return metadata
}

func createNode(numbers []int) (*Node, int) {
	fmt.Println("createNode", numbers)

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
	p += mc

	return &node, p
}
