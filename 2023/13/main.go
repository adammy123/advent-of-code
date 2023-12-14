package main

import (
	"fmt"
	"os"
	"strings"
)

// 33110 = too low

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part2(inputRaw, true))
	fmt.Println("part2 ans: ", part2(inputRaw, false))
}

func part2(input string, smudgeCleaned bool) int {
	result := 0
	blocks := strings.Split(input, "\n\n")
	
	BlockLoop:
	for i, block := range blocks {
		fmt.Print("block: ", i)
		horizontalIndex := getHorizontalIndex(block, smudgeCleaned)
		if horizontalIndex > 0 {
			result += 100*horizontalIndex
			fmt.Println(" horizontal at ", horizontalIndex)
			continue BlockLoop
		}

		verticalIndex := getVerticalIndex(block, smudgeCleaned)
		if verticalIndex < 0 {
			panic("no horizontal or vertical mirror found!")
		}
		fmt.Println(" vertical at ", verticalIndex)
		result += verticalIndex
	}

	return result
}


func getHorizontalIndex(block string, originalSmudgeCleaned bool) int {
	rows := strings.Split(block, "\n")
	numRows := len(rows)
	var smudgeCleaned bool

	RowLoop:
	for i:=0; i<numRows-1; i++ {
		smudgeCleaned = originalSmudgeCleaned
		if rows[i] != rows[i+1] {
			if !smudgeCleaned && canClean(rows[i], rows[i+1]) {
				smudgeCleaned = true
				toCheck := min(i, numRows-i-2)
				for j:=1; j<=toCheck; j++ {
					if rows[i-j] != rows[i+1+j] {
						if smudgeCleaned {
							continue RowLoop
						}
					}
				}
				// fmt.Println("found horizontal after row", i+1)
				return i+1
			}
		}
		if rows[i] == rows[i+1] {
			toCheck := min(i, numRows-i-2)
			for j:=1; j<=toCheck; j++ {
				if rows[i-j] != rows[i+1+j] {
					if smudgeCleaned {
						continue RowLoop
					}
					if !smudgeCleaned && canClean(rows[i-j], rows[i+1+j]) {
						smudgeCleaned = true
					}
				}
			}
			if smudgeCleaned {
				// fmt.Println("found horizontal after row", i+1)
				return i+1
			}
		}
	}

	// no horizontal mirror found
	// fmt.Println("no horizontal mirror found")
	return -1
}

func canClean(firstRow, secondRow string) bool {
	cleaned := false
	for i:=0; i<len(firstRow); i++ {
		if firstRow[i:i+1] != secondRow[i:i+1] {
			if cleaned {
				return false
			}
			cleaned = true
		}
	}
	// fmt.Printf("first: %s\nsecond:%s\ncanclean:%v\n", firstRow, secondRow, cleaned)
	return true
}

func getVerticalIndex(block string, originalSmudgeCleaned bool) int {
	rows := strings.Split(block, "\n")
	numRows := len(rows)
	var smudgeCleaned bool
	var isMirrored bool

	topRow := rows[0]
	rowLen := len(topRow)

	RowLoop:
	for i:=0; i<rowLen-1; i++ {
		smudgeCleaned = originalSmudgeCleaned
		isMirrored = false
		isMirrored, smudgeCleaned = isMirroredAfterWithCleanPossibility(topRow, i, smudgeCleaned)
		if isMirrored {
			for j:=1; j<numRows; j++ {
				isMirrored, smudgeCleaned = isMirroredAfterWithCleanPossibility(rows[j], i, smudgeCleaned)
				if !isMirrored {
					continue RowLoop
				}
			}
			if smudgeCleaned {
				// fmt.Println("found horizontal after row", i+1)
				return i+1
			}
		}
	}

	// no vertical mirror found
	return -1
}

func isMirroredAfterWithCleanPossibility(row string, idx int, cleaned bool) (bool, bool) {
	rowLen := len(row)
	toCheck := min(idx, rowLen-idx-2)
	// fmt.Println(row, idx)

	for i:=0; i<=toCheck; i++ {
		left := row[idx-i:idx-i+1]
		right := row[idx+i+1:idx+i+2]
		// fmt.Printf("%s:%s\n", left, right)
		if left != right {
			if cleaned {
				return false, true
			}
			cleaned = true
		}
	}
	return true, cleaned
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func part1(input string) int {
// 	result := 0
// 	blocks := strings.Split(input, "\n\n")

// 	BlockLoop:
// 	for _, block := range blocks {
// 		horizontalIndex := getHorizontalIndex(block)
// 		if horizontalIndex > 0 {
// 			result += 100*horizontalIndex
// 			continue BlockLoop
// 		}
// 		verticalIndex := getVerticalIndex(block)
// 		if verticalIndex < 0 {
// 			panic("no horizontal or vertical mirror found!")
// 		}
// 		result += verticalIndex
// 	}

// 	return result
// }

// func isMirroredAfter(row string, idx int) bool {
// 	rowLen := len(row)
// 	toCheck := min(idx, rowLen-idx-2)
	
// 	for i:=0; i<=toCheck; i++ {
// 		if row[idx-i:idx-i+1] != row[idx+i+1:idx+i+2] {
// 			return false
// 		}
// 	}
// 	return true
// }