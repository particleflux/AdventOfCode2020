package main

import (
	"fmt"
	"github.com/mpvl/unique"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Bag struct {
	Color    string
	Amount   int
	Children map[string]Bag
}

func parseRules(input string) map[string]Bag {
	rules := make(map[string]Bag)
	rulesInput := strings.Split(strings.TrimSpace(input), "\n")

	r := regexp.MustCompile(`(^(\w+ \w+))|(\d+) (\w+ \w+) bags?[,.]`)
	for _, rule := range rulesInput {
		matches := r.FindAllStringSubmatch(rule, -1)

		bag := Bag{
			Color:    matches[0][0],
			Amount:   0,
			Children: nil,
		}
		// single match result is "contains no other bag"
		if len(matches) > 1 {
			bag.Children = make(map[string]Bag, 2)
			for i := 1; i < len(matches); i++ {
				amount, _ := strconv.Atoi(matches[i][3])
				bag.Children[matches[i][4]] = Bag{
					Color:    matches[i][4],
					Amount:   amount,
					Children: nil,
				}
			}
		}

		rules[bag.Color] = bag
	}

	return rules
}

func solvePart1(rules map[string]Bag) int {
	// a map correlating children to possible parents
	relations := make(map[string][]string, 100)
	for _, bag := range rules {
		if bag.Children == nil {
			continue
		}

		for color := range bag.Children {
			_, ok := relations[color]
			if !ok {
				relations[color] = make([]string, 0)
			}
			relations[color] = append(relations[color], bag.Color)
		}
	}

	parents := getParents("shiny gold", relations)
	sort.Strings(parents)
	unique.Strings(&parents)
	return len(parents)
}

func getParents(child string, relations map[string][]string) []string {
	parents := make([]string, 0)
	for _, parent := range relations[child] {
		parents = append(parents, parent)
		parents = append(parents, getParents(parent, relations)...)
	}

	return parents
}

func solvePart2(rules map[string]Bag, bag string) int {
	sum := 1
	for _, child := range rules[bag].Children {
		sum += child.Amount * solvePart2(rules, child.Color)
	}
	return sum
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	rules := parseRules(string(input))

	fmt.Println(solvePart1(rules))
	fmt.Println(solvePart2(rules, "shiny gold") - 1)
}
