package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, []int) {
	decks := strings.Split(strings.TrimSpace(input), "\n\n")
	players := make([][]int, 2)

	for player, deck := range decks {
		lines := strings.Split(deck, "\n")
		players[player] = make([]int, len(lines) - 1)

		for i, line := range lines {
			if i == 0 {
				continue
			}

			players[player][i - 1], _ = strconv.Atoi(line)
		}
	}

	return players[0], players[1]
}

func solvePart1(player1, player2 []int) int {
	for {
		if len(player1) == 0 || len(player2) == 0 {
			break
		}

		c1, c2 := player1[0], player2[0]
		player1, player2 = player1[1:], player2[1:]
		if c1 > c2 {
			player1 = append(player1, c1, c2)
		} else {
			player2 = append(player2, c2, c1)
		}
	}

	score := 0
	if len(player1) > 0 {
		fmt.Println("player 1 wins")
		score = calcScore(player1)
	} else {
		fmt.Println("player 2 wins")
		score = calcScore(player2)
	}

	return score
}

func calcScore(deck []int) int {
	score := 0
	for i, val := range deck {
		score += val * (len(deck) - i)
	}

	return score
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	player1, player2 := parseInput(string(input))

	fmt.Println(solvePart1(player1, player2))
}
