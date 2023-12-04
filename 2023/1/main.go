package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("error reading file: ", err)
	}
	inputRaw := string(data)
	// fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	sum := 0
	re := regexp.MustCompile("[1-9]")
	for _, row := range strings.Split(input, "\n") {
		// fmt.Println("row: ", row)
		nums := re.FindAllString(row, -1)
		// fmt.Println("nums: ", nums)
		firstDigit := nums[0]
		lastDigit := nums[len(nums)-1]
		strNum := firstDigit + lastDigit
		// fmt.Println("strNum: ", strNum)
		num, err := strconv.Atoi(strNum)
		if err != nil {
			log.Fatal("err converting strNum to num: ", err)
		}
		sum += num
		// fmt.Println("")
	}

	return sum

}

func part2(input string) int {
	newInput := []string{}
	numStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, row := range strings.Split(input, "\n") {

		fmt.Println("old row", row)
		minIndex := math.MaxInt
		stringToReplace := ""
		numToReplaceTo := 0

		for i, numString := range numStrings {
			// fmt.Println("Row: ", row)
			// fmt.Println("numString: ", numString)
			numIndex := strings.Index(row, numString)
			if numIndex > -1 && numIndex < minIndex {
				// fmt.Println("numIndex: ", numIndex)
				stringToReplace = numString
				numToReplaceTo = i + 1
				minIndex = numIndex
				// fmt.Println("stringToReplace: ", stringToReplace)
			}
		}
		stringToReplaceTo := strconv.Itoa(numToReplaceTo)
		row = strings.Replace(row, stringToReplace, stringToReplaceTo, 1)

		maxIndex := math.MinInt
		stringToReplace = ""
		numToReplaceTo = 0

		for i, numString := range numStrings {
			// fmt.Println("Row: ", row)
			// fmt.Println("numString: ", numString)
			numIndex := strings.Index(row, numString)
			if numIndex > -1 && numIndex > maxIndex {
				// fmt.Println("numIndex: ", numIndex)
				stringToReplace = numString
				numToReplaceTo = i + 1
				maxIndex = numIndex
				// fmt.Println("stringToReplace: ", stringToReplace)
			}
		}
		stringToReplaceTo = strconv.Itoa(numToReplaceTo)
		row = strings.Replace(row, stringToReplace, stringToReplaceTo, 1)

		newInput = append(newInput, row)
		// fmt.Println("new row", row)
		// fmt.Println("")
	}
	return part1(strings.Join(newInput, "\n"))
}
