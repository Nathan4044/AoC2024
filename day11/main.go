package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// stones := getStones("example_input.txt")
	stones := getStones("input.txt")

	// solve(stones, 25)
	solve(stones, 75)
}

func solve(stones []int, iterations int) {
	stoneCounts := map[int]int{}

	for _, s := range stones {
		stoneCounts[s]++
	}

	for i := 0; i < iterations; i++ {
		stoneCounts = iterCounts(stoneCounts)
	}

	count := 0
	for _, v := range stoneCounts {
		count += v
	}
	fmt.Println(count)
}

func iterCounts(counts map[int]int) map[int]int {
	newCount := map[int]int{}

	for k, v := range counts {
		if k == 0 {
			newCount[1] += v
		} else if len(digits(k))%2 == 0 {
			for _, s := range splitStone(k) {
				newCount[s] += v
			}
		} else {
			newCount[k*2024] += v
		}
	}

	return newCount
}

func digits(i int) string {
	return fmt.Sprintf("%d", i)
}

func splitStone(s int) []int {
	d := digits(s)
	half := len(d) / 2

	a := d[:half]
	b := d[half:]

	first, _ := strconv.Atoi(a)
	second, _ := strconv.Atoi(b)

	return []int{first, second}
}

func getStones(filename string) []int {
	f, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	numStrings := strings.Split(string(f[:len(f)-1]), " ")

	nums := []int{}

	for _, s := range numStrings {
		n, err := strconv.Atoi(s)

		if err != nil {
			panic(err)
		}

		nums = append(nums, n)
	}

	return nums
}
