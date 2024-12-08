package main

import (
	"bufio"
	"os"
)

func main() {
	// data := getData("example_input.txt")
	data := getData("input.txt")

	// solve(data, part1)
	solve(data, part2)
}

type coord struct {
	x, y int
}

func (c *coord) in(h, w int) bool {
	return c.x >= 0 &&
		c.x < w &&
		c.y >= 0 &&
		c.y < h
}

func solve(data []string, f func([]coord, int, int) []coord) {
	height := len(data)
	width := len(data[0])

	antennae := make(map[rune][]coord)

	for y, line := range data {
		for x, c := range line {
			if c == '.' {
				continue
			}

			if _, ok := antennae[c]; !ok {
				antennae[c] = []coord{}
			}

			antennae[c] = append(antennae[c], coord{
				x: x,
				y: y,
			})
		}
	}

	unique := make([][]bool, height)
	count := 0
	for i := range unique {
		unique[i] = make([]bool, width)
	}

	for _, coords := range antennae {
		points := f(coords, height, width)

		for _, p := range points {
			if unique[p.y][p.x] != true {
				count++
			}
			unique[p.y][p.x] = true
		}
	}

	println(count)
}

func part1(cs []coord, height, width int) []coord {
	result := []coord{}

	for i, c1 := range cs {
		for _, c2 := range cs[i+1:] {
			xDiff := c1.x - c2.x
			yDiff := c1.y - c2.y

			newCoords := []coord{
				{
					x: c1.x + xDiff,
					y: c1.y + yDiff,
				},
				{
					x: c2.x - xDiff,
					y: c2.y - yDiff,
				},
			}

			for _, c := range newCoords {
				if c.in(height, width) {
					result = append(result, c)
				}
			}
		}
	}

	return result
}

func part2(cs []coord, height, width int) []coord {
	result := []coord{}

	for i, c1 := range cs {
		if len(cs) > 1 {
			result = append(result, c1)
		}

		for _, c2 := range cs[i+1:] {
			xDiff := c1.x - c2.x
			yDiff := c1.y - c2.y

			nc := coord{
				x: c1.x + xDiff,
				y: c1.y + yDiff,
			}
			for nc.in(height, width) {
				result = append(result, nc)

				nc = coord{
					x: nc.x + xDiff,
					y: nc.y + yDiff,
				}
			}

			nc = coord{
				x: c2.x - xDiff,
				y: c2.y - yDiff,
			}
			for nc.in(height, width) {
				result = append(result, nc)

				nc = coord{
					x: nc.x - xDiff,
					y: nc.y - yDiff,
				}
			}
		}
	}

	return result
}

func getData(filename string) []string {
	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
