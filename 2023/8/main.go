package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	left string
	right string
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	totalSteps := 0
	rows := strings.Split(input, "\n")
	instructions := rows[0]
	nodesStr := rows[2:]
	nodesMap := map[string]node{}

	for _, nodeStr := range nodesStr {
		nodeStr = strings.ReplaceAll(nodeStr, " ", "")
		parts := strings.Split(nodeStr, "=")
		key := parts[0]
		leftRightParts := strings.Split(parts[1], ",")
		left := strings.ReplaceAll(leftRightParts[0], "(", "")
		right := strings.ReplaceAll(leftRightParts[1], ")", "")
		nodesMap[key] = node{left: left, right: right}
	}

	start := "AAA"

	for {
		for _, instruction := range strings.Split(instructions, "") {
			totalSteps += 1
			if instruction == "L" {
				start = nodesMap[start].left
			} else {
				start = nodesMap[start].right
			}
			if start == "ZZZ" {
				return totalSteps
			}
		}
	}
}

func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}

func part2(input string) int {
	rows := strings.Split(input, "\n")
	instructions := rows[0]
	nodesStr := rows[2:]
	nodesMap := map[string]node{}

	for _, nodeStr := range nodesStr {
		nodeStr = strings.ReplaceAll(nodeStr, " ", "")
		parts := strings.Split(nodeStr, "=")
		key := parts[0]
		leftRightParts := strings.Split(parts[1], ",")
		left := strings.ReplaceAll(leftRightParts[0], "(", "")
		right := strings.ReplaceAll(leftRightParts[1], ")", "")
		nodesMap[key] = node{left: left, right: right}
	}

	starts := []string{}
	minSteps := []int{}

	for k := range nodesMap {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}

	fmt.Println(starts)
	StartLoop:
	for _, start := range starts {
		// fmt.Println("start: ", start)
		steps := 0
		for {
			for _, instruction := range strings.Split(instructions, "") {
				steps += 1
				if instruction == "L" {
					start = nodesMap[start].left
				} else {
					start = nodesMap[start].right
				}
				if strings.HasSuffix(start, "Z") {
					minSteps = append(minSteps, steps)
					// fmt.Println(minSteps)
					continue StartLoop
				}
			}
		}
	}

	return LCM(minSteps[0], minSteps[1], minSteps[2:]...)


	
	
}
