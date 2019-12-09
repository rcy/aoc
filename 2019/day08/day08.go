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

func main() {
	flag.Parse()

	lines := linesFromInput(*input)

	width := 25
	height := 6

	part1(lines[0], width, height)
	part2(lines[0], width, height)
}

func part2(input string, width int, height int) {
	result := make([]rune, width * height)
	layers := getLayers(input, width, height)

	for i := len(layers) - 1; i >= 0; i-- {
		layer := layers[i]
		for i, ch := range layer {
			if ch != '2' {
				result[i] = ch
			}
		}
	}

	for i, ch := range result {
		if i % width == 0 {
			fmt.Println()
		}
		if (ch == '1') {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}
}

func part1(input string, width int, height int) {
	min := -1
	var result int
	for i, layer := range getLayers(input, width, height) {
		counts := countChars(layer)
		if min == -1 || counts[0] < min {
			min = counts[0]
			result = counts[1] * counts[2]
			fmt.Println("layer", i, counts, layer, result)
		}
	}
}

func countChars(layer string) []int {
	result := []int{0,0,0}
	for _, ch := range layer {
		result[ch - '0']++
	}
	return result
}

func getLayers(data string, width int, height int) []string{
	layerSize := width * height

	numLayers := len(data) / layerSize

	layers := make([]string, numLayers)

	var layer string

	for i := 0; i < numLayers; i++ {
		offset := i * layerSize
		layer = data[offset:offset + layerSize]
		layers[i] = layer
	}
	return layers
}
