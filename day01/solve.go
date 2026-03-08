package day01

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func PyMod(a, b int) int {
	return (a%b + b) % b
}

func ProcessData(input string) []int {
	output := make([]int, 0)
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		direction := 1
		if line[0] == 'L' {
			direction = -1
		}
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println(err)
			fmt.Println(line)
			panic(err)
		}
		output = append(output, num*direction)

	}
	return output
}

func SolvePuzzle1(input string) int {
	// TODO: solve puzzle 1
	data := ProcessData(input)
	position := 50
	total := 0

	for _, d := range data {
		position += d
		if position < 0 {
			position %= -100
		} else if position > 99 {
			position %= 100
		}
		if position == 0 {
			total++
		}
	}

	return total
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	data := ProcessData(input)
	position := 50
	total := 0
	increase := 0

	for _, d := range data {
		fmt.Println(position)
		if d < 0 {
			increase = (PyMod(-position, 100) - d) / 100
		} else {
			increase = (position + d) / 100
		}
		total += increase
		position += d
		if position < 0 {
			position = PyMod(position, 100)
		} else if position > 99 {
			position %= 100
		}
		fmt.Println(position)
		fmt.Println("Total:", total)
	}

	return total
}
