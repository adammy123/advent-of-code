package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type lens struct {
	code string
	focalLength int
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part2(input string) int {
	boxes := make([][]lens, 256)
	codes := strings.Split(input, ",")
	for _, code := range codes {
		// fmt.Println(code)
		if toAdd(code) { //contains =
			fields := strings.Split(code, "=")
			newCode := fields[0]
			focalLength, _ := strconv.Atoi(fields[1])
			lens := lens{newCode, focalLength}
			boxNum := getHashFromCode(newCode)

			//if box[boxNum] already has box with code, replace
			existingIndex := getExistingLensIndex(boxes[boxNum], newCode)
			if existingIndex >= 0 {
				boxes[boxNum][existingIndex].focalLength = focalLength
			} else {
				//else append
				boxes[boxNum] = append(boxes[boxNum], lens)
			}

		} else { //cotains -
			fields := strings.Split(code, "-")
			newCode := fields[0]
			boxNum := getHashFromCode(newCode)
			box := boxes[boxNum]

			existingIndex := getExistingLensIndex(box, newCode)
			if existingIndex == 0 {
				if len(box) == 1 {
					boxes[boxNum] = []lens{}
				} else {
					boxes[boxNum] = box[1:]
				}
			} else if existingIndex > 0 && existingIndex == len(box)-1 {
				boxes[boxNum] = box[:len(box)-1]
			} else if existingIndex > 0 {
				boxes[boxNum] = append(box[0:existingIndex], box[existingIndex+1:]...)
			}
		}
		// printBoxes(boxes)
	}

	result := 0
	for i, box := range boxes {
		if len(box) != 0 {
			for j, lens := range box {
				result += (i+1)*(j+1)*lens.focalLength
			}
		}
	}

	return result
}

func printBoxes(boxes [][]lens) {
	for i, box := range boxes {
		if len(box) > 0 {
			fmt.Println(i, box)
		}
	}
	fmt.Println()
}

func toAdd(code string) bool {
	return !strings.HasSuffix(code, "-")
}

func getExistingLensIndex(lenses []lens, code string) int {
	for i, lens := range lenses {
		if lens.code == code {
			return i
		}
	}
	return -1
}

func part1(input string) int {
	result := 0
	codes := strings.Split(input, ",")

	for _, code := range codes {
		result += getHashFromCode(code)
	}
	return result
}

func getHashFromCode(code string) int {
	currHash := 0
	for _, char := range code {
		ascii := int(char)
		currHash += ascii
		currHash *= 17
		currHash = int(math.Mod(float64(currHash), 256))
	}
	// fmt.Println(currHash)
	return currHash
}
