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

func str2num(str string) int {
	n := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '#' {
			n++
		}
		n <<= 1
	}
	return n >> 1
}

func loadFromFile() (int, [32]int) {
	flag.Parse()

	lines := linesFromInput(*input)

	rowRe := regexp.MustCompile("^initial state: (.+)$")
	match := rowRe.FindStringSubmatch(lines[0])
	fmt.Println(match[1])
	n := str2num(match[1])

	var rules [32]int
	ruleRe := regexp.MustCompile("^(.....) => (.)$")
	for _, line := range lines[2:] {
		if len(line) > 0 {
			match = ruleRe.FindStringSubmatch(line)
			if match[2] == "#" {
				rules[str2num(match[1])] = 1
			}
		}
	}
	return n, rules
}

func render(n int) {
	for ; n > 0; n >>= 1 {
		if n & 1 == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
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
	n, rules := loadFromFile()
	//fmt.Println(n)
	m := 0
	fmt.Println("load", n, rules)
	n <<= 2
	for ; n >= 1; n >>= 1 {
		//fmt.Print(n)
		//render(n & 31)
		for r := 0; r < 32; r++ {
			fmt.Print(n, r)
			if rules[r] == 1 {
				if n & 31 == r {
					fmt.Print("match")
					m += 1
				}
			}
			fmt.Println(" ", m)
		}
		m <<= 1
	}
	m <<= 1
	fmt.Println("\n---", m)
	render(m)
}
