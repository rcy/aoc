package main

import "fmt"

type record struct {
	total int
	minutes [60]int
}

func main() {
	m := make(map[int]*record)
	fmt.Println(m[0] == nil)
	m[0] = &record{total: 100}

	m[0].total++
	fmt.Println(m[0].total)
}
