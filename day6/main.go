package main

import (
	"bytes"
	"fmt"
	"os"
)

var (
	guardDirs = []byte{'^', '>', 'v', '<'}
)

func main() {
	// m := getMap("example_input.txt")
	m := getMap("input.txt")

	// part1(m)
	part2(m)
}

func part1(m [][]byte) {
	x, y, err := startPos(m)
	if err != nil {
		panic(err)
	}

	x, y, m, finished := moveGuard(x, y, m)
	for !finished {
		x, y, m, finished = moveGuard(x, y, m)
	}

	fmt.Println(stepCount(m))
}

func part2(m [][]byte) {
	count := 0
	m2 := make([][]byte, len(m))
	for i := range m2 {
		m2[i] = make([]byte, len(m[0]))
	}

	for i := range m {
		for j := range m[i] {
			if m[i][j] == '.' {
				for y := range m {
					for x := range m[y] {
						m2[y][x] = m[y][x]
					}
				}

				m2[i][j] = '#'

				if hasLoop(m2) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func hasLoop(m [][]byte) bool {
	x, y, err := startPos(m)

	if err != nil {
		panic("no guard")
	}

	path := make([][][]byte, len(m))
	for i := range path {
		path[i] = make([][]byte, len(m[0]))
	}

	x, y, m, path, finished, loop := trackGuard(x, y, m, path)
	for !finished {
		x, y, m, path, finished, loop = trackGuard(x, y, m, path)
	}

	return loop
}

func moveGuard(x, y int, m [][]byte) (int, int, [][]byte, bool) {
	guard := m[y][x]
	switch guard {
	case '^':
		if y-1 < 0 {
			m[y][x] = 'X'
			return 0, 0, m, true
		}

		next := m[y-1][x]

		if next == '#' {
			m[y][x] = '>'
			return x, y, m, false
		} else {
			m[y-1][x] = '^'
			m[y][x] = 'X'
			return x, y - 1, m, false
		}
	case '>':
		if x+1 >= len(m[0]) {
			m[y][x] = 'X'
			return 0, 0, m, true
		}

		next := m[y][x+1]

		if next == '#' {
			m[y][x] = 'v'
			return x, y, m, false
		} else {
			m[y][x+1] = '>'
			m[y][x] = 'X'
			return x + 1, y, m, false
		}
	case 'v':
		if y+1 >= len(m) {
			m[y][x] = 'X'
			return 0, 0, m, true
		}

		next := m[y+1][x]

		if next == '#' {
			m[y][x] = '<'
			return x, y, m, false
		} else {
			m[y+1][x] = 'v'
			m[y][x] = 'X'
			return x, y + 1, m, false
		}
	case '<':
		if x-1 < 0 {
			m[y][x] = 'X'
			return 0, 0, m, true
		}

		next := m[y][x-1]

		if next == '#' {
			m[y][x] = '^'
			return x, y, m, false
		} else {
			m[y][x-1] = '<'
			m[y][x] = 'X'
			return x - 1, y, m, false
		}
	}
	panic("invalid guard")
}

func trackGuard(x, y int, m [][]byte, path [][][]byte) (int, int, [][]byte, [][][]byte, bool, bool) {
	guard := m[y][x]

	for _, g := range path[y][x] {
		if g == guard {
			return 0, 0, m, path, true, true
		}
	}

	path[y][x] = append(path[y][x], guard)

	switch guard {
	case '^':
		if y-1 < 0 {
			m[y][x] = 'X'
			return 0, 0, m, path, true, false
		}

		next := m[y-1][x]

		if next == '#' {
			m[y][x] = '>'
			return x, y, m, path, false, false
		} else {
			m[y-1][x] = '^'
			m[y][x] = 'X'
			return x, y - 1, m, path, false, false
		}
	case '>':
		if x+1 >= len(m[0]) {
			m[y][x] = 'X'
			return 0, 0, m, path, true, false
		}

		next := m[y][x+1]

		if next == '#' {
			m[y][x] = 'v'
			return x, y, m, path, false, false
		} else {
			m[y][x+1] = '>'
			m[y][x] = 'X'
			return x + 1, y, m, path, false, false
		}
	case 'v':
		if y+1 >= len(m) {
			m[y][x] = 'X'
			return 0, 0, m, path, true, false
		}

		next := m[y+1][x]

		if next == '#' {
			m[y][x] = '<'
			return x, y, m, path, false, false
		} else {
			m[y+1][x] = 'v'
			m[y][x] = 'X'
			return x, y + 1, m, path, false, false
		}
	case '<':
		if x-1 < 0 {
			m[y][x] = 'X'
			return 0, 0, m, path, true, false
		}

		next := m[y][x-1]

		if next == '#' {
			m[y][x] = '^'
			return x, y, m, path, false, false
		} else {
			m[y][x-1] = '<'
			m[y][x] = 'X'
			return x - 1, y, m, path, false, false
		}
	}
	panic("invalid guard")
}

func stepCount(m [][]byte) int {
	count := 0

	for _, line := range m {
		for _, b := range line {
			if b == 'X' {
				count++
			}
		}
	}

	return count
}

func startPos(m [][]byte) (int, int, error) {
	for j, line := range m {
		for i, c := range line {
			if isGuard(c) {
				return i, j, nil
			}
		}
	}

	return 0, 0, fmt.Errorf("no guard")
}

func isGuard(c byte) bool {
	for _, r := range guardDirs {
		if c == r {
			return true
		}
	}

	return false
}

func getMap(filename string) [][]byte {
	b, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	lines := bytes.Split(b, []byte{'\n'})
	return lines[:len(lines)-1]
}
