package main

import (
	"bufio"
	"fmt"
	"os"
)

type coords struct {
	x, y int
}

func main() {
	// trail := readTrail("example_input.txt")
	trail := readTrail("input.txt")

	// part1(trail)
	part2(trail)
}

func part1(trail [][]byte) {
	count := 0
	for y, line := range trail {
		for x, n := range line {
			if n == 0 {
				count += len(getPaths(x, y, n, trail))
			}
		}
	}
	fmt.Println(count)
}

func part2(trail [][]byte) {
	count := 0
	for y, line := range trail {
		for x, n := range line {
			if n == 0 {
				count += getPaths2(x, y, n, trail)
			}
		}
	}
	fmt.Println(count)
}

func getPaths(x, y int, val byte, trail [][]byte) map[coords]bool {
	endPoints := map[coords]bool{}

	opts := []coords{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}

	for _, o := range opts {
		if !(o.x >= 0 && o.x < len(trail[0]) &&
			o.y >= 0 && o.y < len(trail) &&
			trail[o.y][o.x] == val+1) {
			continue
		}

		if trail[o.y][o.x] == 9 {
			endPoints[o] = true
			continue
		}

		for c := range getPaths(o.x, o.y, val+1, trail) {
			endPoints[c] = true
		}
	}

	return endPoints
}

func getPaths2(x, y int, val byte, trail [][]byte) int {
	count := 0

	opts := []coords{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}

	for _, o := range opts {
		if !(o.x >= 0 && o.x < len(trail[0]) &&
			o.y >= 0 && o.y < len(trail) &&
			trail[o.y][o.x] == val+1) {
			continue
		}

		if trail[o.y][o.x] == 9 {
			count++
			continue
		}

		count += getPaths2(o.x, o.y, val+1, trail)
	}

	return count
}

func readTrail(filename string) [][]byte {
	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	trail := [][]byte{}

	for scanner.Scan() {
		line := []byte{}

		for _, c := range scanner.Bytes() {
			line = append(line, c-'0')
		}

		trail = append(trail, line)
	}

	return trail
}
