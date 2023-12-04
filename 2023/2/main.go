package main

import (
	"fmt"
	"log"
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
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	// red, blue, green
	masterColours := []int{12, 14, 13}
	// fmt.Println("masterColours: ", masterColours)
	sumIDs := 0
	re := regexp.MustCompile("[0-9]+")

	for _, row := range strings.Split(input, "\n") {
		toAdd := true
		headers := strings.Split(row, ":")
		gameNum, _ := strconv.Atoi(strings.Split(headers[0], " ")[1])
		// fmt.Println("gameNum: ", gameNum)

		sets := strings.Split(headers[1], ";")
		for _, set := range sets {
			//format
			colours := []int{0, 0, 0}
			for _, colour := range strings.Split(set, ",") {
				//red
				if strings.Contains(colour, "red") {
					countString := re.FindAllString(colour, 1)
					if countString != nil {
						countNum, _ := strconv.Atoi(countString[0])
						colours[0] = countNum
					}
				}

				//blue
				if strings.Contains(colour, "blue") {
					countString := re.FindAllString(colour, 1)
					if countString != nil {
						countNum, _ := strconv.Atoi(countString[0])
						colours[1] = countNum
					}
				}

				//green
				if strings.Contains(colour, "green") {
					countString := re.FindAllString(colour, 1)
					if countString != nil {
						countNum, _ := strconv.Atoi(countString[0])
						colours[2] = countNum
					}
				}

			}
			//check if not possible, then add gameNum to sumIDs
			// fmt.Println("colours: ", colours)
			for i, masterCount := range masterColours {
				if masterCount < colours[i] {
					toAdd = false
					// fmt.Println("setting toAdd to false")
				}
			}
		}
		// fmt.Println("toAdd: ", toAdd)
		if toAdd {
			sumIDs += gameNum
		}
	}
	return sumIDs
}

func part2(input string) int {
	sum := 0
	re := regexp.MustCompile("[0-9]+")

	for _, row := range strings.Split(input, "\n") {
		headers := strings.Split(row, ":")
		// red, blue, green
		colours := []int{0, 0, 0}

		sets := strings.Split(headers[1], ";")
		for _, set := range sets {
			//format
			for _, colour := range strings.Split(set, ",") {
				//red
				if strings.Contains(colour, "red") {
					countString := re.FindAllString(colour, 1)
					if countString != nil {
						countNum, _ := strconv.Atoi(countString[0])
						if countNum > colours[0] {
							colours[0] = countNum
						}
					}
				}

				//blue
				if strings.Contains(colour, "blue") {
					countString := re.FindAllString(colour, 1)
					if countString != nil {
						countNum, _ := strconv.Atoi(countString[0])
						if countNum > colours[1] {
							colours[1] = countNum
						}
					}
				}

				//green
				if strings.Contains(colour, "green") {
					countString := re.FindAllString(colour, 1)
					if countString != nil {
						countNum, _ := strconv.Atoi(countString[0])
						if countNum > colours[2] {
							colours[2] = countNum
						}
					}
				}

			}
		}
		// fmt.Println("colours: ", colours)
		sum += (colours[0] * colours[1] * colours[2])
	}
	return sum
}
