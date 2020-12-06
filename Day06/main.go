package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func countQuestions(group []byte) (int, int) {
	counts := map[byte]int{}
	numPeople := 1
	for _, question := range group {
		if question == '\n' {
			numPeople++
		} else {
			counts[question]++
		}
	}

	numAllYes := 0
	for _, c := range counts {
		if c == numPeople {
			numAllYes++
		}
	}

	return len(counts), numAllYes
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	groups := bytes.Split(input, []byte("\n\n"))

	sum1, sum2 := 0, 0
	for _, group := range groups {
		a, b := countQuestions(group)
		sum1 += a
		sum2 += b
	}

	fmt.Println(sum1, sum2)
}
