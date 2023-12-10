package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	startRow, startCol int
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	// fmt.Println("part1 ans: ", part1(inputRaw))
	// fmt.Println("part2 ans: ", part2(inputRaw))
	fmt.Println("part3 ans: ", part3(inputRaw))
}

func part1(input string) int {
	paddedRows := padInput(input)
	grid := [][]string{}

	for _, row := range paddedRows {
		gridRow := []string{}
		gridRow = append(gridRow, strings.Split(row, "")...)
		grid = append(grid, gridRow)
	}

	startRow, startCol = findStartPos(grid)
	fmt.Printf("Starting position: [%d, %d]\n\n", startRow, startCol)

	sourceRow, sourceCol := startRow, startCol
	currRow, currCol := startRow+1, startCol //hack
	tempRow, tempCol := 0, 0

	loopLength := 0
	for {
		if currRow < 0 {
			break
		}
		loopLength += 1
		tempRow, tempCol = currRow, currCol
		currRow, currCol = findNextRowCol(grid, currRow, currCol, sourceRow, sourceCol) //replace current with found next
		sourceRow, sourceCol = tempRow, tempCol                                         //replace source with old current
	}

	fmt.Println("Loop length: ", loopLength)
	return loopLength / 2
}

func findNextRowCol(grid [][]string, thisRow, thisCol, sourceRow, sourceCol int) (int, int) {
	thisChar := grid[thisRow][thisCol]
	// fmt.Println("thisChar: ", thisChar)

	switch thisChar {
	case "S":
		return -1, -1 //stop
	case "|":
		if thisRow > sourceRow {
			return thisRow + 1, thisCol //coming from bottom, go up
		}
		return thisRow - 1, thisCol //coming from top, go down
	case "-":
		if thisCol > sourceCol {
			return thisRow, thisCol + 1 //coming from left, go right
		}
		return thisRow, thisCol - 1 //coming from right, go left
	case "F":
		if thisCol < sourceCol {
			return thisRow + 1, thisCol //coming from right, go down
		}
		return thisRow, thisCol + 1 // comfing from bottom, go right
	case "7":
		if thisCol > sourceCol {
			return thisRow + 1, thisCol //coming from left, go down
		}
		return thisRow, thisCol - 1 // coming from bottowm, go left
	case "J":
		if thisRow > sourceRow {
			return thisRow, thisCol - 1 //coming from top, go left
		}
		return thisRow - 1, thisCol //coming from left, go top
	case "L":
		if thisRow > sourceRow {
			return thisRow, thisCol + 1 //coming from top, go right
		}
		return thisRow - 1, thisCol //coming from right, go top
	default:
		panic(fmt.Sprintf("Invalid char: %s", thisChar))
	}
}

func findStartPos(grid [][]string) (int, int) {
	for i, row := range grid {
		for j, letter := range row {
			if letter == "S" {
				return i, j
			}
		}
	}
	panic("Can't find start position!")
}

// pad grid with ground
func padInput(input string) []string {
	paddedInput := []string{}
	rawRows := strings.Split(input, "\n")
	rowLength := len(rawRows[0])

	groundRow := ""
	for i := 0; i < rowLength+2; i++ {
		groundRow += "."
	}
	paddedInput = append(paddedInput, groundRow)
	for _, row := range rawRows {
		paddedInput = append(paddedInput, "."+row+".")
	}
	paddedInput = append(paddedInput, groundRow)
	return paddedInput

}

func part2(input string) int {
	paddedRows := padInput(input)
	grid := [][]string{}
	originalGrid := [][]string{}
	trappedBits := [][]int{}

	for _, row := range paddedRows {
		gridRow := []string{}
		gridRow = append(gridRow, strings.Split(row, "")...)
		grid = append(grid, gridRow)
	}

	for _, row := range paddedRows {
		gridRow := []string{}
		gridRow = append(gridRow, strings.Split(row, "")...)
		originalGrid = append(originalGrid, gridRow)
	}

	for _, row := range paddedRows {
		trappedBits = append(trappedBits, make([]int, len(row)))
	}

	startRow, startCol = findStartPos(grid)
	fmt.Printf("Starting position: [%d, %d]\n\n", startRow, startCol)

	sourceRow, sourceCol := startRow, startCol
	currRow, currCol := startRow+1, startCol //hack. go bottom
	tempRow, tempCol := 0, 0

	for {
		if currRow < 0 {
			break
		}
		tempRow, tempCol = currRow, currCol
		currRow, currCol = findNextRowCol(grid, currRow, currCol, sourceRow, sourceCol) //replace current with found next
		sourceRow, sourceCol = tempRow, tempCol                                         //replace source with old current

		if sourceRow > 0 {
			grid[sourceRow][sourceCol] = "x" // to mark part of loop
		}
	}

	fmt.Println("original grid")
	for _, row := range originalGrid {
		fmt.Println(row)
	}
	fmt.Println()

	fmt.Println("grid")
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()

	sourceRow, sourceCol = startRow, startCol
	currRow, currCol = startRow+1, startCol //hack. go bottom
	tempRow, tempCol = 0, 0
	insideDirection := "right" //hack 1122/968/929/266
	numTrapped := 0
	currChar := ""
	sourceChar := ""

	for {
		currChar = originalGrid[currRow][currCol]
		sourceChar = originalGrid[sourceRow][sourceCol]
		fmt.Println(sourceChar, currChar)
		if originalGrid[currRow][currCol] == "-" || originalGrid[currRow][currCol] == "|" {
			fmt.Printf("looking %s at [%d, %d]\n", insideDirection, currRow, currCol)
			switch insideDirection {
			case "bottom":
				if grid[currRow+1][currCol] != "x" {
					trappedBits[currRow+1][currCol] = 1
					fmt.Println("1")
				}
			case "top":
				if grid[currRow-1][currCol] != "x" {
					trappedBits[currRow-1][currCol] = 1
					fmt.Println("1")
				}
			case "right":
				if grid[currRow][currCol+1] != "x" {
					trappedBits[currRow][currCol+1] = 1
					fmt.Println("1")
				}
			case "left":
				if grid[currRow][currCol-1] != "x" {
					trappedBits[currRow][currCol-1] = 1
					fmt.Println("1")
				}
			}
		}

		if originalGrid[currRow][currCol] == "L" {
			if insideDirection == "bottom" || insideDirection == "left" {
				if grid[currRow+1][currCol] != "x" {
					fmt.Printf("looking bottom at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow+1][currCol] = 1
					fmt.Println("1")
				}
				if grid[currRow+1][currCol-1] != "x" {
					fmt.Printf("looking bottom-left at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow+1][currCol-1] = 1
					fmt.Println("1")
				}
				if grid[currRow][currCol-1] != "x" {
					fmt.Printf("looking left at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow][currCol-1] = 1
					fmt.Println("1")
				}
			}
		}

		if originalGrid[currRow][currCol] == "J" {
			if insideDirection == "bottom" || insideDirection == "right" {
				if grid[currRow+1][currCol] != "x" {
					fmt.Printf("looking bottom at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow+1][currCol] = 1
					fmt.Println("1")
				}
				if grid[currRow+1][currCol+1] != "x" {
					fmt.Printf("looking bottom-right at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow+1][currCol+1] = 1
					fmt.Println("1")
				}
				if grid[currRow][currCol+1] != "x" {
					fmt.Printf("looking right at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow][currCol+1] = 1
					fmt.Println("1")
				}
			}
		}

		if originalGrid[currRow][currCol] == "7" {
			if insideDirection == "top" || insideDirection == "right" {
				if grid[currRow-1][currCol] != "x" {
					fmt.Printf("looking top at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow-1][currCol] = 1
					fmt.Println("1")
				}
				if grid[currRow-1][currCol+1] != "x" {
					fmt.Printf("looking top-right at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow-1][currCol+1] = 1
					fmt.Println("1")
				}
				if grid[currRow][currCol+1] != "x" {
					fmt.Printf("looking right at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow][currCol+1] = 1
					fmt.Println("1")
				}
			}
		}

		if originalGrid[currRow][currCol] == "F" {
			if insideDirection == "top" || insideDirection == "left" {
				if grid[currRow-1][currCol] != "x" {
					fmt.Printf("looking top at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow-1][currCol] = 1
					fmt.Println("1")
				}
				if grid[currRow-1][currCol-1] != "x" {
					fmt.Printf("looking top-left at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow-1][currCol-1] = 1
					fmt.Println("1")
				}
				if grid[currRow][currCol-1] != "x" {
					fmt.Printf("looking left at [%d, %d]\n", currRow, currCol)
					trappedBits[currRow][currCol-1] = 1
					fmt.Println("1")
				}
			}
		}

		tempRow, tempCol = currRow, currCol
		currRow, currCol = findNextRowCol(originalGrid, currRow, currCol, sourceRow, sourceCol) //replace current with found next
		sourceRow, sourceCol = tempRow, tempCol                                                 //replace source with old current

		if currRow < 0 {
			break
		}

		switch originalGrid[currRow][currCol] {
		case "7":
			if insideDirection == "bottom" {
				insideDirection = "left"
			} else if insideDirection == "top" {
				insideDirection = "right"
			} else if insideDirection == "right" {
				insideDirection = "top"
			} else if insideDirection == "left" {
				insideDirection = "bottom"
			}
		case "J":
			if insideDirection == "bottom" {
				insideDirection = "right"
			} else if insideDirection == "top" {
				insideDirection = "left"
			} else if insideDirection == "right" {
				insideDirection = "bottom"
			} else if insideDirection == "left" {
				insideDirection = "top"
			}
		case "L":
			if insideDirection == "bottom" {
				insideDirection = "left"
			} else if insideDirection == "top" {
				insideDirection = "right"
			} else if insideDirection == "right" {
				insideDirection = "top"
			} else if insideDirection == "left" {
				insideDirection = "bottom"
			}
		case "F":
			if insideDirection == "bottom" {
				insideDirection = "right"
			} else if insideDirection == "top" {
				insideDirection = "left"
			} else if insideDirection == "right" {
				insideDirection = "bottom"
			} else if insideDirection == "left" {
				insideDirection = "top"
			}
		case "-":
		case "|":
		}
	}

	fmt.Println("trapped bits")
	for _, row := range trappedBits {
		fmt.Println(row)
	}

	for _, row := range trappedBits {
		for _, val := range row {
			numTrapped += val
		}
	}

	return numTrapped
}

type part struct {
	val string
	insides []string
}

func (p part) isPartOfMainLoop() bool {
	return p.val != ""
}

func isBitInsidePart(comingFrom string, part part) bool {
	for _, inside := range part.insides {
		if comingFrom == inside {
			return true
		}
	}
	return false
}

func part3(input string) int {
	paddedRows := padInput(input)
	grid := [][]string{}
	partsGrid := [][]part{}
	for i:=0; i<len(paddedRows); i++ {
		partsGrid = append(partsGrid, make([]part, len(paddedRows[0])))
	}

	for _, row := range paddedRows {
		gridRow := []string{}
		gridRow = append(gridRow, strings.Split(row, "")...)
		grid = append(grid, gridRow)
	}

	// startRow, startCol = findStartPos(grid)
	startRow, startCol = 23, 115 //hack

	sourceRow, sourceCol := startRow, startCol
	currRow, currCol := startRow+1, startCol //hack. go bottom
	tempRow, tempCol := 0, 0
	insideDirection := "left" //hack
	var insides []string
	numTrapped := 0

	for {

		insides = []string{insideDirection}

		// if "|" or "-", no change to direction
		switch currChar := grid[currRow][currCol]; currChar {
		case "L":
			if insideDirection == "left" { //come from top
				insideDirection = "bottom"
				insides = append(insides, insideDirection)
			} else if insideDirection == "right" { //come from top
				insideDirection = "top"
			} else if insideDirection == "top" { //come from right
				insideDirection = "right"
			} else if insideDirection == "bottom" { //come from right
				insideDirection = "left"
				insides = append(insides, insideDirection)
			}
		case "J":
			if insideDirection == "right" { //come from top
				insideDirection = "bottom"
				insides = append(insides, insideDirection)
			} else if insideDirection == "left" { //come from top
				insideDirection = "top"
			} else if insideDirection == "top" { //come from left
				insideDirection = "left"
			} else if insideDirection == "bottom" { //come from left
				insideDirection = "right"
				insides = append(insides, insideDirection)
			}
		case "7":
			if insideDirection == "right" { //come from bottom
				insideDirection = "top"
				insides = append(insides, insideDirection)
			} else if insideDirection == "left" { //come from bottom
				insideDirection = "bottom"
			} else if insideDirection == "top" { //come from left
				insideDirection = "right"
				insides = append(insides, insideDirection)
			} else if insideDirection == "bottom" { //come from left
				insideDirection = "left"
			}
		case "F":
			if insideDirection == "left" { //come from bottom
				insideDirection = "top"
				insides = append(insides, insideDirection)
			} else if insideDirection == "right" { //come from bottom
				insideDirection = "bottom"
			} else if insideDirection == "top" { //come from right
				insideDirection = "left"
				insides = append(insides, insideDirection)
			} else if insideDirection == "bottom" { //come from right
				insideDirection = "right"
			}
		}

		partsGrid[currRow][currCol] = part{val: grid[currRow][currCol], insides: insides}

		if currRow == startRow && currCol == startCol { //back to starting position, exit loop
			break
		}

		tempRow, tempCol = currRow, currCol
		currRow, currCol = findNextRowCol(grid, currRow, currCol, sourceRow, sourceCol) //replace current with found next
		sourceRow, sourceCol = tempRow, tempCol                                         //replace source with old current
	}

	fmt.Println("grid")
	for _, row := range grid {
		fmt.Println(row)
		
	}
	fmt.Println()

	fmt.Println("partsGrid")
	for _, row := range partsGrid {
		fmt.Println(row)
	}
	fmt.Println()

	// for each part, look all the way to the left. once hit a pipe, if inside ++
	for rowInx, row := range partsGrid {
		ColLoop:
		for colIdx, part := range row {
			if !part.isPartOfMainLoop() {
				fmt.Printf("Looking at bit [%d, %d]", rowInx, colIdx)
				for j:=colIdx+1;j<len(row);j++{
					targetpart := row[j]
					if targetpart.isPartOfMainLoop() {
						if isBitInsidePart("left", targetpart) {
							numTrapped += 1
							fmt.Print(" 1")
						}
						fmt.Println()
						continue ColLoop
					}
				}
				fmt.Println()
			}
		}
	}

	// for {
	// 	if originalGrid[currRow][currCol] == "-" || originalGrid[currRow][currCol] == "|" {
	// 		fmt.Printf("looking %s at [%d, %d]\n", insideDirection, currRow, currCol)
	// 		switch insideDirection {
	// 		case "bottom":
	// 			if grid[currRow+1][currCol] != "x" {
	// 				trappedBits[currRow+1][currCol] = 1
	// 				fmt.Println("1")
	// 			}
	// 		case "top":
	// 			if grid[currRow-1][currCol] != "x" {
	// 				trappedBits[currRow-1][currCol] = 1
	// 				fmt.Println("1")
	// 			}
	// 		case "right":
	// 			if grid[currRow][currCol+1] != "x" {
	// 				trappedBits[currRow][currCol+1] = 1
	// 				fmt.Println("1")
	// 			}
	// 		case "left":
	// 			if grid[currRow][currCol-1] != "x" {
	// 				trappedBits[currRow][currCol-1] = 1
	// 				fmt.Println("1")
	// 			}
	// 		}
	// 	}

	// 	if originalGrid[currRow][currCol] == "L" {
	// 		if insideDirection == "bottom" || insideDirection == "left" {
	// 			if grid[currRow+1][currCol] != "x" {
	// 				fmt.Printf("looking bottom at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow+1][currCol] = 1
	// 				fmt.Println("1")
	// 			}
	// 			if grid[currRow+1][currCol-1] != "x" {
	// 				fmt.Printf("looking bottom-left at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow+1][currCol-1] = 1
	// 				fmt.Println("1")
	// 			}
	// 			if grid[currRow][currCol-1] != "x" {
	// 				fmt.Printf("looking left at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow][currCol-1] = 1
	// 				fmt.Println("1")
	// 			}
	// 		}
	// 	}

	// 	if originalGrid[currRow][currCol] == "J" {
	// 		if insideDirection == "bottom" || insideDirection == "right" {
	// 			if grid[currRow+1][currCol] != "x" {
	// 				fmt.Printf("looking bottom at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow+1][currCol] = 1
	// 				fmt.Println("1")
	// 			}
	// 			if grid[currRow+1][currCol+1] != "x" {
	// 				fmt.Printf("looking bottom-right at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow+1][currCol+1] = 1
	// 				fmt.Println("1")
	// 			}
	// 			if grid[currRow][currCol+1] != "x" {
	// 				fmt.Printf("looking right at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow][currCol+1] = 1
	// 				fmt.Println("1")
	// 			}
	// 		}
	// 	}

	// 	if originalGrid[currRow][currCol] == "7" {
	// 		if insideDirection == "top" || insideDirection == "right" {
	// 			if grid[currRow-1][currCol] != "x" {
	// 				fmt.Printf("looking top at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow-1][currCol] = 1
	// 				fmt.Println("1")
	// 			}
	// 			if grid[currRow-1][currCol+1] != "x" {
	// 				fmt.Printf("looking top-right at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow-1][currCol+1] = 1
	// 				fmt.Println("1")
	// 			}
	// 			if grid[currRow][currCol+1] != "x" {
	// 				fmt.Printf("looking right at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow][currCol+1] = 1
	// 				fmt.Println("1")
	// 			}
	// 		}
	// 	}

	// 	if originalGrid[currRow][currCol] == "F" {
	// 		if insideDirection == "top" || insideDirection == "left" {
	// 			if grid[currRow-1][currCol] != "x" {
	// 				fmt.Printf("looking top at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow-1][currCol] = 1
	// 				fmt.Println("1")
	// 			}
	// 			if grid[currRow-1][currCol-1] != "x" {
	// 				fmt.Printf("looking top-left at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow-1][currCol-1] = 1
	// 				fmt.Println("1")
	// 			}
	// 			if grid[currRow][currCol-1] != "x" {
	// 				fmt.Printf("looking left at [%d, %d]\n", currRow, currCol)
	// 				trappedBits[currRow][currCol-1] = 1
	// 				fmt.Println("1")
	// 			}
	// 		}
	// 	}

	// 	tempRow, tempCol = currRow, currCol
	// 	currRow, currCol = findNextRowCol(originalGrid, currRow, currCol, sourceRow, sourceCol) //replace current with found next
	// 	sourceRow, sourceCol = tempRow, tempCol                                                 //replace source with old current

	// 	if currRow < 0 {
	// 		break
	// 	}

	// 	switch originalGrid[currRow][currCol] {
	// 	case "7":
	// 		if insideDirection == "bottom" {
	// 			insideDirection = "left"
	// 		} else if insideDirection == "top" {
	// 			insideDirection = "right"
	// 		} else if insideDirection == "right" {
	// 			insideDirection = "top"
	// 		} else if insideDirection == "left" {
	// 			insideDirection = "bottom"
	// 		}
	// 	case "J":
	// 		if insideDirection == "bottom" {
	// 			insideDirection = "right"
	// 		} else if insideDirection == "top" {
	// 			insideDirection = "left"
	// 		} else if insideDirection == "right" {
	// 			insideDirection = "bottom"
	// 		} else if insideDirection == "left" {
	// 			insideDirection = "top"
	// 		}
	// 	case "L":
	// 		if insideDirection == "bottom" {
	// 			insideDirection = "left"
	// 		} else if insideDirection == "top" {
	// 			insideDirection = "right"
	// 		} else if insideDirection == "right" {
	// 			insideDirection = "top"
	// 		} else if insideDirection == "left" {
	// 			insideDirection = "bottom"
	// 		}
	// 	case "F":
	// 		if insideDirection == "bottom" {
	// 			insideDirection = "right"
	// 		} else if insideDirection == "top" {
	// 			insideDirection = "left"
	// 		} else if insideDirection == "right" {
	// 			insideDirection = "bottom"
	// 		} else if insideDirection == "left" {
	// 			insideDirection = "top"
	// 		}
	// 	case "-":
	// 	case "|":
	// 	}
	// }

	return numTrapped
}
