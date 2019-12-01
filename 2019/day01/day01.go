package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"strconv"
	"math"
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

	total := 0.0

	for _, line := range lines {
		if len(line) > 0 {
			n, _ := strconv.ParseFloat(line, 32)
			total += calcFuel(n)
		}
	}
	fmt.Printf("%f\n", total)
}

func calcFuel(n float64) float64 {
	fuel := math.Floor(n / 3) - 2
	if (fuel > 6) {
		fuel += calcFuel(fuel)
	}
	return fuel
}
