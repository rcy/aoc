package main

import (
	"flag"
	"strconv"
	"fmt"
	"regexp"
	"strings"
	"io/ioutil"
	"log"
)

var input = flag.String("input", "../inputs/day04.txt", "The input to the problem.")

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

type record struct {
	id int
	total int
	minutes [60]int
}

type sleepmap = map[int]*record

var sleepMap = make(sleepmap)
var id int

func main() {
	flag.Parse()

	lines := linesFromInput(*input)

	lineRe := regexp.MustCompile("^\\[.+(..)\\] (.+)$")
	guardRe := regexp.MustCompile("Guard #([0-9]+) begins shift$")
	
	minute := 0
	event := ""
	since := -1

	for _, line := range lines {
		//fmt.Printf("%d %s\n", i, line)
			
		match := lineRe.FindStringSubmatch(line)
		if len(match) > 0 {
			minute, _ = strconv.Atoi(match[1])
			event = match[2]
			//fmt.Println(minute, "...", event)
		}
		
		if len(line) > 0 {
			match := guardRe.FindStringSubmatch(line)
			if (len(match) > 0) {
				id, _ = strconv.Atoi(match[1])
				if sleepMap[id] == nil {
					sleepMap[id] = &record{id: id}
				}
			} else {
				if event == "falls asleep" {
					since = minute
				} else {
					//fmt.Println(id, "slept from", since, "to", minute)
					sleepMap[id].total += minute - since
					for i := since; i < minute; i++ {
						sleepMap[id].minutes[i]++
					}
				}
			}
		}
	}

	partA(sleepMap)
	partB(sleepMap)
}

func partA(m sleepmap) {
	max := 0
	var maxRecord record
	var maxMinute int

	for _, v := range m {
		if v.total > max {
			max = v.total
			maxRecord = *v
		}
	}
	
	max = 0
	for k, v := range maxRecord.minutes {
		if v > max {
			max = v
			maxMinute = k
		}
	}
	fmt.Println("partA:", maxRecord.id * maxMinute)
}

func partB(m sleepmap) {
	maxValue := 0
	maxMinute := 0
	var maxRecord record
	for _, r := range m {
		for k, v := range r.minutes {
			if v > maxValue {
				maxValue = v
				maxMinute = k
				maxRecord = *r
			}
		}
	}
	fmt.Println("partB:", maxRecord.id * maxMinute)
}
