package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	// fmt.Println("part1 ans: ", part1(inputRaw))
	// fmt.Println("part2 ans: ", part2(inputRaw))
	// fmt.Println("part2 ans: ", part3(inputRaw))
	fmt.Println("part2 ans: ", part4(inputRaw))

}

func part1(input string) int {
	result := 0
	rawRows := strings.Split(input, "\n")

	for _, row := range rawRows {
		fmt.Println(row)
		allPossibilities := generateAllPossibilities(row)
		for _, possibility := range allPossibilities {
			fields := strings.Fields(possibility)
			patternString := fields[0]
			toMatch := convertStringToIntSlice(fields[1])
			if isPatternMatched(strings.Split(patternString, ""), toMatch) {
				result += 1
			}
		}
	}

	return result
}

func part2(input string) int {
	result := 0
	rawRows := strings.Split(input, "\n")

	for _, row := range rawRows {
		// fmt.Println(row)
		allPossibilities := generateAllPossibilitiesFiveTimes(row)
		for _, possibility := range allPossibilities {
			// fmt.Printf("row: %d, permutation %d of %d\n", i, j, len(allPossibilities))
			fields := strings.Fields(possibility)
			patternString := fields[0]
			toMatch := convertStringToIntSlice(fields[1])
			if isPatternMatched(strings.Split(patternString, ""), toMatch) {
				result += 1
			}
		}
	}

	return result
}

func part3(input string) int {
	result := 0
	rawRows := strings.Split(input, "\n")

	firstResults := []int{}
	for _, row := range rawRows {
		firstResult := 0
		allPossibilities := generateAllPossibilities(row)
		for _, possibility := range allPossibilities {
			fields := strings.Fields(possibility)
			patternString := fields[0]
			toMatch := convertStringToIntSlice(fields[1])
			if isPatternMatched(strings.Split(patternString, ""), toMatch) {
				firstResult += 1
			}
		}
		firstResults = append(firstResults, firstResult)
	}

	fmt.Println("starting second results")
	secondResults := []int{}
	for i, row := range rawRows {
		secondResult := 0
		allPossibilities := generateAllPossibilitiesFiveTimes(row)
		for j, possibility := range allPossibilities {
			fmt.Printf("row %d, %d of %d", i, j, len(allPossibilities))
			fields := strings.Fields(possibility)
			patternString := fields[0]
			toMatch := convertStringToIntSlice(fields[1])
			if isPatternMatched(strings.Split(patternString, ""), toMatch) {
				secondResult += 1
			}
		}
		secondResults = append(secondResults, secondResult)
	}

	for i, firstRes := range firstResults {
		secondRes := float64(secondResults[i])
		power := secondRes/float64(firstRes)
		result += int(math.Pow(power, 4)) * firstRes
	}

	return result
}

func part4(input string) int {
	result := 0
	rawRows := strings.Split(input, "\n")

	for _, row := range rawRows {
		fmt.Println(row)
		fields := strings.Fields(row)
		patternString := fields[0]
		toMatch := convertStringToIntSlice(fields[1])
		
		
	}

	return result
}

func getNumPossibilities(pattern string, remainingCombis []int, currCounter, totalPossibilities int) int {

	// last char in pattern, remainingCombis not empty
	if len(pattern) == 1 {
		if pattern[0:1] != "." {
			if cc
		}
	}

	currChar := pattern[0:1]
	if currChar == "#" || currChar == "?" {
		currCounter += 1
		if len(remainingCombis) > 0 && currCounter <= remainingCombis[0] {
			totalPossibilities += getNumPossibilities(pattern[1:], remainingCombis, currCounter, totalPossibilities)
		}
	}
	if currChar == "." || currChar == "?" {
		if currCounter > 0 {
			if len(remainingCombis) > 0 && currCounter == remainingCombis[0] {
				// good, still match
				remainingPattern := pattern[]
			}

			currCounter = 0
		} else {

		}

		
	}

	return totalPossibilities
}

func convertStringToIntSlice(pattern string) []int {
	result := []int{}
	for _, valStr := range strings.Split(pattern, ",") {
		val, _ := strconv.Atoi(valStr)
		result = append(result, val)
	}
	return result
}

func generateAllPossibilitiesFiveTimes(row string) []string {
	times := 2
	fields := strings.Fields(row)
	leftString := fields[0]
	rightString := fields[1]

	leftStringFiveTimes := ""
	leftStringSlice := []string{}
	for i:=0;i<times;i++{
		leftStringSlice = append(leftStringSlice, leftString)
	}
	leftStringFiveTimes = strings.Join(leftStringSlice, "?")

	rightStringFiveTimes := ""
	rightStringSlice := []string{}
	for i:=0;i<times;i++{
		rightStringSlice = append(rightStringSlice, rightString)
	}
	rightStringFiveTimes = strings.Join(rightStringSlice, ",")

	// fmt.Println(leftStringFiveTimes+" ",rightStringFiveTimes)
	return generateAllPossibilities(leftStringFiveTimes+" "+rightStringFiveTimes)
}

// convert xxx?xx? to [xxx.xx., xxx#xx., xxx.xx#, xxx#xx#]
func generateAllPossibilities(pattern string) []string {
	result := []string{}
	// ch := make(chan string, 2)

	firstReplacement := strings.Replace(pattern, "?", ".", 1)
	if strings.Contains(firstReplacement, "?") {
		result = append(result, generateAllPossibilities(firstReplacement)...)
	} else {
		fmt.Println(firstReplacement)
		result = append(result, firstReplacement)
	}

	secondReplacement := strings.Replace(pattern, "?", "#", 1)
	if strings.Contains(firstReplacement, "?") {
		result = append(result, generateAllPossibilities(secondReplacement)...)
	} else {
		fmt.Println(secondReplacement)
		result = append(result, secondReplacement)
	}

	return result
}

// isPatternMatched returns if pattern match the numbers
func isPatternMatched(pattern[]string, toMatch[]int) bool {
	// fmt.Println(pattern)
	// fmt.Println(toMatch)

	currDamagedCount := 0
	totalDamagedSets := 0

	for i, spring := range pattern {
		// fmt.Println(i, spring)
		switch spring {
		case ".":
			if currDamagedCount > 0 {
				if totalDamagedSets == len(toMatch) {
					return false
				}
				if currDamagedCount != toMatch[totalDamagedSets] {
					return false
				}
				totalDamagedSets += 1
				currDamagedCount = 0
			}
		case "#":
			currDamagedCount += 1
			if i == len(pattern)-1 {
				if totalDamagedSets == len(toMatch) {
					return false
				}
				if currDamagedCount != toMatch[totalDamagedSets] {
					return false
				}
				totalDamagedSets += 1
			}
		}
	}

	return totalDamagedSets == len(toMatch)
}