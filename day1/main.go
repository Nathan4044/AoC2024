package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read '%s': %s\n", filename, err)
	}

	scanner := bufio.NewScanner(file)

	var left, right []int

	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "   ")

		num, _ := strconv.Atoi(items[0])
		left = append(left, num)

		num, _ = strconv.Atoi(items[1])
		right = append(right, num)
	}

	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}

func partOne(left, right []int) {
	diffs := 0

	for i, n := range left {
		diffs += abs(n - right[i])
	}

	fmt.Printf("%d\n", diffs)
}

func partTwo(left, right []int) {
	rightIndex := 0
	result := 0
	matches := 0
	lastLeft := -1

	for _, n := range left {
		if n == lastLeft {
			result += n * matches
			continue
		}

		matches = 0

		if rightIndex > len(right)-1 {
			break
		}

		for right[rightIndex] <= n {
			if right[rightIndex] == n {
				matches++
			}

			rightIndex++

			if rightIndex > len(right)-1 {
				break
			}
		}

		result += n * matches
		lastLeft = n
	}

	fmt.Println(result)
}

func main() {
	left, right := readInput("input.txt")

	partOne(left, right)
	partTwo(left, right)
}
