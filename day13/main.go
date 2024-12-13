package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	x, y int
}

type Prize Button

type Game struct {
	buttonA, buttonB Button
	prize            Prize
}

func main() {
	// games := getGames("example_input.txt")
	games := getGames("input.txt")

	// part1(games)
	part2(games)
}

func part1(games []Game) {
	var total int64 = 0
	for _, g := range games {
		result, ok := solveGame(g)

		if ok {
			total += result
		}
	}

	fmt.Println(total)
}

func part2(games []Game) {
	var total int64 = 0
	for _, g := range games {
		g.prize.x += 10000000000000
		g.prize.y += 10000000000000
		result, ok := solveGame(g)

		if ok {
			total += result
		}
	}

	fmt.Println(total)
}

// func solveGame(game Game) (solution int, solved bool) {
//     solution = int(^uint(0) >> 1) // maximum int value
//     maxB := game.prize.x / game.buttonB.x
//     maxBsY := game.prize.y / game.buttonB.y
//
//     if maxBsY < maxB {
//         maxB = maxBsY
//     }
//
//     fmt.Printf("B button permutations: %d\n", maxB)
//
//     for i := 0; i <= maxB; i++ {
//         remainingX := game.prize.x - (game.buttonB.x * i)
//         if remainingX % game.buttonA.x != 0 {
//             continue
//         }
//
//         remainingY := game.prize.y - (game.buttonB.y * i)
//         if remainingY % game.buttonA.y != 0 {
//             continue
//         }
//
//         x := remainingX / game.buttonA.x
//         y := remainingY / game.buttonA.y
//         if x != y {
//             continue
//         }
//
//         val := (3 * x) + i
//
//         if val < solution {
//             solution = val
//             solved = true
//         }
//     }
//
//     return
// }

func solveGame(game Game) (int64, bool) {
	fmt.Println(game)
	a := float64(game.buttonA.x)
	b := float64(game.buttonA.y)
	c := float64(game.buttonB.x)
	d := float64(game.buttonB.y)
	e := float64(game.prize.x)
	f := float64(game.prize.y)

	n := (d*e - c*f) / (d*a - c*b)
	fmt.Printf("n = %g\n", n)

	if float64(int64(n)) == n {
		m := int64((f - n*b) / d)
		fmt.Printf("solution: n = %d, m = %d\n\n", int64(n), int64(m))
		return 3*int64(n) + m, true
	}

	fmt.Printf("no solution\n\n")
	return 0, false
}

func getGames(filename string) []Game {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(f), "\n\n")
	games := []Game{}

	for _, g := range input {
		games = append(games, parseGame(g))
	}

	return games
}

func parseGame(s string) Game {
	lines := strings.Split(s, "\n")
	var g Game

	g.buttonA = parseButton(lines[0])
	g.buttonB = parseButton(lines[1])
	g.prize = Prize(parseButton(lines[2]))

	return g
}

func parseButton(s string) Button {
	r := regexp.MustCompile(`\d+`)
	var b Button

	ns := r.FindAllString(s, -1)

	n, _ := strconv.Atoi(ns[0])
	b.x = n
	n, _ = strconv.Atoi(ns[1])
	b.y = n

	return b
}
