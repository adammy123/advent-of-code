package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	rawRaws := strings.Split(input, "\n")
	inputRows := make([][]int, len(rawRaws))
	lastVals := make([]int, len(rawRaws))

	for i, row := range rawRaws {
		inputRows[i] = convertStringToIntSlice(strings.Fields(row))
	}

	for i, row := range inputRows {
		lastVal := getLastVal(row, 0)
		fmt.Println(lastVal)
		lastVals[i] = getLastVal(row, 0)
		fmt.Println()
	}

	return getSumOfIntSlice(lastVals)
}

func part2(input string) int {
	rawRaws := strings.Split(input, "\n")
	inputRows := make([][]int, len(rawRaws))
	firstVals := make([]int, len(rawRaws))

	for i, row := range rawRaws {
		inputRows[i] = convertStringToIntSlice(strings.Fields(row))
	}

	for i, row := range inputRows {
		firstVal := getFirstVal(row, 0)
		fmt.Println("firstVal: ", firstVal)
		firstVals[i] = getFirstVal(row, 0)
		fmt.Println()
	}

	return getSumOfIntSlice(firstVals)
}

func getFirstVal(intSlice []int, firstVal int) int {
	fmt.Println("row: ", intSlice)
	rowLen := len(intSlice)
	if isLastRow(intSlice) {
		return intSlice[0] - firstVal
	}

	nextRow := make([]int, rowLen-1)
	for i:=0; i<rowLen-1; i++ {
		nextRow[i] = intSlice[i+1]-intSlice[i]
	}

	return intSlice[0] - getFirstVal(nextRow, firstVal)
}

func getLastVal(intSlice []int, lastVal int) int {
	fmt.Println("row: ", intSlice)
	rowLen := len(intSlice)
	if isLastRow(intSlice) {
		return intSlice[rowLen-1] + lastVal
	}
	nextRow := make([]int, rowLen-1)
	for i:=0; i<rowLen-1; i++ {
		nextRow[i] = intSlice[i+1]-intSlice[i]
	}
	return intSlice[rowLen-1] + getLastVal(nextRow, lastVal)
}

func convertStringToIntSlice(stringSlice []string) []int {
	result := make([]int, len(stringSlice))
	for i, stringVal := range stringSlice {
		num, _ := strconv.Atoi(stringVal)
		result[i] = num
	}
	return result
}

func getSumOfIntSlice(intSlice []int) int {
	result := 0
	for _, val := range intSlice {
		result += val
	}
	return result
}

func isLastRow(intSlice []int) bool {
	firstVal := intSlice[0]
	for _, val := range intSlice {
		if val != firstVal {
			return false
		}
	}
	return true
}