package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var symbols = []string{"*", "-", "+", "/", "@", "=", "#", "&", "%", "$"}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("error reading file: ", err)
	}
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part2(input string) int {
	sum := 0

	re := regexp.MustCompile("[*]")
	rows := strings.Split(input, "\n")

	// add blank row to top and bottom
	blankRow := "............................................................................................................................................"
	rows = append([]string{blankRow}, rows...)
	rows = append(rows, blankRow)

	// add a . to the start and end of each row
	for i, row := range rows {
		rows[i] = "." + row + "."
	}

	for i := 1; i < len(rows)-1; i++ {
		// fmt.Println("row: ", i)
		// nums := re.FindAllString(rows[i], -1)
		asterisksStartEndIndices := re.FindAllStringIndex(rows[i], -1)
		// fmt.Println("asterisksStartEndIndices: ", asterisksStartEndIndices)
		for _, asterisksStartEndIndex := range asterisksStartEndIndices {
			asteriskIndex := asterisksStartEndIndex[0]
			// fmt.Println("asteriskIndex: ", asteriskIndex)
			sum += addAsterisk(asteriskIndex, rows[i-1:i+2])
		}
	}
	return sum
}

func addAsterisk(asteriskIndex int, threeRows []string) int {
	re := regexp.MustCompile("[0-9]+")
	candidates := []int{}

	for _, row := range threeRows {
		nums := re.FindAllStringIndex(row, -1)
		for _, numIndices := range nums {
			// fmt.Println("numIndices: ", numIndices)
			numStart := numIndices[0]
			numEnd := numIndices[1]
			if asteriskIndex >= numStart && asteriskIndex < numEnd ||
				numEnd == asteriskIndex || numStart == asteriskIndex+1 {
				num, _ := strconv.Atoi(row[numStart:numEnd])
				candidates = append(candidates, num)
			}
		}
	}

	// fmt.Println("candidates: ", candidates)
	if len(candidates) == 2 {
		result := candidates[0] * candidates[1]
		// fmt.Printf("[%d] +%d\n\n", asteriskIndex, result)
		return result
	}
	// fmt.Printf("[%d] -\n\n", asteriskIndex)
	return 0
}

func part1(input string) int {
	sum := 0
	re := regexp.MustCompile("[0-9]+")
	rows := strings.Split(input, "\n")

	// add blank row to top and bottom
	blankRow := "............................................................................................................................................"
	rows = append([]string{blankRow}, rows...)
	rows = append(rows, blankRow)

	// add a . to the start and end of each row
	for i, row := range rows {
		rows[i] = "." + row + "."
	}

	for i := 1; i < len(rows)-1; i++ {
		// fmt.Println("row: ", i)
		nums := re.FindAllStringIndex(rows[i], -1)
		for _, num := range nums {
			sum += addNum(num, rows[i-1:i+2])
		}
	}
	return sum
}

func addNum(numIndices []int, theeRows []string) int {
	// topRow := theeRows[0]
	targetRow := theeRows[1]
	// bottomRow := theeRows[2]
	numStartIndex := numIndices[0]
	numEndIndex := numIndices[1]
	num := targetRow[numStartIndex:numEndIndex]

	// startIndex := strings.Index(targetRow, num)

	for i := numStartIndex; i <= numEndIndex; i++ {
		for _, row := range theeRows {
			if strings.ContainsAny(row[i-1:i+1], strings.Join(symbols, "")) {
				result, _ := strconv.Atoi(num)
				// fmt.Println("+: ", num)
				return result
			}
		}
	}
	// fmt.Println("- ", num)
	return 0
}
