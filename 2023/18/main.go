package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 69698 too low

type step struct {
	direction string
	number int
	colour string
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	// fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	steps := convertInputToSteps(input)
	
	// start at pseudo (0, 0)
	// get max top-left and bottom-right coords
	maxTop := 0
	maxLeft := 0
	maxRight := 0
	maxBottom := 0
	coords := []int{0, 0}

	for _, step := range steps {
		switch step.direction {
		case "R":
			coords[1] += step.number
		case "L":
			coords[1] -= step.number
		case "U":
			coords[0] -= step.number
		case "D":
			coords[0] += step.number
		}
		maxTop = min(maxTop, coords[0])
		maxBottom = max(maxBottom, coords[0])
		maxLeft = min(maxLeft, coords[1])
		maxRight = max(maxRight, coords[1])
	}

	fmt.Println("maxTop: ", maxTop)
	fmt.Println("maxBottom: ", maxBottom)
	fmt.Println("maxRight: ", maxRight)
	fmt.Println("maxLeft: ", maxLeft)

	// make empty grid
	height := maxBottom-maxTop+1
	width := maxRight-maxLeft+1
	grid := makeGrid(height, width)

	// draw trench
	startPos := []int{-maxTop, -maxLeft}
	for i, step := range steps {
		switch step.direction {
		case "R":
			if i>0 {
				if steps[i-1].direction == "D" {
					grid[startPos[0]][startPos[1]] = "L"
				} else {
					grid[startPos[0]][startPos[1]] = "F"
				}
			}
			for i:=1; i<=step.number; i++ {
				grid[startPos[0]][startPos[1]+1] = "-"
				startPos[1] += 1
			}
		case "L":
			if i>0 {
				if steps[i-1].direction == "D" {
					grid[startPos[0]][startPos[1]] = "J"
				} else {
					grid[startPos[0]][startPos[1]] = "7"
				}
			}
			for i:=1; i<=step.number; i++ {
				grid[startPos[0]][startPos[1]-1] = "-"
				startPos[1] -= 1
			}
		case "U":
			if i>0 {
				if steps[i-1].direction == "R" {
					grid[startPos[0]][startPos[1]] = "J"
				} else {
					grid[startPos[0]][startPos[1]] = "L"
				}
			}
			for i:=1; i<=step.number; i++ {
				grid[startPos[0]-1][startPos[1]] = "|"
				startPos[0] -= 1
			}
		case "D":
			if i>0 {
				if steps[i-1].direction == "R" {
					grid[startPos[0]][startPos[1]] = "7"
				} else {
					grid[startPos[0]][startPos[1]] = "F"
				}
			}
			for i:=1; i<=step.number; i++ {
				grid[startPos[0]+1][startPos[1]] = "|"
				startPos[0] += 1
			}
		}
	}
	grid[startPos[0]][startPos[1]] = "F"

	// calculate area
	area := 0

	for _, row := range grid {
		inside := false
		prevCorner := ""
		for _, val := range row {
			switch val {
			case "-":
				area += 1
			case ".":
				if inside {
					area += 1
				}
			case "|":
				area += 1
				inside = !inside
			case "F":
				area += 1
				inside = !inside
				prevCorner = "F"
			case "L":
				area += 1
				inside = !inside
				prevCorner = "L"
			case "7":
				area += 1
				if prevCorner == "F" {
					inside = !inside
				}
			case "J":
				area += 1
				if prevCorner == "L" {
					inside = !inside
				}
			}
		}
		fmt.Println(area)
	}


	for _, row := range grid {
		fmt.Println(row)
	}

	return area
}

func convertInputToSteps(input string) []step {
	steps := []step{}
	for _, row := range strings.Split(input, "\n") {
		fields := strings.Fields(row)
		thisDir := fields[0]
		thisNum, _ := strconv.Atoi(fields[1])
		thisColour := strings.ReplaceAll(fields[2][1:], ")", "")
		steps = append(steps, step{thisDir, thisNum, thisColour})
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func makeGrid(height, width int) [][]string {
	grid := [][]string{}
	for i:=0; i<height; i++ {
		thisRow := []string{}
		for j:=0; j<width; j++ {
			thisRow = append(thisRow, ".")
		}
		grid = append(grid, thisRow)
	}
	return grid
}