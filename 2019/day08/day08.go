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

	part1(lines[0])
}

func part1(input string) {
	min := -1
	var result int
	for i, layer := range getLayers(input, 25, 6) {
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

	fmt.Println("numLayers", numLayers)
	var layer string

	for i := 0; i < numLayers; i++ {
		offset := i * layerSize
		layer = data[offset:offset + layerSize]
		layers[i] = layer
	}
	return layers
}
