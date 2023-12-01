package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./1.txt")
	if err != nil {
		log.Fatalln("error reading file: ", err)
	}
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	elves := [][]int{}
	for _, group := range strings.Split(input, "\n\n") {
		row := []int{}
		for _, line := range strings.Split(group, "\n") {
			num, _ := strconv.Atoi(line)
			row = append(row, num)
		}
		elves = append(elves, row)
	}

	maxSum := 0
	for _, elf := range elves {
		currSum := 0
		for _, score := range elf {
			currSum += score
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum

}

func part2(input string) int {
	elves := [][]int{}
	for _, group := range strings.Split(input, "\n\n") {
		row := []int{}
		for _, line := range strings.Split(group, "\n") {
			num, _ := strconv.Atoi(line)
			row = append(row, num)
		}
		elves = append(elves, row)
	}

	first := 0
	second := 0
	third := 0
	for _, elf := range elves {
		currSum := 0
		for _, score := range elf {
			currSum += score
		}

		if currSum > first {
			third = second
			second = first
			first = currSum
		} else if currSum > second {
			third = second
			second = currSum
		} else if currSum > third {
			third = currSum
		}
	}

	return first + second + third

}
