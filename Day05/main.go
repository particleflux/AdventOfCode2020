package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getNumber(id string) int {
	num := 0
	l := len(id)
	for i := l - 1; i >= 0; i-- {
		if id[i] == 'B' || id[i] == 'R' {
			num |= 1 << (l - i - 1)
		}
	}
	return num
}

func getId(row, col int) int {
	return (row << 3) + col
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	ids := make([]int, 0)
	highest := -1
	for {
		var line string
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		id := getId(getNumber(line[:7]), getNumber(line[7:]))
		if id > highest {
			highest = id
		}

		ids = append(ids, id)
	}

	fmt.Println(highest)

	sort.Ints(ids)
	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1]-ids[i] == 2 {
			fmt.Println(ids[i] + 1)
		}
	}
}
