package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const numMaps = 7

type seedMap struct {
	destStart   int
	sourceStart int
	steps       int
}

func (sm seedMap) getSourceEnd() int {
	return sm.sourceStart + sm.steps
}

type newSeed struct {
	start int
	steps int
}

func main() {
	// data, err := os.ReadFile("./input.txt")
	// if err != nil {
	// log.Fatalln("error reading file: ", err)
	// }
	// inputRaw := string(data)
	// fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2())
}

func part1(input string) int {
	result := 0
	seeds := []int{}
	locations := []int{}

	seedsData, _ := os.ReadFile("./seeds.txt")
	for _, seedString := range strings.Fields(string(seedsData)) {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	maps := make([][]seedMap, numMaps)
	for i := 1; i <= numMaps; i++ {
		maps[i-1] = getSeedMapList(fmt.Sprintf("./%d.txt", i))
	}

	for _, seed := range seeds {
		// fmt.Println("seed: ", seed)
		source := seed

	SeedMapLoop:
		for _, seedMapList := range maps {
			// fmt.Println("source, map: ", source, seedMapList)
			for _, seedMap := range seedMapList {
				if source >= seedMap.sourceStart && source <= seedMap.getSourceEnd() {
					source = seedMap.destStart + (source - seedMap.sourceStart)
					continue SeedMapLoop
				}
			}
		}
		locations = append(locations, source)
	}

	fmt.Println(locations)
	result = getSmallestNum(locations)

	return result
}

func getSmallestNum(nums []int) int {
	smallest := math.MaxInt
	for _, num := range nums {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}

func getSeedMapList(fileName string) []seedMap {
	result := []seedMap{}
	data, _ := os.ReadFile(fileName)
	for _, row := range strings.Split(string(data), "\n") {
		data := strings.Fields(row)
		destStart, _ := strconv.Atoi(data[0])
		sourceStart, _ := strconv.Atoi(data[1])
		count, _ := strconv.Atoi(data[2])

		result = append(result, seedMap{destStart, sourceStart, count})
	}
	return result
}

func part2() int {
	seeds := []newSeed{}

	seedsData, _ := os.ReadFile("./seeds.txt")
	seedsDataSlice := strings.Fields(string(seedsData))
	for j := 0; j < len(seedsDataSlice); j += 2 {
		start, _ := strconv.Atoi(seedsDataSlice[j])
		steps, _ := strconv.Atoi(seedsDataSlice[j+1])
		seeds = append(seeds, newSeed{start, steps})
	}

	maps := make([][]seedMap, numMaps)
	for i := 1; i <= numMaps; i++ {
		maps[i-1] = getSeedMapList(fmt.Sprintf("./%d.txt", i))
	}

	smallestLocation := math.MaxInt
	ch := make(chan int)
	for _, newSeed := range seeds {
		go getSmallestLocationPerNewSeed(newSeed, maps, ch)
	}

	for k := 0; k < len(seeds); k++ {
		res := <-ch
		if res < smallestLocation {
			smallestLocation = res
		}
	}

	// below logic was before parallelism using go routines

	// smallestLocation := math.MaxInt
	// for i, newSeed := range seeds {
	// 	fmt.Println("newSeed numer: ", i)
	// 		for k := 0; k<newSeed.steps; k++ {
	// 		source := newSeed.start+k
	// 		SeedMapLoop:
	// 		for _, seedMapList := range maps {
	// 			// fmt.Println("source, map: ", source, seedMapList)
	// 			for _, seedMap := range seedMapList {
	// 				if source >= seedMap.sourceStart && source <= seedMap.getSourceEnd() {
	// 					source = seedMap.destStart + (source-seedMap.sourceStart)
	// 					continue SeedMapLoop
	// 				}
	// 			}
	// 		}
	// 		if source < smallestLocation {
	// 			smallestLocation = source
	// 		}
	// 	}
	// }

	return smallestLocation
}

func getSmallestLocationPerNewSeed(newSeed newSeed, maps [][]seedMap, ch chan int) {
	smallestLocation := math.MaxInt
	for k := 0; k < newSeed.steps; k++ {
		source := newSeed.start + k
	SeedMapLoop:
		for _, seedMapList := range maps {
			for _, seedMap := range seedMapList {
				if source >= seedMap.sourceStart && source <= seedMap.getSourceEnd() {
					source = seedMap.destStart + (source - seedMap.sourceStart)
					continue SeedMapLoop
				}
			}
		}
		if source < smallestLocation {
			smallestLocation = source
		}
	}
	ch <- smallestLocation
}
