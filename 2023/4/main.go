package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("error reading file: ", err)
	}
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	result := 0
	toAdd := 0
	for _, row := range strings.Split(input, "\n") {
		toAdd = 0
		headers := strings.Split(row, ":")
		cardsRow := headers[1]
		cards := strings.Split(cardsRow, "|")
		winningCard := cards[0]
		handCard := cards[1]

		winningNumsString := strings.Fields(winningCard)
		handNumsString := strings.Fields(handCard)

		for _, handNumString := range handNumsString {
			for _, winningNumString := range winningNumsString {
				if handNumString == winningNumString {
					if toAdd == 0 {
						toAdd = 1
					} else {
						toAdd = toAdd * 2
					}
				}
			}
		}

		result += toAdd
	}
	return result
}

func part2(input string) int {
	rows := strings.Split(input, "\n")
	multipliers := make([]int, len(rows))

	for i, row := range rows {
		headers := strings.Split(row, ":")
		cardsRow := headers[1]
		cards := strings.Split(cardsRow, "|")
		winningCard := cards[0]
		handCard := cards[1]

		winningNumsString := strings.Fields(winningCard)
		handNumsString := strings.Fields(handCard)

		numMatches := 0
		for _, handNumString := range handNumsString {
			for _, winningNumString := range winningNumsString {
				if handNumString == winningNumString {
					numMatches += 1
				}
			}
		}

		for j := 1; j <= numMatches; j++ {
			toAddMultiplier := multipliers[i] + 1
			multipliers[i+j] += toAddMultiplier
		}
	}

	// fmt.Println("multipliers: ", multipliers)
	result := 0
	for _, multiplier := range multipliers {
		result = result + multiplier + 1
	}

	return result
}
