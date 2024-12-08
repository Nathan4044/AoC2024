package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type equation struct {
    result int
    operands []int
}

func main() {
    // data := getData("example_input.txt")
    data := getData("input.txt")

    // solve(data, part1)
    solve(data, part2)
}

func solve(equations []equation, f func(int, []int) bool) {
    result := 0

    for _, e := range equations {
        if f(e.result, e.operands) {
            result += e.result
        }
    }

    fmt.Println(result)
}

func part1(result int, operands []int) bool {
    switch len(operands) {
    case 2:
        if operands[0] + operands[1] == result {
            return true
        }
        if operands[0] * operands[1] == result {
            return true
        }
        return false
    default:
        n1 := operands[0] + operands[1]
        n2 := operands[0] * operands[1]

        ops := append([]int{n1}, operands[2:]...)

        if part1(result, ops) {
            return true
        }

        ops = append([]int{n2}, operands[2:]...)

        if part1(result, ops) {
            return true
        }

        return false
    }
}

func part2(result int, operands []int) bool {
    switch len(operands) {
    case 2:
        if operands[0] + operands[1] == result {
            return true
        }
        if operands[0] * operands[1] == result {
            return true
        }
        if concat(operands[0], operands[1]) == result {
            return true
        }
        return false
    default:
        n1 := operands[0] + operands[1]
        n2 := operands[0] * operands[1]
        n3 := concat(operands[0], operands[1])

        ops := append([]int{n1}, operands[2:]...)

        if part2(result, ops) {
            return true
        }

        ops = append([]int{n2}, operands[2:]...)

        if part2(result, ops) {
            return true
        }

        ops = append([]int{n3}, operands[2:]...)

        if part2(result, ops) {
            return true
        }

        return false
    }
}

func concat(a, b int) int {
    c, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
    return c
}

func getData(filename string) []equation {
    f, err := os.Open(filename)
    defer f.Close()

    if err != nil {
        panic(err)
    }

    r := regexp.MustCompile(`\d+`)
    scanner := bufio.NewScanner(f)
    equations := []equation{}

    for scanner.Scan() {
        numStrings := r.FindAllString(scanner.Text(), -1)

        operands := []int{}

        for _, s := range numStrings[1:] {
            n, err := strconv.Atoi(s)

            if err != nil {
                panic(err)
            }

            operands = append(operands, n)
        }

        res, err := strconv.Atoi(numStrings[0])

        if err != nil {
            panic(err)
        }

        equations = append(equations, equation{
            result: res,
            operands: operands,
        })
    }

    return equations
}
