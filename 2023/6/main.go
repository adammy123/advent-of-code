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
	result := 1
	rows := strings.Split(input, "\n")
	timeRow := rows[0]
	distRow := rows[1]

	times := strings.Fields(strings.Split(timeRow, ":")[1])
	distances := strings.Fields(strings.Split(distRow, ":")[1])

	// fmt.Println("times: ", times)
	// fmt.Println("distances: ", distances)

	for i, timeStr := range times {
		margins := 0

		time, _ := strconv.Atoi(timeStr)
		dist, _ := strconv.Atoi(distances[i])

		for k := 1; k<time; k++ {
			coveredDist := k*(time-k)
			// fmt.Println("speed (k): ", k)
			// fmt.Println("time (time-k): ", time-k)
			// fmt.Println("distance: ", coveredDist)
			// fmt.Println()
			if coveredDist > dist {
				margins += 1
			}
		}
		fmt.Println("margins: ", margins)
		result *= margins
	}

	return result
}

func part2(input string) int {
	result := 1
	rows := strings.Split(input, "\n")
	timeRow := rows[0]
	distRow := rows[1]

	times := strings.Fields(strings.Split(timeRow, ":")[1])
	distances := strings.Fields(strings.Split(distRow, ":")[1])

	// fmt.Println("times: ", times)
	// fmt.Println("distances: ", distances)

	times = []string{strings.Join(times, "")}
	distances = []string{strings.Join(distances, "")}

	for i, timeStr := range times {
		margins := 0

		time, _ := strconv.Atoi(timeStr)
		dist, _ := strconv.Atoi(distances[i])

		for k := 1; k<time; k++ {
			coveredDist := k*(time-k)
			if coveredDist > dist {
				margins += 1
			}
		}
		fmt.Println("margins: ", margins)
		result *= margins
	}

	return result
}
