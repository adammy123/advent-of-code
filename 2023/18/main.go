package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var numDirMap map[string]string = map[string]string{
	"0": "R",
	"1": "D",
	"2": "L",
	"3": "U",
}

type step struct {
	direction string
	number int
}

type pipe struct {
	colIdx int
	dir string
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part2(input string) int {
	// steps := convertInputToNewSteps(input)
	steps := convertInputToSteps(input)
	maxTop, _, _, maxLeft := getMaxCorners(steps)

	// {rowIdx: [{colIdx, "dir"}]}}
	// "dir" can be |, F, 7, J or L
	pipeMap := map[int][]pipe{}

	// [rowIdx, colIdx]
	currRow := -maxTop
	currCol := -maxLeft
	prevDir := "D" //hack

	for _, step := range steps {
		// fmt.Println(step)
		switch step.direction {
		case "R":
			if prevDir == "U" {
				addToPipeMap(pipeMap, currRow, currCol, "F")
			} else {//prevDir == "D"
				addToPipeMap(pipeMap, currRow, currCol, "L")
			}
			currCol += step.number
		case "L":
			if prevDir == "U" {
				addToPipeMap(pipeMap, currRow, currCol, "7")
			} else {//prevDir == "D"
				addToPipeMap(pipeMap, currRow, currCol, "J")
			}
			currCol -= step.number
		case "U":
			if prevDir == "R" {
				addToPipeMap(pipeMap, currRow, currCol, "J")
			} else {//prevDir == "L"
				addToPipeMap(pipeMap, currRow, currCol, "L")
			}
			// add "|" all the way up
			addStraightPipesToMap(pipeMap, currRow, currCol, step.number, step.direction)
			currRow -= step.number
		case "D":
			if prevDir == "R" {
				addToPipeMap(pipeMap, currRow, currCol, "7")
			} else {//prevDir == "L"
				addToPipeMap(pipeMap, currRow, currCol, "F")
			}
			// add "|" all the way up
			addStraightPipesToMap(pipeMap, currRow, currCol, step.number, step.direction)
			currRow += step.number
		}
		prevDir = step.direction
		// for k, v := range pipeMap {
		// 	fmt.Println(k, v)
		// }
		// fmt.Println()
	}

	// for k, v := range pipeMap {
	// 	fmt.Println(k, v)
	// }


	return calculateArea(pipeMap)
}

func calculateArea(pipeMap map[int][]pipe) int {
	var prevCol int
	var prevDir string
	var inside bool
	area := 0
	for _, pipes := range pipeMap {
		sorted := sortPipes(pipes)

		inside = false
		for _, pipe := range pipes {
			switch pipe.dir {
			case "|":
				if inside {
					area += pipe.colIdx-prevCol
				}
				inside = !inside
			case "F":
				
			case "L":
			case "7":
			case "J":
			}
			prevCol = pipe.colIdx
			prevDir = pipe.dir
		}
	}
	return area
}

func sortPipes(pipes []pipe) []pipe {
	if len(pipes) == 1 {
		return pipes
	}
	sort.Slice(pipes, func(i, j int) bool {
		return pipes[i].colIdx < pipes[j].colIdx
	  })
	return pipes
}

func addStraightPipesToMap(pipeMap map[int][]pipe, currRow, currCol, numSteps int, dir string) {
	multiplier := 1
	if dir == "U" {
		multiplier = -1
	}
	for i:=1; i<numSteps; i++ {
		addToPipeMap(pipeMap, currRow+(i*multiplier), currCol, "|")
	}
}

func addToPipeMap(pipeMap map[int][]pipe, rowIdx, colIdx int, dir string) {
	toAdd := pipe{colIdx: colIdx, dir: dir}
	if pipes, ok := pipeMap[rowIdx]; ok {
		pipeMap[rowIdx] = append(pipes, toAdd)
	} else {
		pipeMap[rowIdx] = []pipe{toAdd}
	}
}

func convertInputToNewSteps(input string) []step {
	inputRows := strings.Split(input, "\n")
	steps := make([]step, len(inputRows))
	for i, row := range inputRows {
		fields := strings.Fields(row)
		hexNum := string(fields[2][2:7])
		dirNum := string(fields[2][7:8])
		num, _ := strconv.ParseInt(hexNum, 16, 64)
		steps[i] = step{direction: numDirMap[dirNum], number: int(num)}
	}
	return steps
}

func part1(input string) int {
	steps := convertInputToSteps(input)
	maxTop, maxBottom, maxRight, maxLeft := getMaxCorners(steps)

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

// returns maxTop, maxBottom, maxRight, maxLeft
func getMaxCorners(steps []step) (int, int, int, int) {
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

	return maxTop, maxBottom, maxRight, maxLeft
}

func convertInputToSteps(input string) []step {
	steps := []step{}
	for _, row := range strings.Split(input, "\n") {
		fields := strings.Fields(row)
		thisDir := fields[0]
		thisNum, _ := strconv.Atoi(fields[1])
		steps = append(steps, step{thisDir, thisNum})
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