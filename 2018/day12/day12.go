package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var input = flag.String("input", "../inputs/day12.txt", "The input to the problem.")

func linesFromInput(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return strings.Split(str, "\n")
}

func loadFromFile() (Row, Rules) {
	flag.Parse()

	lines := linesFromInput(*input)

	rowRe := regexp.MustCompile("^initial state: (.+)$")
	match := rowRe.FindStringSubmatch(lines[0])
	row := Row{pots: match[1], origin: 0}

	ruleRe := regexp.MustCompile("^(.....) => (.)$")
	var rules Rules
	for _, line := range lines[1:] {
	 	if len(line) > 0 {
			match = ruleRe.FindStringSubmatch(line)
			rule := Rule{match[1], match[2]}
			rules = append(rules, rule)
	 	}
	}
	return row, rules
}

type Row struct {
	pots string
	origin int
}

type Rule struct {
	pat string
	out string
}

type Rules []Rule

func getPot(r Row, index int) bool {
	return r.pots[index] == '#'
}

func checkRuleMatch(pots string, pos int, rule Rule) bool {
	var str string
	str = pots[pos-2:pos+3]
	match := str == rule.pat
	return match
}

func applyRules(row Row, pos int, rules Rules) string {
	for _, rule := range rules {
		if checkRuleMatch(row.pots, pos, rule) {
			return rule.out
		}
	}
	return "."
}

func sequence(row Row, rules Rules) Row {
	str := ".." 
	for i := 2; i < len(row.pots) - 2; i++ {
		str = str + applyRules(row, i, rules)
		//fmt.Println(i, row.pots[i-2:i+3])
	}
	row.pots = str
	return row
}

func pad(row Row) Row {
	padding := "...."
	for i := 0; i < 5; i++ {
		if row.pots[i] == '#' {
			row.pots = padding[:4-i] + row.pots
			row.origin += 4-i
			break
		}
	}
	for i := 0; i < 5; i++ {
		if row.pots[len(row.pots) - i - 1] == '#' {
			row.pots = row.pots + padding[:4-i]
		}
	}
	return row
}

func main() {
	row, rules := loadFromFile()
	fmt.Println(rules)
	fmt.Println("================")
	row = pad(row)
	fmt.Println(row)

	for i := 0; i < 20; i++ {
		row = pad(sequence(row, rules))
		fmt.Println(row)
	}

	total := 0
	for i := 0; i < len(row.pots); i++ {
		if row.pots[i] == '#' {
			total += i - row.origin
		}
	}
	fmt.Println(total)
}
