package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type report []int

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func safePair(one, two int, shouldAscend bool) bool {
	diff := two - one
	ascending := diff > 0

	if ascending != shouldAscend {
		return false
	}

	diff = abs(diff)
	if diff < 1 || diff > 3 {
		return false
	}

	return true
}

func (r report) isSafe() bool {
	if abs(r[0]-r[1]) < 1 || abs(r[0]-r[1]) > 3 {
		return false
	}

	startsAscending := (r[1] - r[0]) > 0

	for i := 1; i < len(r)-1; i++ {
		if !safePair(r[i], r[i+1], startsAscending) {
			return false
		}
	}

	return true
}

func (r report) isSafeDampened() bool {
	initial := 1
	dampened := false

	if abs(r[0]-r[1]) < 1 || abs(r[0]-r[1]) > 3 {
		return r[1:].isSafe() || append(report{r[0]}, r[2:]...).isSafe()
	}

	startsAscending := (r[1] - r[0]) > 0

	if ((r[2] - r[1]) > 0) != startsAscending {
		if r[1:].isSafe() {
			return true
		} else {
			if abs(r[0]-r[2]) < 1 || abs(r[0]-r[2]) > 3 {
				return false
			}

			initial = 2
			dampened = true
			startsAscending = (r[2] - r[0]) > 0
		}
	}

	for i := initial; i < len(r)-1; i++ {
		if safePair(r[i], r[i+1], startsAscending) {
			continue
		}

		if dampened {
			return false
		}

		if i+2 > len(r)-1 {
			continue
		}

		dampened = true

		if !safePair(r[i], r[i+2], startsAscending) {
			return false
		}

		i++
	}

	return true
}

func getReports(filename string) []report {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open '%s': %s\n", filename, err)
		os.Exit(1)
	}

	var reports []report
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var r report

		for _, num := range line {
			n, err := strconv.Atoi(num)

			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to parse %s as number", num)
				os.Exit(2)
			}
			r = append(r, n)
		}

		reports = append(reports, r)
	}

	return reports
}

func part1(reports []report) {
	safeCount := 0

	for _, r := range reports {
		if r.isSafe() {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func part2(reports []report) {
	safeCount := 0

	for _, r := range reports {
		if r.isSafeDampened() {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func main() {
	// reports := getReports("sample_input.txt")
	reports := getReports("input.txt")

	// part1(reports)
	part2(reports)
}
