package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
)

var input = flag.String("input", "../inputs/day09.txt", "The input to the problem.")

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

type Node struct {
	next *Node
	prev *Node
	value int
}

type State struct {
	player int
	node *Node
	first *Node
	numPlayers int
	scores map[int]int
}

func main() {
	flag.Parse()

	// n := &Node{value: 1}
	// n.next = n
	// m := &Node{value: 2}
	// insertn(n, m)
	// fmt.Println(n.value, n.next.value)
	// return

	//scores := run(9, 25)
	//scores := run(10, 1618)
	scores := run(476, 7143100)
	//scores := run(476, 71431)
	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	fmt.Println("max", max)
}


// insert node2 after node1
func insertn(node *Node, newNode *Node) {
	//fmt.Println("insertn", node.value, newNode.value)
	newNode.next = node.next
	newNode.prev = node
	node.next = newNode
	newNode.next.prev = newNode
	//node2.prev = node1
}

func run(players int, worth int) map[int]int {
	node := &Node{value: 0}
	node.prev = node
	node.next = node
	state := State{ first: node, node: node, player: 0, numPlayers: players, scores: make(map[int]int) }

	//dumpState(state)

	for i := 1; i <= worth; i++ {
		state = placeMarble(state, i)
		//dumpState(state)
		// if i % 10000 == 0 {
		// 	fmt.Println(i)
		// }
	}
	return state.scores
}

func placeMarble(state State, value int) State {
	state.player = (state.player + 1) % state.numPlayers
	if value % 23 == 0 {
		//fmt.Println("***23 thing")
		state.node = state.node.prev.prev.prev.prev.prev.prev.prev
		score := value + state.node.value
		state.scores[state.player] += score
		//fmt.Println("do stuff 23", state.circle[removePos])
		state.node.prev.next = state.node.next
		state.node = state.node.next
	} else {
		//state.node = state.node.next
		m := &Node{value: value}
		insertn(state.node.next, m)
		state.node = m
	}
	return state
}

func dumpState(state State) {
	fmt.Printf("[%d] ", state.player)
	
	node := state.first
	for {
		if node == state.node {
			fmt.Printf("(%2d)", node.value)
		} else {
			fmt.Printf(" %2d ", node.value)
		}
		node = node.next
		if node == state.first {
			break
		}
	}
	// for i, v := range state.circle {
	// 	fmt.Printf("%2d", v)
	// 	if state.position == i {
	// 		fmt.Print(")")
	// 	} else if state.position - 1 == i {
	// 		fmt.Print("(")
	// 	} else {
	// 		fmt.Print(" ")
	// 	}
	// }
	fmt.Println()
	//fmt.Println(state.first.value, state.first.next.value)
}
