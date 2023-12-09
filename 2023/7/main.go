package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const sortedCards = "AKQT98765432J"

// const sortedCardsJoker = "AKQT98765432J"

type hand struct {
	originalString string
	mappedCards    map[string]int
	bid            int
	numJokers      int
}

// func (h hand) getRank() int {
// 	hasTriplet := false
// 	hasPair := false

// 	for _, v := range h.mappedCards {
// 		if v == 5 {
// 			return 7 // five of a kind
// 		}
// 		if v == 4 {
// 			return 6 // four of a kind
// 		}
// 		if v == 3 {
// 			if hasPair {
// 				return 5 // Full house
// 			}
// 			hasTriplet = true
// 		}
// 		if v == 2 {
// 			if hasTriplet {
// 				return 5 // Full house
// 			}
// 			if hasPair {
// 				return 3 // two pairs
// 			}
// 			hasPair = true
// 		}
// 	}
// 	if hasTriplet {
// 		return 4 // three of a kind
// 	}
// 	if hasPair {
// 		return 2 // one pair
// 	}
// 	return 1 // high card
// }

func (h hand) getRank() int {
	hasTriplet := false
	hasPair := false
	if h.numJokers == 5 {
		return 7
	}

	for _, v := range h.mappedCards {
		if v == 5 {
			return 7 // five of a kind
		}
		if v == 4 {
			return 6 + h.numJokers // four of a kind
		}
		if v == 3 {
			if hasPair {
				return 5 // Full house
			}
			if h.numJokers > 0 {
				return 5 + h.numJokers
			}
			hasTriplet = true
		}
		if v == 2 {
			if hasTriplet {
				return 5 // Full house
			}
			if hasPair {
				if h.numJokers == 1 {
					return 5 // full house
				} else {
					return 3 // two pairs
				}
			}
			hasPair = true
		}
	}
	if hasTriplet { // jokers taken care of alread
		return 4 // three of a kind
	}
	if hasPair { // only one pair
		if h.numJokers == 3 {
			return 7 // five of a kind
		}
		if h.numJokers == 2 {
			return 6 // four of a kind
		}
		if h.numJokers == 1 {
			return 4 // three of a kind
		}
		return 2 // one pair
	}

	if h.numJokers == 4 {
		return 7
	}
	if h.numJokers == 3 {
		return 6
	}
	if h.numJokers == 2 {
		return 4
	}
	if h.numJokers == 1 {
		return 2
	}
	return 1 // high card
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

// return if hand1 > hand2
func sortTwoHands(hand1, hand2 hand) bool {
	if hand1.getRank() > hand2.getRank() {
		return true
	}
	if hand1.getRank() < hand2.getRank() {
		return false
	}
	return sortTwoHandsWithSameRank(hand1, hand2)
}

func sortTwoHandsWithSameRank(hand1, hand2 hand) bool {
	for i := 0; i < len(hand1.originalString); i++ {
		if strings.Index(sortedCards, string(hand1.originalString[i])) < strings.Index(sortedCards, string(hand2.originalString[i])) {
			return true
		}
		if strings.Index(sortedCards, string(hand1.originalString[i])) > strings.Index(sortedCards, string(hand2.originalString[i])) {
			return false
		}
	}
	return true
}

func handStrToHandStruct(handStr string) hand {
	handSlice := strings.Fields(handStr)
	bidInt, _ := strconv.Atoi(handSlice[1])
	mapped := map[string]int{}

	numJokers := 0

	for _, card := range strings.Split(handSlice[0], "") {
		if card == "J" {
			numJokers += 1
		} else {
			mapped[card] += 1
		}
	}

	return hand{mappedCards: mapped, bid: bidInt, originalString: handSlice[0], numJokers: numJokers}
}

func part1(input string) int {
	result := 0
	rows := strings.Split(input, "\n")
	hands := []hand{}
	// handsByRank := []hand{}

	for _, row := range rows {
		// fmt.Printf("Row %d: %s", i, row)
		hands = append(hands, handStrToHandStruct(row))
	}
	// fmt.Println("hands: ", hands)

	// for _, hand := range hands {
	// 	fmt.Printf("hand: %+v, rank: %d\n", hand, hand.getRank())
	// }
	sort.Slice(hands, func(i, j int) bool {
		return sortTwoHands(hands[i], hands[j])
	})

	// fmt.Println("sorted hands: ", hands)

	for i, hand := range hands {
		multiplier := len(hands) - (i)
		fmt.Printf("multipler: %d, hand: %+v\n", multiplier, hand)
		result += hand.bid * multiplier
	}
	return result
}

func part2(input string) int {
	result := 0
	rows := strings.Split(input, "\n")
	hands := []hand{}
	// handsByRank := []hand{}

	for _, row := range rows {
		// fmt.Printf("Row %d: %s", i, row)
		hands = append(hands, handStrToHandStruct(row))
	}
	// fmt.Println("hands: ", hands)

	// for _, hand := range hands {
	// 	fmt.Printf("hand: %+v, rank: %d\n", hand, hand.getRank())
	// }
	sort.Slice(hands, func(i, j int) bool {
		return sortTwoHands(hands[i], hands[j])
	})

	// fmt.Println("sorted hands: ", hands)

	for i, hand := range hands {
		multiplier := len(hands) - (i)
		fmt.Printf("multipler: %d, hand: %+v, rank:%d\n", multiplier, hand, hand.getRank())
		result += hand.bid * multiplier
	}
	return result
}
