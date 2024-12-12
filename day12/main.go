package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    // lines := getInput("example_input.txt")
    // lines := getInput("example_input2.txt")
    // lines := getInput("example_input3.txt")
    lines := getInput("input.txt")

    solve(lines)
}

type coord struct {
    x, y int
}

func solve(lines []string) {
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
        perimeter := part1(p)

        fenceCost += perimeter * len(p)
    }

    fmt.Println(fenceCost)
}

func part1(plot []coord) int {
    perimeter := 0

    for _, c := range plot {
        p := 4

        neighbours := []coord{
            { c.x-1, c.y },
            { c.x+1, c.y },
            { c.x, c.y-1 },
            { c.x, c.y+1 },
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

func part2(plot []coord) int {
    // todo
    return 0
}

func getNeighbours(x, y int, lines []string, seen map[coord]bool) []coord {
    plot := []coord{
        { x, y },
    }
    seen[coord{x, y}] = true

    c := lines[y][x]

    neighbours := []coord{
        { x-1, y },
        { x+1, y },
        { x, y-1 },
        { x, y+1 },
    }

    for _, n := range neighbours {
        if n.x < 0 || n.x > len(lines[0]) - 1 ||
        n.y < 0 || n.y > len(lines) - 1 {
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
