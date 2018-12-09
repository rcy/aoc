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

type State struct {
	player int
	circle []int
	position int
	value int
	numPlayers int
	scores map[int]int
}

func main() {
	flag.Parse()

	scores := run(9, 50)
	//scores := run(13, 79990)
	//scores := run(476, 71431)
	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}


func run(players int, worth int) map[int]int {
	state := State{ circle: make([]int, 1), player: 0, position: 0, value: 0, numPlayers: players, scores: make(map[int]int) }

	//dumpState(state)

	for i := 0; i < worth; i++ {
		state = placeMarble(state)
		dumpState(state)
		if i % 10000 == 0 {
			fmt.Println(i)
		}
	}
	return state.scores
}

func placeMarble(state State) State {
	state.value++
	state.player = (state.player + 1) % state.numPlayers

	if state.value % 23 == 0 {
		//fmt.Println(state.position)
		removePos := 0
		if state.position < 7 {
			removePos = len(state.circle) - (7 - state.position)
		} else {
			removePos = (state.position - 7) % len(state.circle)
		}
		score := state.value + state.circle[removePos]
		state.scores[state.player] += score
		//fmt.Println("do stuff 23", state.circle[removePos])
		state.circle = remove(state.circle, removePos)
		state.position = removePos
	} else {
		state.position = (state.position + 1) % len(state.circle) + 1
		state.circle = insert(state.circle, state.position, state.value)
	}

	return state
}

func insert(slice []int, index, value int) []int {
	new := make([]int, len(slice) + 1)
	for i, v := range slice[0:index] {
		new[i] = v
	}
	new[index] = value
	for i, v := range slice[index:] {
		new[index + i + 1] = v
	}
	return new
}

func remove(slice []int, index int) []int {
	new := make([]int, len(slice) - 1)
	for i, v := range slice[0:index] {
		new[i] = v
	}
	for i, v := range slice[index+1:] {
	 	new[index + i] = v
	}
	return new
}

func dumpState(state State) {
	fmt.Printf("[%d]", state.player)
	for i, v := range state.circle {
		fmt.Printf("%2d", v)
		if state.position == i {
			fmt.Print(")")
		} else if state.position - 1 == i {
			fmt.Print("(")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
