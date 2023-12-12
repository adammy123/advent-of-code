package main

import (
	"fmt"
	"os"
	"strings"
)
//78419242 too low
//702152204842
const expansionFactor = 1000000

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	// fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	result := 0
	rawRows := strings.Split(input, "\n")
	image := [][]string{}
	emptyRowIndices := map[int]bool{}
	nonEmptyColIndices := make([]bool, len(rawRows[0]))

	for rowIdx, row := range rawRows {
		imageRow := []string{}
		rowIsEmpty := true
		cols := strings.Split(row, "")
		for colIdx, val := range cols {
			if val == "#" {
				rowIsEmpty = false
				nonEmptyColIndices[colIdx] = true
			}
			imageRow = append(imageRow, val)
		}
		if rowIsEmpty {
			emptyRowIndices[rowIdx] = true
		}
		image = append(image, imageRow)
	}

	expandedImage := [][]string{}
	for i, row := range image {
		expandedImage = append(expandedImage, row)
		if _, ok := emptyRowIndices[i]; ok {
			expandedImage = append(expandedImage, row)
		}
	}

	fmt.Println(nonEmptyColIndices)
	for i, row := range expandedImage {
		counter := 0
		// fmt.Println(i)
		for colIdx, isNonEmpty := range nonEmptyColIndices {
			if !isNonEmpty {
				// fmt.Println("old row: ", row)
				row = append(row[:colIdx+counter], append([]string{"."}, row[colIdx+counter:]...)...)
				counter += 1
				// fmt.Println("new row: ", row)
			}
		}
		expandedImage[i] = row
	}

	galaxyPositions := [][]int{}
	for rowIdx, imageRow := range expandedImage {
		fmt.Println(imageRow)
		for colIdx, val := range imageRow {
			if val == "#" {
				galaxyPositions = append(galaxyPositions, []int{rowIdx, colIdx})
			}
		}
	}

	// fmt.Println(galaxyPositions)
	for i, pos := range galaxyPositions {
		for _, otherPos := range galaxyPositions[i+1:] {
			result += abs(pos[0]-otherPos[0]) + abs(pos[1]-otherPos[1])
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func part2(input string) int {
	result := 0
	rawRows := strings.Split(input, "\n")
	image := [][]string{}
	emptyRowIndices := map[int]bool{}
	nonEmptyColIndices := make([]bool, len(rawRows[0]))

	for rowIdx, row := range rawRows {
		imageRow := []string{}
		rowIsEmpty := true
		cols := strings.Split(row, "")
		for colIdx, val := range cols {
			if val == "#" {
				rowIsEmpty = false
				nonEmptyColIndices[colIdx] = true
			}
			imageRow = append(imageRow, val)
		}
		if rowIsEmpty {
			emptyRowIndices[rowIdx] = true
		}
		image = append(image, imageRow)
	}

	galaxyPositions := [][]int{}
	for rowIdx, imageRow := range image {
		fmt.Println(imageRow)
		for colIdx, val := range imageRow {
			if val == "#" {
				galaxyPositions = append(galaxyPositions, []int{rowIdx, colIdx})
			}
		}
	}

	fmt.Println("empty rows: ", emptyRowIndices)
	fmt.Println("nonEmpty cols: ", nonEmptyColIndices)
	for _, row := range image {
		fmt.Println(row)
	}

	fmt.Println(galaxyPositions)
	for i, pos := range galaxyPositions {
		for _, otherPos := range galaxyPositions[i+1:] {
			rowDist := getDistWithRowExpansion(pos[0], otherPos[0], emptyRowIndices)
			colDist := getDistWithColExpansion(pos[1], otherPos[1], nonEmptyColIndices)
			result += rowDist + colDist
		}
	}

	return result
}

func getDistWithRowExpansion(x, y int, expandedRows map[int]bool) int {
	dist := abs(x-y)
	for k := range expandedRows {
		if isBetween(x, y, k) {
			dist += expansionFactor-1
		}
	}
	return dist
}

func getDistWithColExpansion(a, b int, expandedCols []bool) int {
	dist := abs(a-b)
	for j, isNotExpanded := range expandedCols{
		if !isNotExpanded {
			if isBetween(a, b, j) {
				dist += expansionFactor-1
			}
		}
	}
	return dist
}

// isBetween returns if x is between a and b
func isBetween(a, b, x int) bool {
	// assume a != b
	if b > a {
		return x < b && x > a
	}
	return x < a && x > b
}

