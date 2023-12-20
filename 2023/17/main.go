package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	value, xPos, yPos int
	previous []*point
	isStart bool // true only for [0][0]
}

func (p point) getTotalValue() int {
	if p.isStart {
		return 0
	}
	return p.value + getValuesFromPointSlice(p.previous)
}

func getValuesFromPointSlice(points []*point) int {
	total := 0
	for _, point := range points {
		total += point.value
	}
	return total
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	// fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	grid := convertInputToGrid(input)
	gridLength := len(grid)

	for i:=1; i<gridLength; i++ {

		// top row down the column && 
		// bottom row rightwards
		for j:=0; j<=i; j++ {
			
		}

		for j:=0; j<=i; j++ {
			
		}

		// bottom right corner of i

		// go backwards to i

	}

	return grid[gridLength-1][gridLength-1].getTotalValue()
}

func convertInputToGrid(input string) [][]point {
	inputRows := strings.Split(input, "\n")
	grid := make([][]point, len(inputRows))
	for i, inputRow := range inputRows {
		gridRow := make([]point, len(inputRow))
		for j, valStr := range strings.Split(inputRow, "") {
			valInt, _ := strconv.Atoi(valStr)
			gridRow[j] = point{value: valInt, xPos: i, yPos: j}
		}
		grid[i] = gridRow
	}
	grid[0][0].isStart = true
	// for _, gridRow := range grid {
		// fmt.Println(gridRow)
	// }
	return grid
}