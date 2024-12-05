package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rules map[int][]int

func main() {
	// updates, rules := parseData("example_input.txt")
	updates, rules := parseData("input.txt")

	part1(rules, updates)
	part2(rules, updates)
}

func part2(rules Rules, updates [][]int) {
	total := 0

	for _, update := range updates {
		inUpdate := make(map[int]int)
		for i, n := range update {
			inUpdate[n] = i
		}

		if !validUpdate(update, inUpdate, rules) {
			newUpdate := fixUpdate(update, rules)
			total += newUpdate[len(newUpdate)/2]
		}
	}

	fmt.Println(total)
}

func fixUpdate(u []int, rules Rules) []int {
	update := make([]int, len(u))
	for i, n := range u {
		update[i] = n
	}

	newRules := make(map[int]map[int]bool)
	for k, vs := range rules {
		for _, v := range vs {
			if _, ok := newRules[k]; !ok {
				newRules[k] = make(map[int]bool)
			}
			newRules[k][v] = true
		}
	}

	i := 0
	for i < len(update) {
		if vals, ok := newRules[update[i]]; ok {
			origI := i

			for j := i - 1; j >= 0; j-- {
				if _, ok := vals[update[j]]; ok {
					update = move(update, i, j)
					i = j
				}
			}

			if i == origI {
				i++
			}
		} else {
			i++
		}
	}

	return update
}

func move(s []int, i, j int) []int {
	t := s[i]

	for n := i; n > j; n-- {
		s[n] = s[n-1]
	}

	s[j] = t

	return s
}

func part1(rules Rules, updates [][]int) {
	total := 0

	for _, update := range updates {
		inUpdate := make(map[int]int)
		for i, n := range update {
			inUpdate[n] = i
		}

		if validUpdate(update, inUpdate, rules) {
			total += update[len(update)/2]
		}
	}

	fmt.Println(total)
}

func validUpdate(update []int, updateMap map[int]int, rules Rules) bool {
	for i, n := range update {
		if rule, ok := rules[n]; ok {
			for _, other := range rule {
				if p, ok := updateMap[other]; ok {
					if p < i {
						return false
					}
				}
			}
		}
	}

	return true
}

func parseData(filename string) ([][]int, Rules) {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(bytes)
	parts := strings.Split(data, "\n\n")

	if len(parts) != 2 {
		panic("wrong number of parts")
	}

	return getUpdates(parts[1]), getRules(parts[0])
}

func getRules(s string) Rules {
	lines := strings.Split(s, "\n")
	rules := make(Rules)

	for _, line := range lines {
		sides := strings.Split(line, "|")
		one, _ := strconv.Atoi(sides[0])
		two, _ := strconv.Atoi(sides[1])

		rules[one] = append(rules[one], two)
	}

	return rules
}

func getUpdates(s string) [][]int {
	lines := strings.Split(s, "\n")
	var updates [][]int

	for _, line := range lines {
		var nums []int
		ns := strings.Split(line, ",")

		for _, n := range ns {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		updates = append(updates, nums)
	}

	return updates
}
