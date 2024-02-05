package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minPos = 200000000000000.0 // inclusive
	maxPos = 400000000000000.0 // inclusive
)

var hails = []hail{}

type hail struct {
	initialPos pos
	vel velocity
	l line
}

type pos struct {
	x, y, z int
}

type velocity struct {
	x, y, z int
}

type line struct {
	m, c float64
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	convertInputToHails(inputRaw)
	// for _, hail := range hails {
	// 	fmt.Println(hail)
	// }

	fmt.Println("part1 ans: ", part1())
}

func part1() int {
	var x, y float64
	var ok bool
	result := 0
	// for each hail, check if it will intersect with every other hail
	for i:=0; i<len(hails)-1; i++ {
		for j:=i+1; j<len(hails); j++ {
			x, y, ok = getIntersectionInFuture(hails[i], hails[j])
			if ok && x >= minPos && x <= maxPos && y >= minPos && y <= maxPos {
				result += 1
			}
		}
	}
	return result
}

func getIntersectionInFuture(first, second hail) (float64, float64, bool) {
	x, y := 0., 0.
	// if parallel, return false
	if isParallel(first, second) {
		return x, y, false
	}

	// get x & y
	x = (second.l.c-first.l.c)/(first.l.m-second.l.m)
	y = first.l.m*x + first.l.c

	// fmt.Println(x, y)
	if isInPast(first, x, y) || isInPast(second, x, y) {
		// fmt.Println("in past")
		return x, y, false
	}
	// if in the past, return false
	return x, y, true
}

func isInPast(h hail, x, y float64) bool {
	if x > float64(h.initialPos.x) {
		return h.vel.x < 0
	}
	return h.vel.x > 0
}

func isParallel(first, second hail) bool {
	return first.l.m == second.l.m
}

func convertInputToHails(input string) {
	rows := strings.Split(input, "\n")
	hails = make([]hail, len(rows))
	for i, row := range rows {
		row = strings.ReplaceAll(row, " ", "")
		fields := strings.Split(row, "@")
		posStrSlice := strings.Split(fields[0], ",")
		posIntSlice := make([]int, len(posStrSlice))
		for j:=0; j<len(posStrSlice); j++ {
			intVal, _ := strconv.Atoi(posStrSlice[j])
			posIntSlice[j] = intVal
		}

		velStrSlice := strings.Split(fields[1], ",")
		velIntSlice := make([]int, len(velStrSlice))
		for j:=0; j<len(velStrSlice); j++ {
			intVal, _ := strconv.Atoi(velStrSlice[j])
			velIntSlice[j] = intVal
		}

		thisPos := pos{posIntSlice[0], posIntSlice[1], posIntSlice[2]}
		thisVel := velocity{velIntSlice[0], velIntSlice[1], velIntSlice[2]}
		thisLine := getLine(thisPos, thisVel)
		hails[i] = hail{thisPos, thisVel, thisLine}

	}
}

func getLine(pos pos, vel velocity) line {
	m := float64(vel.y)/float64(vel.x)
	c := float64(pos.y) - float64(pos.x)*m
	return line{m, c}
}
