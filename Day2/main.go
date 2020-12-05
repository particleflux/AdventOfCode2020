package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValid(min, max int, chr byte, password string) bool {
	count := strings.Count(password, string(chr))
	return count >= min && count <= max
}

func isValidV2(min, max int, chr byte, password string) bool {
	return (password[min - 1] == chr) != (password[max - 1] == chr)
}

func main() {
	var min, max, numValid, numValidV2 int
	var required byte
	var password string

	for {
		var r, c string
		if _, err := fmt.Scanf("%s %s %s", &r, &c, &password); err != nil {
			break
		}

		required = c[0]
		parts := strings.Split(r, "-")
		min,_ = strconv.Atoi(parts[0])
		max,_ = strconv.Atoi(parts[1])

		if isValid(min, max, required, password) {
			numValid++
		}
		if isValidV2(min, max, required, password) {
			numValidV2++
		}
	}

	fmt.Println(numValid)
	fmt.Println(numValidV2)
}
