package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pos struct {
	x, y, z int
}

type brick struct {
	startPos    pos
	endPos      pos
	hasMoved    bool
	bricksBelow []*brick
	bricksAbove []*brick
}

func (b brick) isOnGround() bool {
	return b.startPos.z == 1 || b.endPos.z == 1
}

func (b brick) getDirection() string {
	if b.startPos.x != b.endPos.x {
		return "x"
	}
	if b.startPos.y != b.endPos.y {
		return "y"
	}
	return "z"
}

func (b brick) isOnPos(targetPos pos) bool {
	return b.startPos.x <= targetPos.x && targetPos.x <= b.endPos.x &&
		b.startPos.y <= targetPos.y && targetPos.y <= b.endPos.y &&
		b.startPos.z <= targetPos.z && targetPos.z <= b.endPos.z
}

func (b *brick) moveDownOne() {
	b.startPos.z -= 1
	b.endPos.z -= 1
	// fmt.Println("moved")
}

func (b *brick) addBrickBelow(targetBrick *brick) {
	if b.bricksBelow != nil {
		b.bricksBelow = removeDuplicate(append(b.bricksBelow, targetBrick))
	} else {
		b.bricksBelow = []*brick{targetBrick}
	}
}

func (b *brick) addBrickAbove(targetBrick *brick) {
	if b.bricksAbove != nil {
		b.bricksAbove = removeDuplicate(append(b.bricksAbove, targetBrick))
	} else {
		b.bricksAbove = []*brick{targetBrick}
	}
}

func (b brick) canBeDisintegrated() bool {
	// if no bricks above, can be disintegrated
	if b.bricksAbove == nil || len(b.bricksAbove) == 0 {
		return true
	} else {
		// for all brick above, if there are >1 bricks below, can be disintegrated
		for _, brickAbove := range b.bricksAbove {
			if len(brickAbove.bricksBelow) == 1 {
				return false
			}
		}
	}
	return true
}

// 78738 too high
// 50970 too low

func main() {
	data, _ := os.ReadFile("./input.txt")
	input := string(data)

	bricks := convertInputToBricks(input)
	sortBricksByStartPosZ(bricks)

	for i := 0; i < len(bricks); i++ {
		moveBrickDown(i, bricks)
	}

	for i := 0; i < len(bricks); i++ {
		setBelowBricks(i, bricks)
		setAboveBricks(i, bricks)
	}

	part1Ans := getNumBricksThatCanBeDisintegrated(bricks)
	fmt.Println("Part1 Ans: ", part1Ans)

	part2Ans := getTotalBricksThatWilLFall(input, len(bricks))
	fmt.Println("Part2 Ans: ", part2Ans)
}

func getTotalBricksThatWilLFall(input string, numBricks int) int {
	// var bricks []*brick
	total := 0
	// totalDisintegrated := 0
	// ch := make(chan int, 10)


	//re-create all bricks every time
	for n:=0; n<numBricks; n++ {
		// go func(n int) {

			fmt.Println("go num ",n)
			// goTotal := 0
			bricks := convertInputToBricks(input)
			sortBricksByStartPosZ(bricks)

			for i := 0; i < len(bricks); i++ {
				moveBrickDown(i, bricks)
			}

			for i := 0; i < len(bricks); i++ {
				setBelowBricks(i, bricks)
				setAboveBricks(i, bricks)
			}

			bricks = append(bricks[:n], bricks[n+1:]...)
			for j := 0; j < len(bricks); j++ {
				if moveBrickDown(j, bricks) {
					total += 1
				}
			}
			// ch <- goTotal
		}
	
	// for n:=0; n<numBricks; n++ {
	// 	fmt.Println("result num ",n)
	// 	total += <- ch
	// }	
	// fmt.Println("total disintegrated: ", totalDisintegrated)
	return total
}

// move brick downwards one step at a time until ground or hits another brick
func moveBrickDown(idx int, bricks []*brick) bool {
	currBrick := bricks[idx]
	moved := false
mainLoop:
	for {
		if currBrick.isOnGround() {
			break mainLoop
		}
		switch currBrick.getDirection() {
		case "z": // vertical
			bottomPos := pos{currBrick.startPos.x, currBrick.startPos.y, currBrick.startPos.z - 1}
			if !hasBrickOnPos(bricks, bottomPos) {
				currBrick.moveDownOne()
				moved = true
			} else {
				break mainLoop
			}
		case "y":
			canMove := true
		yLoop:
			for y := currBrick.startPos.y; y <= currBrick.endPos.y; y++ {
				bottomPos := pos{currBrick.startPos.x, y, currBrick.startPos.z - 1}
				if hasBrickOnPos(bricks, bottomPos) {
					canMove = false
					break yLoop
				}
			}
			if canMove {
				currBrick.moveDownOne()
				moved = true
			} else {
				break mainLoop
			}
		case "x":
			canMove := true
		xLoop:
			for x := currBrick.startPos.x; x <= currBrick.endPos.x; x++ {
				bottomPos := pos{x, currBrick.startPos.y, currBrick.startPos.z - 1}
				if hasBrickOnPos(bricks, bottomPos) {
					canMove = false
					break xLoop
				}
			}
			if canMove {
				currBrick.moveDownOne()
				moved = true
			} else {
				break mainLoop
			}
		}
	}
	return moved
}

func getNumBricksThatCanBeDisintegrated(bricks []*brick) int {
	total := 0
	for _, brick := range bricks {
		if brick.canBeDisintegrated() {
			total += 1
		}
	}
	return total
}

func setBelowBricks(idx int, bricks []*brick) {
	currBrick := bricks[idx]
	if !currBrick.isOnGround() {
		switch currBrick.getDirection() {
		case "z":
			bottomPos := pos{currBrick.startPos.x, currBrick.startPos.y, currBrick.startPos.z - 1}
			bottomBrick := getBrickOnPos(bricks, bottomPos)
			if bottomBrick != nil {
				currBrick.addBrickBelow(bottomBrick)
			}
		case "y":
			for y := currBrick.startPos.y; y <= currBrick.endPos.y; y++ {
				bottomPos := pos{currBrick.startPos.x, y, currBrick.startPos.z - 1}
				bottomBrick := getBrickOnPos(bricks, bottomPos)
				if bottomBrick != nil {
					currBrick.addBrickBelow(bottomBrick)
				}
			}
		case "x":
			for x := currBrick.startPos.x; x <= currBrick.endPos.x; x++ {
				bottomPos := pos{x, currBrick.startPos.y, currBrick.startPos.z - 1}
				bottomBrick := getBrickOnPos(bricks, bottomPos)
				if bottomBrick != nil {
					currBrick.addBrickBelow(bottomBrick)
				}
			}
		}
	}
}

func setAboveBricks(idx int, bricks []*brick) {
	currBrick := bricks[idx]
	switch currBrick.getDirection() {
	case "z":
		topPos := pos{currBrick.endPos.x, currBrick.endPos.y, currBrick.endPos.z + 1}
		topBrick := getBrickOnPos(bricks, topPos)
		if topBrick != nil {
			currBrick.addBrickAbove(topBrick)
		}
	case "y":
		for y := currBrick.startPos.y; y <= currBrick.endPos.y; y++ {
			topPos := pos{currBrick.startPos.x, y, currBrick.startPos.z + 1}
			topBrick := getBrickOnPos(bricks, topPos)
			if topBrick != nil {
				currBrick.addBrickAbove(topBrick)
			}
		}
	case "x":
		for x := currBrick.startPos.x; x <= currBrick.endPos.x; x++ {
			topPos := pos{x, currBrick.startPos.y, currBrick.startPos.z + 1}
			topBrick := getBrickOnPos(bricks, topPos)
			if topBrick != nil {
				currBrick.addBrickAbove(topBrick)
			}
		}
	}
}

func sortBricksByStartPosZ(bricks []*brick) {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].startPos.z < bricks[j].startPos.z
	})
}

func convertInputToBricks(input string) []*brick {
	rows := strings.Split(input, "\n")
	bricks := make([]*brick, len(rows))
	for i, row := range rows {
		parts := strings.Split(row, "~")

		bricks[i] = &brick{startPos: convertStringPosToPos(parts[0]), endPos: convertStringPosToPos(parts[1])}
	}
	return bricks
}

func convertStringPosToPos(posStr string) pos {
	valsStr := strings.Split(posStr, ",")
	x, _ := strconv.Atoi(valsStr[0])
	y, _ := strconv.Atoi(valsStr[1])
	z, _ := strconv.Atoi(valsStr[2])
	return pos{x, y, z}
}

func printBricks(bricks []*brick) {
	for _, brick := range bricks {
		fmt.Println(brick)
	}
}

func hasBrickOnPos(bricks []*brick, targetPos pos) bool {
	for _, brick := range bricks {
		if brick.isOnPos(targetPos) {
			return true
		}
	}
	return false
}

func getBrickOnPos(bricks []*brick, targetPos pos) *brick {
	for _, brick := range bricks {
		if brick.isOnPos(targetPos) {
			return brick
		}
	}
	return nil
}

func removeDuplicate(bricks []*brick) []*brick {
	allKeys := make(map[*brick]bool)
	list := []*brick{}
	for _, item := range bricks {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
