package main

import (
	"strconv"
	"fmt"
	"bufio"
	"os"
	"log"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var freq int64 = 0
	seen := map[int64]bool{}

	for {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			num, _ := strconv.ParseInt(text, 10, 64)
			freq += num
			if (seen[freq]) {
				fmt.Println("seen", freq)
				goto done
			}
			seen[freq] = true

		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		
		file.Seek(0, 0)
	}
done:
}
