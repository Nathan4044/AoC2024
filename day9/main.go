package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// data := getPartOneData("example_input.txt")
	// data := getPartOneData("input.txt")
	// partOne(data)

	// data := getPartTwoData("example_input.txt")
	data := getPartTwoData("input.txt")
	partTwo(data)
}

func partOne(data []rune) {
	back := len(data) - 1

	for i := 0; i < back; i++ {
		if data[i] != '.' {
			continue
		}

		for j := back; j > i; j-- {
			if data[j] == '.' {
				continue
			}

			data[i] = data[j]
			data[j] = '.'
			back = j
			break
		}
	}

	total := 0
	for i, c := range data {
		if c == '.' {
			break
		}

		total += i * int(c-'0')
	}

	fmt.Println(total)
}

func partTwo(blocks []block) {
	first := blocks[0].char

	for i := len(blocks) - 1; blocks[i].char != first; i-- {
		if blocks[i].char == '.' {
			continue
		}

		l := blocks[i].len

		for j := 0; j < i; j++ {
			if !blocks[j].empty || blocks[j].len < l {
				continue
			}

			if blocks[j].len == l {
				blocks[j].char = blocks[i].char
				blocks[j].empty = false
				blocks[i].char = '.'
				blocks[i].empty = true
				break
			}

			blocks = append(blocks, block{
				char:  '.',
				len:   blocks[j].len - blocks[i].len,
				empty: true,
			})

			blocks[j].char = blocks[i].char
			blocks[j].empty = false
			blocks[j].len = blocks[i].len
			blocks[i].char = '.'
			blocks[i].empty = true

			for k := len(blocks) - 1; k > j+1; k-- {
				temp := blocks[k]

				blocks[k] = blocks[k-1]
				blocks[k-1] = temp
			}
			i++
			break
		}
	}

	count := 0
	index := 0

	for _, b := range blocks {
		if b.empty {
			index += b.len
			continue
		}

		for i := 0; i < b.len; i++ {
			count += index * int(b.char-'0')
			index++
		}
	}

	fmt.Println(count)
}

func getPartOneData(filename string) []rune {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer([]byte{})

	for i, b := range string(file[:len(file)-1]) {
		bn := int(b - '0')
		char := '.'

		if i%2 == 0 {
			char = rune('0' + (i / 2))
		}

		for n := 0; n < bn; n++ {
			buf.WriteRune(char)
		}
	}

	return []rune(buf.String())
}

type block struct {
	len   int
	char  rune
	empty bool
}

func getPartTwoData(filename string) []block {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	blocks := []block{}

	for i, b := range string(file[:len(file)-1]) {
		bn := int(b - '0')
		char := '.'
		empty := true

		if i%2 == 0 {
			char = rune('0' + (i / 2))
			empty = false
		}

		blocks = append(
			blocks,
			block{
				len:   bn,
				char:  char,
				empty: empty,
			},
		)
	}

	return blocks
}
