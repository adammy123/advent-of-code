package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

const startingRuleName = "in"

var rulesMap map[string][]condition

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	fmt.Println("part1 ans: ", part1(inputRaw))
	fmt.Println("part2 ans: ", part2(inputRaw))
}

func part2(input string) int {
	inputParts := strings.Split(input, "\n\n")
	rulesMap = getRulesMap(inputParts[0])
	allCombis := []partCombis{{[][]int{{1, 4000},{1, 4000},{1, 4000},{1, 4000}}}}
	initialConditions := rulesMap[startingRuleName]

	return getApprovedCombis(allCombis, initialConditions)
}

func getApprovedCombis(combis []partCombis, conditions []condition) int {
	if len(combis) == 0 {
		return 0
	}
	if len(conditions) == 1 {
		return getCombiConditionResult(conditions[0].result, combis)
	}

	thisCondition := conditions[0]
	remainingConditions := conditions[1:]

	trueCombis, falseCombis := splitCombis(combis, thisCondition)
	return getCombiConditionResult(thisCondition.result, trueCombis) + getApprovedCombis(falseCombis, remainingConditions)
}

var categoryMap map[string]int = map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}

func splitCombis(combis []partCombis, condition condition) (trueResults []partCombis, falseResults []partCombis) {
	// fmt.Println("Condition: ", condition)
	for _, combi := range combis {
		idx := categoryMap[condition.category]
		originalVals := combi.values[idx]
		var trueVal, falseVal []int
		if condition.operator == "<" {
			if originalVals[0] < condition.threshold {
				if originalVals[1] < condition.threshold {
					trueVal = originalVals
				} else {
					trueVal = []int{originalVals[0], condition.threshold-1}
					falseVal = []int{condition.threshold, originalVals[1]}
				}
			} else {
				falseVal = originalVals
			}
		} else { // operator == ">"
			if originalVals[1] > condition.threshold {
				if originalVals[0] > condition.threshold {
					trueVal = originalVals
				} else {
					trueVal = []int{condition.threshold+1, originalVals[1]}
					falseVal = []int{originalVals[0], condition.threshold}
				}
			} else {
				falseVal = originalVals
			}
		}

		// todo: don't make results if val is empty
		if trueVal != nil {
			trueVals := make([][]int, 4)
			for j:=0; j<4; j++ {
				if j==idx {
					trueVals[j] = trueVal
				} else {
					trueVals[j] = combi.values[j]
				}
			}
			trueResults = append(trueResults, partCombis{trueVals})
		}
		if falseVal != nil {
			falseVals := make([][]int, 4)
			for j:=0; j<4; j++ {
				if j==idx {
					falseVals[j] = falseVal
				} else {
					falseVals[j] = combi.values[j]
				}
			}
			falseResults = append(falseResults, partCombis{falseVals})
		}
	}
	// fmt.Println("trueResults: ", trueResults)
	// fmt.Println("falseResults: ", falseResults)
	return trueResults, falseResults
}

func getCombiConditionResult(result string, combis []partCombis) int {
	switch result {
	case "A":
		total := 0
		for _, combi := range combis {
			total += combi.getNumPossibleCombis()
			// fmt.Printf("%+v +%d\n", combi, combi.getNumPossibleCombis())
		}
		return total
	case "R":
		return 0
	default:
		return getApprovedCombis(combis, rulesMap[result])
	}
}

// [lowestNum, highestNum] inclusive
type partCombis struct {
	values [][]int
}

func (pc partCombis) getNumPossibleCombis() int {
	return (pc.values[0][1]-pc.values[0][0]+1)*(pc.values[1][1]-pc.values[1][0]+1)*(pc.values[2][1]-pc.values[2][0]+1)*(pc.values[3][1]-pc.values[3][0]+1)
}

func part1(input string) int {
	result := 0
	inputParts := strings.Split(input, "\n\n")

	rulesMap = getRulesMap(inputParts[0])
	// for k, v := range rulesMap {
	// 	fmt.Println(k, v)
	// }
	parts := getPartsList(inputParts[1])
	// for _, part := range parts {
	// 	fmt.Println(part)
	// }

	for _, part := range parts {
		isAccepted := putPartThroughRule(part, startingRuleName)
		if isAccepted {
			result += part.getSumRatings()
		}
	}

	return result
}

func putPartThroughRule(p part, ruleName string) bool {
	conditions := rulesMap[ruleName]

	conditionLoop:
	for _, condition := range conditions {
		if condition.operator == "" {
			return getConditionResult(p, condition.result)
		}
		var partVal int
		switch condition.category {
		case "x":
			partVal = p.x
		case "m":
			partVal = p.m
		case "a":
			partVal = p.a
		case "s":
			partVal = p.s
		}
		if condition.operator == ">" {
			if partVal > condition.threshold {
				return getConditionResult(p, condition.result)
			}
			continue conditionLoop
		} else {
			if partVal < condition.threshold {
				return getConditionResult(p, condition.result)
			}
			continue conditionLoop
		}
	}
	return false //should not reach here
}

func getConditionResult(p part, result string) bool {
	switch result {
	case "A":
		return true
	case "R":
		return false
	default:
		return putPartThroughRule(p, result)
	}
}

type condition struct {
	category string
	operator string
	threshold int
	result string
}

type part struct {
	x, m, a, s int
}

func (p part) getSumRatings() int {
	return p.x + p.m + p.a + p.s
}

func getRulesMap(input string) map[string][]condition {
	rulesMap := map[string][]condition{}
	for _, row := range strings.Split(input, "\n") {
		fields := strings.Split(row, "{")
		ruleName := fields[0]
		rulesString := strings.ReplaceAll(fields[1], "}", "")
		rulesSlice := strings.Split(rulesString, ",")

		conditions := make([]condition, len(rulesSlice))
		for i:=0; i<len(rulesSlice); i++ {
			ruleString := rulesSlice[i]
			if i==len(rulesSlice)-1 {
				conditions[i] = condition{result: ruleString}
			} else {
				conditionFields := strings.Split(ruleString, ":")
				result := conditionFields[1]
				left := conditionFields[0]
				var op string
				if strings.Contains(left, ">") {
					op = ">"
				} else {
					op = "<"
				}
				thisRuleFields := strings.Split(left, op)
				category := thisRuleFields[0]
				threshold, _ := strconv.Atoi(thisRuleFields[1])
				conditions[i] = condition{category, op, threshold, result}

			}
		}
		rulesMap[ruleName] = conditions
	}
	return rulesMap
}

func getPartsList(input string) []part {
	rows := strings.Split(input, "\n")
	parts := make([]part, len(rows))

	re := regexp.MustCompile("[0-9]+")

	for i:=0; i<len(rows); i++ {
		nums := re.FindAllString(rows[i], -1)
		x, _ := strconv.Atoi(nums[0])
		m, _ := strconv.Atoi(nums[1])
		a, _ := strconv.Atoi(nums[2])
		s, _ := strconv.Atoi(nums[3])
		parts[i] = part{x, m, a, s}
	}
	return parts
}


