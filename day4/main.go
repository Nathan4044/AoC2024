package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// lines := getLines("example_input.txt")
	lines := getLines("input.txt")

	// part1(lines)
	part2(lines)
}

func getLines(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func xmasInstances(x, y int, lines []string) int {
	if lines[y][x] != 'X' {
		return 0
	}

	count := 0
	inLeft := x-3 >= 0
	inRight := x+3 < len(lines[0])
	inTop := y-3 >= 0
	inBottom := y+3 < len(lines)

	if inRight &&
		lines[y][x+1:x+4] == "MAS" {
		count++
	}

	if inLeft &&
		lines[y][x-3:x] == "SAM" {
		count++
	}

	if inBottom &&
		string([]byte{lines[y+1][x], lines[y+2][x], lines[y+3][x]}) == "MAS" {
		count++
	}

	if inTop &&
		string([]byte{lines[y-1][x], lines[y-2][x], lines[y-3][x]}) == "MAS" {
		count++
	}

	if inLeft && inTop &&
		string([]byte{lines[y-1][x-1], lines[y-2][x-2], lines[y-3][x-3]}) == "MAS" {
		count++
	}

	if inRight && inTop &&
		string([]byte{lines[y-1][x+1], lines[y-2][x+2], lines[y-3][x+3]}) == "MAS" {
		count++
	}

	if inLeft && inBottom &&
		string([]byte{lines[y+1][x-1], lines[y+2][x-2], lines[y+3][x-3]}) == "MAS" {
		count++
	}

	if inRight && inBottom &&
		string([]byte{lines[y+1][x+1], lines[y+2][x+2], lines[y+3][x+3]}) == "MAS" {
		count++
	}

	return count
}

func crossMasInstance(x, y int, lines []string) bool {
	inLeft := x-1 >= 0
	inRight := x+1 < len(lines[0])
	inTop := y-1 >= 0
	inBottom := y+1 < len(lines)

	if !(inLeft && inRight && inTop && inBottom) {
		return false
	}

	if lines[y][x] != 'A' {
		return false
	}

	diag1 := string([]byte{lines[y-1][x-1], lines[y+1][x+1]})
	diag2 := string([]byte{lines[y-1][x+1], lines[y+1][x-1]})

	if (diag1 == "MS" || diag1 == "SM") &&
		(diag2 == "MS" || diag2 == "SM") {
		return true
	}

	return false
}

func part1(lines []string) {
	count := 0
	for y, line := range lines {
		for x := range line {
			count += xmasInstances(x, y, lines)
		}
	}

	fmt.Println(count)
}

func part2(lines []string) {
	count := 0
	for y, line := range lines {
		for x := range line {
			if crossMasInstance(x, y, lines) {
				count++
			}
		}
	}

	fmt.Println(count)
}
