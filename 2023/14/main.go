package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// 90950 too low

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1(input string) int {
	rawRows := strings.Split(input, "\n")

	grid := [][]string{}
	for _, row := range rawRows {
		grid = append(grid, strings.Split(row, ""))
	}
	
	numCycles := 1
	for i:=1; i<=numCycles; i++{
		rollNorth(grid)
	}

	return getResult(grid)
}

func part2(input string) int {
	rawRows := strings.Split(input, "\n")

	grid := [][]string{}
	for _, row := range rawRows {
		grid = append(grid, strings.Split(row, ""))
	}
	
	numCycles := 1000
	results := map[int]int{}
	consecutiveHits := []int{}
	toMinus := 0

	for i:=1; i<+numCycles; i++{
		rollNorth(grid)
		rollWest(grid)
		rollSouth(grid)
		rollEast(grid)
		
		result := getResult(grid)
		if _, ok := results[result]; ok {
			results[result] += 1
			// fmt.Printf("HIT: %d, count: %d, cycle:%d\n", result, val, i)

			if len(consecutiveHits)>0 && consecutiveHits[0]==result {
				fmt.Println("consecutive hit cycle hit. length: ", len(consecutiveHits))
				toMinus = numCycles-i
				break
			}

			consecutiveHits = append(consecutiveHits, result)
		} else {
			results[result] = 1
			consecutiveHits = []int{} //reset cache
		}
	}

	patternIdx := math.Mod(float64(numCycles-toMinus), float64(len(consecutiveHits)))


	return consecutiveHits[int(patternIdx)]
}

func getResult(grid [][]string) int {
	result := 0
	for i:=0; i<len(grid); i++ {
		for _, val := range grid[i] {
			if val == "O" {
				result += len(grid)-i
			}
		}
	}
	return result
}

func rollWest(grid [][]string) {
	for _, row := range grid {
		for i:=1; i<len(row); i++ {
			value := row[i]
			if value == "O" {
				newRockColIdx := i
				for newColIdx:=i-1; newColIdx>=0; newColIdx-- {
					if row[newColIdx] != "." {
						break
					}
					newRockColIdx = newColIdx
				}
				row[i] = "."
				row[newRockColIdx] = "O"
			}
		}
	}
}

func rollEast(grid [][]string) {
	for _, row := range grid {
		for i:=len(row)-2; i>=0; i-- {
			value := row[i]
			if value == "O" {
				newRockColIdx := i
				for newColIdx:=i+1; newColIdx<len(row); newColIdx++ {
					if row[newColIdx] != "." {
						break
					}
					newRockColIdx = newColIdx
				}
				row[i] = "."
				row[newRockColIdx] = "O"
			}
		}
	}
}

func rollNorth(grid [][]string) {
	for rowIdx:=1; rowIdx<len(grid); rowIdx++ {
		for colIdx:=0; colIdx<len(grid[0]); colIdx++ {
			value := grid[rowIdx][colIdx]
			if value == "O" {
				newRockRowIdx := rowIdx
				for newRowIdx:=rowIdx-1; newRowIdx>=0; newRowIdx-- {
					if grid[newRowIdx][colIdx] != "." {
						break
					}
					newRockRowIdx = newRowIdx
				}
				grid[rowIdx][colIdx] = "."
				grid[newRockRowIdx][colIdx] = "O"
			}
		}
	}
}

func rollSouth(grid [][]string) {
	for rowIdx:=len(grid)-2; rowIdx>=0; rowIdx-- {
		for colIdx:=0; colIdx<len(grid[0]); colIdx++ {
			value := grid[rowIdx][colIdx]
			if value == "O" {
				newRockRowIdx := rowIdx
				for newRowIdx:=rowIdx+1; newRowIdx<len(grid[0]); newRowIdx++ {
					if grid[newRowIdx][colIdx] != "." {
						break
					}
					newRockRowIdx = newRowIdx
				}
				grid[rowIdx][colIdx] = "."
				grid[newRockRowIdx][colIdx] = "O"
			}
		}
	}
}

func intSliceContains(slice []int, target int) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}