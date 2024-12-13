package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// lines := getInput("example_input.txt")
	// lines := getInput("example_input2.txt")
	// lines := getInput("example_input3.txt")
	// lines := getInput("example_input4.txt")
	// lines := getInput("example_input5.txt")
	lines := getInput("input.txt")

	solve(lines, part2)
}

type coord struct {
	x, y int
}

func solve(lines []string, f func([]coord, []string) int) {
	seen := make(map[coord]bool)
	plots := [][]coord{}

	for y, line := range lines {
		for x := range line {
			if _, ok := seen[coord{x, y}]; ok {
				continue
			}

			plot := getNeighbours(x, y, lines, seen)

			plots = append(plots, plot)
		}
	}

	fenceCost := 0
	for _, p := range plots {
		perimeter := f(p, lines)

		fenceCost += perimeter * len(p)
	}

	fmt.Println(fenceCost)
}

func part1(plot []coord, lines []string) int {
	perimeter := 0

	for _, c := range plot {
		p := 4

		neighbours := []coord{
			{c.x - 1, c.y},
			{c.x + 1, c.y},
			{c.x, c.y - 1},
			{c.x, c.y + 1},
		}

		for _, n := range neighbours {
			for _, pp := range plot {
				if n.x == pp.x && n.y == pp.y {
					p--
				}
			}
		}
		perimeter += p
	}

	return perimeter
}

func part2(plot []coord, lines []string) int {
	xMin, xMax, yMin, yMax := getFullArea(plot)
	sides := 0

	for x := xMin; x <= xMax; x++ {
		sides += sidesAtX(x, plot, lines)
	}

	for y := yMin; y <= yMax; y++ {
		sides += sidesAtY(y, plot, lines)
	}

	return sides
}

func sidesAtX(x int, plot []coord, lines []string) int {
	left := []coord{}
	right := []coord{}
	val := lines[plot[0].y][plot[0].x]

	for _, c := range plot {
		if c.x != x {
			continue
		}

		if x-1 < 0 || lines[c.y][x-1] != val {
			left = append(left, coord{
				x - 1, c.y,
			})
		}
		if x+1 >= len(lines[0]) || lines[c.y][x+1] != val {
			right = append(right, coord{
				x + 1, c.y,
			})
		}
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i].y < left[j].y
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i].y < right[j].y
	})

	sides := 0
	for _, l := range [][]coord{left, right} {
		last := -3

		for _, c := range l {
			if c.y > last+1 {
				sides++
			}

			last = c.y
		}
	}

	return sides
}

func sidesAtY(y int, plot []coord, lines []string) int {
	top := []coord{}
	bottom := []coord{}
	val := lines[plot[0].y][plot[0].x]

	for _, c := range plot {
		if c.y != y {
			continue
		}

		if y-1 < 0 || lines[y-1][c.x] != val {
			top = append(top, coord{
				c.x, y - 1,
			})
		}
		if y+1 >= len(lines[0]) || lines[y+1][c.x] != val {
			bottom = append(bottom, coord{
				c.x, y + 1,
			})
		}
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].x < top[j].x
	})

	sort.Slice(bottom, func(i, j int) bool {
		return bottom[i].x < bottom[j].x
	})

	sides := 0
	for _, l := range [][]coord{top, bottom} {
		last := -3

		for _, c := range l {
			if c.x > last+1 {
				sides++
			}

			last = c.x
		}
	}

	return sides
}

func getFullArea(plot []coord) (int, int, int, int) {
	xMin := plot[0].x
	xMax := plot[0].x
	yMin := plot[0].y
	yMax := plot[0].y

	for _, c := range plot {
		if c.x < xMin {
			xMin = c.x
		} else if c.x > xMax {
			xMax = c.x
		}

		if c.y < yMin {
			yMin = c.y
		} else if c.y > yMax {
			yMax = c.y
		}
	}

	return xMin, xMax, yMin, yMax
}

func getNeighbours(x, y int, lines []string, seen map[coord]bool) []coord {
	plot := []coord{
		{x, y},
	}
	seen[coord{x, y}] = true

	c := lines[y][x]

	neighbours := []coord{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}

	for _, n := range neighbours {
		if n.x < 0 || n.x > len(lines[0])-1 ||
			n.y < 0 || n.y > len(lines)-1 {
			continue
		}

		if _, ok := seen[n]; ok {
			continue
		}

		if lines[n.y][n.x] != c {
			continue
		}

		seen[n] = true
		ns := getNeighbours(n.x, n.y, lines, seen)
		plot = append(plot, ns...)
	}

	return plot
}

func getInput(filename string) []string {
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
