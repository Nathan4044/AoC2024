package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	maxX = 101
	maxY = 103
	// for example
	// maxX = 11
	// maxY = 7
	moves = 100
)

type coord struct {
	x, y int
}

type robot struct {
	position coord
	velocity coord
}

func (r *robot) move(n int) {
	r.position.x = mod((r.position.x + (r.velocity.x * n)), maxX)
	r.position.y = mod((r.position.y + (r.velocity.y * n)), maxY)
}

func main() {
	// robots := getRobots("example_input.txt")
	robots := getRobots("input.txt")
	// part1(robots)
	part2(robots)
}

func part1(robots []robot) {
	upperLeft := 0
	upperRight := 0
	lowerLeft := 0
	lowerRight := 0

	centreX := maxX / 2
	centreY := maxY / 2

	for _, r := range robots {
		r.move(moves)

		if r.position.x > centreX {
			if r.position.y > centreY {
				lowerRight++
			} else if r.position.y < centreY {
				upperRight++
			}
		} else if r.position.x < centreX {
			if r.position.y > centreY {
				lowerLeft++
			} else if r.position.y < centreY {
				upperLeft++
			}
		}
	}

	fmt.Println(upperLeft * upperRight * lowerLeft * lowerRight)
}

func part2(robots []robot) {
	for i := 0; ; i++ {
		for i, _ := range robots {
			robots[i].move(1)
		}

		fmt.Printf("%d\n", i+1)
		printArea(robots)
	}
}

func printArea(robots []robot) {
	scanner := bufio.NewScanner(os.Stdin)
	possibleTree := false
	buf := bytes.Buffer{}
	lines := [maxY][maxX]int{}

	for _, r := range robots {
		lines[r.position.y][r.position.x]++
	}

	for y, l := range lines {
		for x, n := range l {
			if n > 0 {
				buf.WriteString(fmt.Sprintf("%d", n))

				if y+2 < maxY && x+2 < maxX && x-2 >= 0 &&
					lines[y+1][x] > 0 &&
					lines[y+1][x+1] > 0 &&
					lines[y+1][x-1] > 0 &&
					lines[y+2][x] > 0 &&
					lines[y+2][x+1] > 0 &&
					lines[y+2][x+2] > 0 &&
					lines[y+2][x-1] > 0 &&
					lines[y+2][x-2] > 0 {
					possibleTree = true
				}
			} else {
				buf.WriteString(fmt.Sprintf("."))
			}
		}
		buf.WriteString("\n")
	}

	if possibleTree {
		fmt.Println(buf.String())
		scanner.Scan()
	}
}

func getRobots(filename string) []robot {
	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`-?\d+`)
	scanner := bufio.NewScanner(f)
	robots := []robot{}

	for scanner.Scan() {
		rob := robot{}
		ns := r.FindAllString(scanner.Text(), -1)

		n, _ := strconv.Atoi(ns[0])
		rob.position.x = n

		n, _ = strconv.Atoi(ns[1])
		rob.position.y = n

		n, _ = strconv.Atoi(ns[2])
		rob.velocity.x = n

		n, _ = strconv.Atoi(ns[3])
		rob.velocity.y = n

		robots = append(robots, rob)
	}

	return robots
}

func mod(a, b int) int {
	return (a%b + b) % b
}
