package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

var (
	cMap = map[string][]string{}
)

// represent if direct connection between components c1 and c2 are cut,
// how many steps does it take to reach each other now
type conn struct {
	c1, c2 string
	newDist int
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	initComponents(inputRaw)
	// for k, v := range cMap {
	// 	fmt.Println(k, v)
	// }
	fmt.Println("part1 ans: ", part1())
	// fmt.Println("part2 ans: ", part2(inputRaw))
}

func part1() int {
	conns := []conn{}
	for comp, neighbours := range cMap {
		for _, neighbour := range neighbours {
			if notCalculated(conns, comp, neighbour) {
				conns = append(conns, conn{c1: comp, c2: neighbour, newDist: findMinDist(comp, neighbour)})
			}
		}
	}

	sort.Slice(conns, func(i, j int) bool {
		return conns[i].newDist > conns[j].newDist
	})
	// for _, conn := range conns {
	// 	fmt.Println(conn)
	// }

	for i:=0; i<3; i++ {
		conn := conns[i]
		cMap[conn.c1] = removeStringFromSlice(cMap[conn.c1], conn.c2)
		cMap[conn.c2] = removeStringFromSlice(cMap[conn.c2], conn.c1)
	}

	for k, v := range cMap {
		fmt.Println(k, v)
	}

	header1 := conns[0].c1
	header2 := conns[0].c2

	return calculateSize(header1)*calculateSize(header2)
}

func calculateSize(start string) int {
	fmt.Println("calculating size for starting component: ", start)
	visited := []string{start}
	toVisit := cMap[start]
	var newToVisit []string

	for {
		fmt.Println("toVisit: ", toVisit)
		visited = append(visited, toVisit...)
		newToVisit = []string{}
		for _, n := range toVisit {
			for _, nn := range cMap[n] {
				 
				if !slices.Contains(visited, nn) && !slices.Contains(newToVisit, nn) {
					newToVisit = append(newToVisit, nn)
				}
			}
		}
		if len(newToVisit) == 0 {
			fmt.Println("size: ", len(visited))
			fmt.Println(visited)
			fmt.Println()
			return len(visited)
		}
		toVisit = newToVisit
	}
}

func findMinDist(start, end string) int {
	// fmt.Println("finding new min dist between: ", start, end)
	visited := []string{start}
	steps := 1
	var newNeighbours []string
	neighbours := removeStringFromSlice(cMap[start], end)

	Loop:
	for {
		newNeighbours = []string{}
		visited = append(visited, neighbours...)

		for _, n := range neighbours {
			if n == end {
				break Loop
			}
			for _, nn := range cMap[n] {
				if !slices.Contains(visited, nn) {
					newNeighbours = append(newNeighbours, nn)
				}
			}
		}
		neighbours = newNeighbours
		steps += 1
	}
	
	return steps
}

func notCalculated(conns []conn, c1, c2 string) bool {
	for _, conn := range conns {
		if (conn.c1 == c1 && conn.c2 == c2) || (conn.c1 == c2 && conn.c2 == c1) {
			return false
		}
	}
	return true
}

func removeStringFromSlice(neighbours []string, end string) []string {
	new := []string{}
	for _, n := range neighbours {
		if n != end {
			new = append(new, n)
		}
	}
	return new
}

func initComponents(input string) {
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		fields := strings.Split(row, ":")
		thisKey := fields[0]
		thisValues := strings.Fields(fields[1])
		if _, ok := cMap[thisKey]; !ok {
			cMap[thisKey] = thisValues
		} else {
			cMap[thisKey] = append(cMap[thisKey], thisValues...)
		}
		for _, value := range thisValues {
			if _, ok := cMap[value]; !ok {
				cMap[value] = []string{thisKey}
			} else {
				cMap[value] = append(cMap[value], thisKey)
			}
		}
	}
}
