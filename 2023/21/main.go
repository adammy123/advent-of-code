package main

import (
	"fmt"
	"os"
	"strings"
)

const numSteps = 6

type ground struct {
	isSteppable bool
	nextSteps [][]int
}

type groundPart2 struct {
	isSteppable bool
	nextSteps []nextStepPart2
}

type nextStepPart2 struct {
	rowIdx int
	colIdx int
	isNextGrid bool
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part2(input string) int {
	grid, start := getGridPart2(input)

	for i, row := range grid {
		for j, currGround := range row {
			if currGround.isSteppable {
				grid[i][j].nextSteps = getPossibleNextStepsPart2(grid, []int{i, j})
			}
		}
		// fmt.Println(row, "\n")
	}

	currGrid := makeIntGrid(len(grid), len(grid[0]))
	prevGrid := makeIntGrid(len(grid), len(grid[0]))
	prevGrid[start[0]][start[1]] = 1

	var currPos [][]int
	var currGround groundPart2
	for i:=0; i<numSteps; i++ {
		currGrid = makeIntGrid(len(grid), len(grid[0]))
		currPos = getPositivePosInIntGrid(prevGrid)
		for _, pos := range currPos {
			currGround = grid[pos[0]][pos[1]]
			for _, nextStep := range currGround.nextSteps {
				// currGrid[nextStep[0]][nextStep[1]] += prevGrid[pos[0]][pos[1]]
				currGrid[nextStep.rowIdx][nextStep.colIdx] += 
			}
		}
		prevGrid = currGrid
		for _, row := range prevGrid {
			fmt.Println(row)
		}
		fmt.Println()
	}

	return sumIntsInGrid(prevGrid)
}

func part1(input string) int {
	grid, start := getGrid(input)

	for i, row := range grid {
		for j := range row {
			grid[i][j].nextSteps = getPossibleNextSteps(grid, []int{i, j})
		}
	}

	currGrid := makeBoolGrid(len(grid), len(grid[0]))
	prevGrid := makeBoolGrid(len(grid), len(grid[0]))
	prevGrid[start[0]][start[1]] = true

	var currPos [][]int
	var currGround ground
	for i:=0; i<numSteps; i++ {
		currGrid = makeBoolGrid(len(grid), len(grid[0]))
		currPos = getTruePosOfBoolGrid(prevGrid)
		for _, pos := range currPos {
			currGround = grid[pos[0]][pos[1]]
			for _, nextStep := range currGround.nextSteps {
				currGrid[nextStep[0]][nextStep[1]] = true
			}
		}
		prevGrid = currGrid
	}

	return numTrueInBoolGrid(prevGrid)
}

func getPossibleNextSteps(grid [][]ground, start []int) [][]int {
	nextSteps := [][]int{}

	if start[1] < len(grid[start[0]])-1 {
		if grid[start[0]][start[1]+1].isSteppable {
			nextSteps = append(nextSteps, []int{start[0], start[1]+1})
		}
	}
	if start[1] > 0 {
		if grid[start[0]][start[1]-1].isSteppable {
			nextSteps = append(nextSteps, []int{start[0], start[1]-1})
		}
	}
	if start[0] < len(grid)-1 {
		if grid[start[0]+1][start[1]].isSteppable {
			nextSteps = append(nextSteps, []int{start[0]+1, start[1]})
		}
	}
	if start[0] > 0 {
		if grid[start[0]-1][start[1]].isSteppable {
			nextSteps = append(nextSteps, []int{start[0]-1, start[1]})
		}
	}
	return nextSteps
}

func getPossibleNextStepsPart2(grid [][]groundPart2, start []int) []nextStepPart2 {
	nextSteps := []nextStepPart2{}

	if start[1] < len(grid[start[0]])-1 {
		if grid[start[0]][start[1]+1].isSteppable {
			nextSteps = append(nextSteps, nextStepPart2{start[0], start[1]+1, false})
		}
	} else if grid[start[0]][0].isSteppable {
		nextSteps = append(nextSteps, nextStepPart2{start[0], 0, true})
	}

	if start[1] > 0 {
		if grid[start[0]][start[1]-1].isSteppable {
			nextSteps = append(nextSteps, nextStepPart2{start[0], start[1]-1, false})
		}
	} else if grid[start[0]][len(grid[start[0]])-1].isSteppable {
		nextSteps = append(nextSteps, nextStepPart2{start[0], len(grid[start[0]])-1, true})
	}

	if start[0] < len(grid)-1 {
		if grid[start[0]+1][start[1]].isSteppable {
			nextSteps = append(nextSteps, nextStepPart2{start[0]+1, start[1], false})
		}
	} else if grid[0][start[1]].isSteppable {
		nextSteps = append(nextSteps, nextStepPart2{0, start[1], true})
	}
	if start[0] > 0 {
		if grid[start[0]-1][start[1]].isSteppable {
			nextSteps = append(nextSteps, nextStepPart2{start[0]-1, start[1], false})
		}
	} else if grid[len(grid)-1][start[1]].isSteppable {
		nextSteps = append(nextSteps, nextStepPart2{len(grid)-1, start[1], true})
	}
	return nextSteps
}

func getGrid(input string) ([][]ground, []int) {
	rawRows := strings.Split(input, "\n")
	grid := make([][]ground, len(rawRows))
	start := make([]int, 2)
	for i, row := range rawRows {
		gridRow := make([]ground, len(row))
		for j, val := range strings.Split(row, "") {
			if val == "S" {
				start[0] = i
				start[1] = j
				val = "."
			}
			gridRow[j] = ground{isSteppable: val=="."}
		}
		grid[i] = gridRow
	}
	return grid, start
}

func getGridPart2(input string) ([][]groundPart2, []int) {
	rawRows := strings.Split(input, "\n")
	grid := make([][]groundPart2, len(rawRows))
	start := make([]int, 2)
	for i, row := range rawRows {
		gridRow := make([]groundPart2, len(row))
		for j, val := range strings.Split(row, "") {
			if val == "S" {
				start[0] = i
				start[1] = j
				val = "."
			}
			gridRow[j] = groundPart2{isSteppable: val=="."}
		}
		grid[i] = gridRow
	}
	return grid, start
}

func makeBoolGrid(numRows, numCols int) [][]bool {
	grid := make([][]bool, numRows)
	for i:=0; i<numRows; i++ {
		grid[i] = make([]bool, numCols)
	}
	return grid
}

func numTrueInBoolGrid(grid [][]bool) int {
	total := 0
	for _, row := range grid {
		for _, val := range row {
			if val {
				total += 1
			}
		}
	}
	return total
}

func getTruePosOfBoolGrid(grid [][]bool) [][]int {
	pos := [][]int{}
	for i, row := range grid {
		for j, val := range row {
			if val {
				pos = append(pos, []int{i, j})
			}
		}
	}

	return pos
}

func makeIntGrid(numRows, numCols int) [][]int {
	grid := make([][]int, numRows)
	for i:=0; i<numRows; i++ {
		grid[i] = make([]int, numCols)
	}
	return grid
}

func sumIntsInGrid(grid [][]int) int {
	total := 0
	for _, row := range grid {
		for _, val := range row {
			total += val
		}
	}
	return total
}


func getPositivePosInIntGrid(grid [][]int) [][]int {
	pos := [][]int{}
	for i, row := range grid {
		for j, val := range row {
			if val > 0 {
				pos = append(pos, []int{i, j})
			}
		}
	}

	return pos
}