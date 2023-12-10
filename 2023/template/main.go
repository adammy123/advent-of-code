package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	// fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	result := 0
	rawRows := strings.Split(input, "\n")

	for _, row := range rawRows {
		fmt.Println(row)
	}

	return result
}

func part2(input string) int {
	result := 0

	return result
}
