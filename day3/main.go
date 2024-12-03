package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getInput(source string) string {
	file, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var output []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return strings.Join(output, "\n")
}

func partOneInstructions(data string) []string {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	return r.FindAllString(data, -1)
}

type matchType int

const (
	MUL matchType = iota
	DO
	DONT
)

type match struct {
	index     int
	string    string
	matchType matchType
}

func matchWithIndex(r *regexp.Regexp, data string, t matchType) []match {
	matches := r.FindAllString(data, -1)
	indices := r.FindAllStringIndex(data, -1)

	result := []match{}

	for i, m := range matches {
		result = append(result, match{
			index:     indices[i][0],
			string:    m,
			matchType: t,
		})
	}

	return result
}

func partTwoInstructions(data string) []match {
	r1 := regexp.MustCompile(`mul\(\d+,\d+\)`)
	r2 := regexp.MustCompile(`do\(\)`)
	r3 := regexp.MustCompile(`don't\(\)`)
	ins := []match{}

	muls := matchWithIndex(r1, data, MUL)
	ins = append(ins, muls...)
	dos := matchWithIndex(r2, data, DO)
	ins = append(ins, dos...)
	donts := matchWithIndex(r3, data, DONT)
	ins = append(ins, donts...)

	sort.Slice(ins, func(i, j int) bool {
		return ins[i].index < ins[j].index
	})

	return ins
}

func execute(instruction string) int {
	r := regexp.MustCompile(`\d+`)

	numStrings := r.FindAllString(instruction, -1)

	if len(numStrings) != 2 {
		return 0
	}

	ints := []int{}

	for _, s := range numStrings {
		if len(s) > 3 {
			return 0
		}
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}

	return ints[0] * ints[1]
}

func part1(data string) {
	ins := partOneInstructions(data)
	total := 0

	for _, i := range ins {
		total += execute(i)
	}

	fmt.Println(total)
}

func part2(data string) {
	ins := partTwoInstructions(data)
	total := 0
	do := true

	for _, i := range ins {
		switch i.matchType {
		case DO:
			do = true
		case DONT:
			do = false
		case MUL:
			if do {
				total += execute(i.string)
			}
		}
	}

	println(total)
}

func main() {
	input := getInput("input.txt")
	// input := getInput("sample_input_2.txt")

	// part1(input)
	part2(input)
}
