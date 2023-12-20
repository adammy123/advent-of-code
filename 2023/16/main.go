package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	rayGrid [][]string
	grid [][]part
	gridLength, gridHeight int
)

type part struct {
	isSplitter bool
	splitterDirection string // '-' or '|'
	hasSplit bool
	isMirror bool
	mirrorDirection string // '/' or '\'

						 //   |      |
	hasReflectedTop bool // -->/ or \<--
	hasReflectedBot bool // -->\ or /<--
						 //   |      |
}

type pos struct {
	row, col int
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part2(input string) int {
	maxRays := 0
	var currDir string
	var currPos pos

	//top row
	for i:=0; i<gridLength; i++ {
		convertInputToGrids(input)
		currPos = pos{-1, i}
		currDir = "down"
		processRay(currPos, currDir)
		maxRays = max(getTotalRayCount(rayGrid), maxRays)
	}
	//bottom row
	for i:=0; i<gridLength; i++ {
		convertInputToGrids(input)
		currPos = pos{gridHeight, i}
		currDir = "up"
		processRay(currPos, currDir)
		maxRays = max(getTotalRayCount(rayGrid), maxRays)
	}
	//left column
	for i:=0; i<gridHeight; i++ {
		convertInputToGrids(input)
		currPos = pos{i, -1}
		currDir = "right"
		processRay(currPos, currDir)
		maxRays = max(getTotalRayCount(rayGrid), maxRays)
	}
	//right column
	for i:=0; i<gridHeight; i++ {
		convertInputToGrids(input)
		currPos = pos{i, gridLength}
		currDir = "left"
		processRay(currPos, currDir)
		maxRays = max(getTotalRayCount(rayGrid), maxRays)
	}

	return maxRays
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func part1(input string) int {
	convertInputToGrids(input)
	currPos := pos{0, -1}
	currDir := "right"

	processRay(currPos, currDir)

	return getTotalRayCount(rayGrid)
}

func processRay(prevPos pos, dir string) {

	// for _, row := range rayGrid {
	// 	fmt.Println(row)
	// }
	// fmt.Println(prevPos)
	// fmt.Println(dir)
	// fmt.Println()
	
	var currPos pos
	switch dir {
	case "right":
		for i:=prevPos.col+1; i<gridLength; i++ {
			currPos = pos{prevPos.row, i}
			rayGrid[currPos.row][currPos.col] = "#"
			thisPart := grid[currPos.row][currPos.col]

			if thisPart.splitterDirection == "|" {
				if !thisPart.hasSplit {
					grid[currPos.row][currPos.col].hasSplit = true
					processRay(currPos, "up")
					processRay(currPos, "down")
				}
				return
			} else if thisPart.isMirror {
				if thisPart.mirrorDirection == "/" && !thisPart.hasReflectedTop {
					grid[currPos.row][currPos.col].hasReflectedTop = true
					processRay(currPos, "up")
				} else if thisPart.mirrorDirection == "\\" && !thisPart.hasReflectedBot {
					grid[currPos.row][currPos.col].hasReflectedBot = true
					processRay(currPos, "down")
				}
				return
			}
		}
	case "left":
		for i:=prevPos.col-1; i>=0; i-- {
			currPos = pos{prevPos.row, i}
			rayGrid[currPos.row][currPos.col] = "#"
			thisPart := grid[currPos.row][currPos.col]

			if thisPart.splitterDirection == "|" {
				if !thisPart.hasSplit {
					grid[currPos.row][currPos.col].hasSplit = true
					processRay(currPos, "up")
					processRay(currPos, "down")
				}
				return
			} else if thisPart.isMirror {
				if thisPart.mirrorDirection == "/" && !thisPart.hasReflectedBot {
					grid[currPos.row][currPos.col].hasReflectedBot = true
					processRay(currPos, "down")
				} else if thisPart.mirrorDirection == "\\" && !thisPart.hasReflectedTop {
					grid[currPos.row][currPos.col].hasReflectedTop = true
					processRay(currPos, "up")
				}
				return
			}
		}
	case "down":
		for i:=prevPos.row+1; i<gridHeight; i++ {
			currPos = pos{i, prevPos.col}
			rayGrid[currPos.row][currPos.col] = "#"
			thisPart := grid[currPos.row][currPos.col]

			if thisPart.splitterDirection == "-" {
				if !thisPart.hasSplit {
					grid[currPos.row][currPos.col].hasSplit = true
					processRay(currPos, "left")
					processRay(currPos, "right")
				}
				return
			} else if thisPart.isMirror {
				if thisPart.mirrorDirection == "/" && !thisPart.hasReflectedTop {
					grid[currPos.row][currPos.col].hasReflectedTop = true
					processRay(currPos, "left")
				} else if thisPart.mirrorDirection == "\\" && !thisPart.hasReflectedTop {
					grid[currPos.row][currPos.col].hasReflectedTop = true
					processRay(currPos, "right")
				}
				return
			}
		}
	case "up":
		for i:=prevPos.row-1; i>=0; i-- {
			currPos = pos{i, prevPos.col}
			rayGrid[currPos.row][currPos.col] = "#"
			thisPart := grid[currPos.row][currPos.col]

			if thisPart.splitterDirection == "-" {
				if !thisPart.hasSplit {
					grid[currPos.row][currPos.col].hasSplit = true
					processRay(currPos, "left")
					processRay(currPos, "right")
				}
				return
			} else if thisPart.isMirror {
				if thisPart.mirrorDirection == "/" && !thisPart.hasReflectedBot {
					grid[currPos.row][currPos.col].hasReflectedBot = true
					processRay(currPos, "right")
				} else if thisPart.mirrorDirection == "\\" && !thisPart.hasReflectedBot {
					grid[currPos.row][currPos.col].hasReflectedBot = true
					processRay(currPos, "left")
				}
				return
			}
		}
	}
}

func convertInputToGrids(input string) ([][]part, [][]string) {
	inputRows := strings.Split(input, "\n")
	gridHeight = len(inputRows)
	grid = make([][]part, gridHeight)
	rayGrid = make([][]string, gridHeight)
	for i, row := range inputRows {
		gridLength = len(row)
		gridRow := make([]part, gridLength)
		rayGridRow := make([]string, gridLength)
		for j, val := range strings.Split(row, "") {
			var thisPart part
			switch val {
			case "/", "\\":
				thisPart = part{isMirror: true, mirrorDirection: val} 
			case "-", "|":
				thisPart = part{isSplitter: true, splitterDirection: val}
			default:
				thisPart = part{}
			}
			gridRow[j] = thisPart
			rayGridRow[j] = "."
		}
		grid[i] = gridRow
		rayGrid[i] = rayGridRow
	}
	return grid, rayGrid
}

func getTotalRayCount(grid [][]string) int {
	total := 0
	for _, row := range grid {
		for _, val := range row {
			if val == "#" {
				total += 1
			}
		}
	}
	return total
}