package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

const (
	rowStart = 0
	colStart = 1
)

var (
	rowEnd, colEnd           int
	gridNumRows, gridNumCols int
	listTiles                []*tile

	nodes []*tile
	edges map[*tile]map[*tile]int // map of startNode: endNode: distance
)

type pos struct {
	row, col int
}

type tile struct {
	coord     pos
	val       string
	nextTiles []*tile
}

func (t tile) isNotForest() bool {
	return t.val != "#"
}

func (t tile) isStart() bool {
	return t.coord.row == rowStart && t.coord.col == colStart
}

func (t tile) isEnd() bool {
	return t.coord.row == rowEnd && t.coord.col == colEnd
}

func (t *tile) setNextTiles(tiles []*tile) {
	t.nextTiles = tiles
}

func (t tile) getNextTilesMinusPrevious(prev *tile) []*tile {
	if prev == nil {
		return t.nextTiles
	}
	if len(t.nextTiles) == 1 {
		return []*tile{}
	}
	// assume previous is always in next tiles
	result := []*tile{}
	for _, tile := range t.nextTiles {
		if tile != prev {
			result = append(result, tile)
		}
	}
	return result
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part2(input string) int {
	gridString := convertInputToGridStringPart2(input)
	gridNumRows = len(gridString)
	gridNumCols = len(gridString[0])
	rowEnd = gridNumRows - 1
	colEnd = gridNumCols - 2

	gridTiles := convertGridStringToGridTiles(gridString)
	listTiles = convertGridStringToTilesList(gridTiles)

	nodes = []*tile{}
	convertTilesToNodes()
	for _, node := range nodes {
		fmt.Println(node)
	}

	edges = map[*tile]map[*tile]int{}
	getEdges()
	for _, edge := range edges {
		fmt.Println(edge)
	}

	visited := []*tile{}
	startNode := getStartTile(listTiles)
	return getMaxPathToEnd(startNode, visited)
}

func getMaxPathToEnd(currTile *tile, visited []*tile) int {

	if currTile.isEnd() {
		return 0
	}

	nextNodesMap := edges[currTile]
	thisNextNodes := make([]*tile, len(nextNodesMap))
	i := 0
	for k := range nextNodesMap {
		thisNextNodes[i] = k
		i++
	}
	nextNodes := []*tile{}
	for _, tile := range thisNextNodes {
		if !sliceContains(visited, tile) {
			nextNodes = append(nextNodes, tile)
		}
	}

	// no more tiles to visit, return big negative number to rule out as possible path
	if len(nextNodes) == 0 {
		return math.MinInt
	}

	visited = append(visited, currTile)
	if len(nextNodes) == 1 {
		nextNode := nextNodes[0]
		currEdgeDist := edges[currTile][nextNode]
		return currEdgeDist + getMaxPathToEnd(nextNode, visited)
	}

	nextDist := make([]int, len(nextNodes))
	for i := 0; i < len(nextNodes); i++ {
		nextNode := nextNodes[i]
		currEdgeDist := edges[currTile][nextNode]
		nextDist[i] = currEdgeDist + getMaxPathToEnd(nextNode, visited)

	}
	return slices.Max(nextDist)

}

func getEdges() {
	for _, node := range nodes {
		edges[node] = map[*tile]int{}
		// for each direction
		for _, nextTile := range node.nextTiles {
			targetNode, dist := getTargetDist(nextTile, node)
			if currDist, ok := edges[node][targetNode]; ok {
				if dist > currDist {
					edges[node][targetNode] = dist
				}
			} else {
				edges[node][targetNode] = dist
			}
		}
	}
}

func getTargetDist(currTile, prevTile *tile) (*tile, int) {
	dist := 1
	for {
		if sliceContains(nodes, currTile) {
			return currTile, dist
		}
		// this should only contain 1 tile since it is not a node
		nextTiles := currTile.getNextTilesMinusPrevious(prevTile)
		if len(nextTiles) > 1 {
			panic("nextTiles >1!")
		}
		prevTile = currTile
		currTile = nextTiles[0]
		dist += 1
	}

}

func convertTilesToNodes() {
	nodes = append(nodes, getStartTile(listTiles))
	for _, tile := range listTiles {
		if len(tile.nextTiles) > 2 {
			nodes = append(nodes, tile)
		}
	}
	nodes = append(nodes, getEndTile(listTiles))
}

func part1(input string) int {
	gridString := convertInputToGridString(input)
	gridNumRows = len(gridString)
	gridNumCols = len(gridString[0])
	rowEnd = gridNumRows - 1
	colEnd = gridNumCols - 2

	gridTiles := convertGridStringToGridTiles(gridString)
	listTiles = convertGridStringToTilesList(gridTiles)

	visited := []*tile{}
	startTile := getStartTile(listTiles)

	return getMaxStepsToEnd(startTile, visited)
}

func getMaxStepsToEnd(currTile *tile, visited []*tile) int {
	// fmt.Println("Current Tile: ", currTile.coord)

	if currTile.isEnd() {
		return 0
	}

	nextTiles := []*tile{}
	for _, tile := range currTile.nextTiles {
		if !sliceContains(visited, tile) {
			nextTiles = append(nextTiles, tile)
		}
	}

	// no more tiles to visit, return big negative number to rule out as possible path
	if len(nextTiles) == 0 {
		return math.MinInt
	}

	visited = append(visited, currTile)
	if len(nextTiles) == 1 {
		return 1 + getMaxStepsToEnd(nextTiles[0], visited)
	}

	return max(1+getMaxStepsToEnd(nextTiles[0], visited), 1+getMaxStepsToEnd(nextTiles[1], visited))
}

func sliceContains(slice []*tile, target *tile) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}

func getStartTile(tiles []*tile) *tile {
	for _, tile := range tiles {
		if tile.isStart() {
			return tile
		}
	}
	panic("no start tile found") //should not reach here
}

func getEndTile(tiles []*tile) *tile {
	for _, tile := range tiles {
		if tile.isEnd() {
			return tile
		}
	}
	panic("no end tile found") //should not reach here
}

func convertGridStringToTilesList(grid [][]*tile) []*tile {
	tiles := []*tile{}
	for _, row := range grid {
		for _, tile := range row {
			if tile.val != "#" {
				setNextTiles(tile, grid)
				tiles = append(tiles, tile)
			}
		}
	}
	return tiles
}

func setNextTiles(currTile *tile, gridTiles [][]*tile) {
	nextTiles := []*tile{}
	currRowIdx := currTile.coord.row
	currColIdx := currTile.coord.col
	var candidateTile *tile
	// set top
	if currRowIdx != 0 {
		candidateTile = gridTiles[currRowIdx-1][currColIdx]
		if candidateTile.isNotForest() && candidateTile.val != "v" {
			nextTiles = append(nextTiles, candidateTile)
		}
	}
	// set bottom
	if currRowIdx < gridNumRows-1 {
		candidateTile = gridTiles[currRowIdx+1][currColIdx]
		if candidateTile.isNotForest() && candidateTile.val != "^" {
			nextTiles = append(nextTiles, candidateTile)
		}
	}
	// set left
	if currColIdx != 0 {
		candidateTile = gridTiles[currRowIdx][currColIdx-1]
		if candidateTile.isNotForest() && candidateTile.val != ">" {
			nextTiles = append(nextTiles, candidateTile)
		}
	}
	// set right
	if currColIdx < gridNumCols-1 {
		candidateTile = gridTiles[currRowIdx][currColIdx+1]
		if candidateTile.isNotForest() && candidateTile.val != "<" {
			nextTiles = append(nextTiles, candidateTile)
		}
	}
	currTile.setNextTiles(nextTiles)
}

func convertGridStringToGridTiles(grid [][]string) [][]*tile {
	gridTiles := make([][]*tile, len(grid))
	for rowIdx, row := range grid {
		gridRow := make([]*tile, len(row))
		for colIdx, val := range row {
			gridRow[colIdx] = &tile{coord: pos{row: rowIdx, col: colIdx}, val: val}
		}
		gridTiles[rowIdx] = gridRow
	}
	return gridTiles
}

func convertInputToGridString(input string) [][]string {
	rows := strings.Split(input, "\n")
	grid := make([][]string, len(rows))
	for rowIdx, row := range rows {
		gridRow := make([]string, len(row))
		for colIdx, val := range strings.Split(row, "") {
			gridRow[colIdx] = val
		}
		grid[rowIdx] = gridRow
	}
	return grid
}

func convertInputToGridStringPart2(input string) [][]string {
	rows := strings.Split(input, "\n")
	grid := make([][]string, len(rows))
	for rowIdx, row := range rows {
		gridRow := make([]string, len(row))
		for colIdx, val := range strings.Split(row, "") {
			if val == ">" || val == "^" || val == "v" || val == "<" {
				val = "."
			}
			gridRow[colIdx] = val
		}
		grid[rowIdx] = gridRow
	}
	return grid
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
