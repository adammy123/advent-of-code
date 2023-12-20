package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	totalLow = 0
	totalHigh = 0
	modules = map[string]module{}
	rxSourceInputs = map[string]int{}
	buttonPresses = 0
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	inputRaw := string(data)
	convertInput(inputRaw)
	part1()
	fmt.Println("part2 ans: ", part2())
}

func part2() int {
	for k := range modules["lv"].(conjunction).inputs {
		rxSourceInputs[k] = -1
	}

	Loop:
	for {
		buttonPresses += 1
		dest := modules["broadcaster"].getDest()
		isPulseHigh := []bool{}
		source := []string{}
		for i:=0; i<len(dest); i++ {
			isPulseHigh = append(isPulseHigh, false)
			source = append(source, "broadcaster")
		}

		var newDest []string
		var newIsPulseHigh []bool
		var newSource []string
		
		innerLoop:
		for {
			if len(dest) == 0 {
				continue Loop
			}
			newDest = []string{}
			newIsPulseHigh = []bool{}
			newSource = []string{}
			for i, moduleName := range dest {
				thisDest, thisIsPulseHigh := sendPulse(isPulseHigh[i], moduleName, source[i])
				if len(thisDest) > 0 {
					newDest = append(newDest, thisDest...)
					for j:=0; j<len(thisDest); j++ {
						newIsPulseHigh = append(newIsPulseHigh, thisIsPulseHigh)
						newSource = append(newSource, moduleName)
					}
				}
			}
			dest = newDest
			isPulseHigh = newIsPulseHigh
			source = newSource


			for _, v := range rxSourceInputs {
				if v<0 {
					continue innerLoop
				}
			}
			return getProduct(rxSourceInputs)
		}
	}
}

func getProduct(m map[string]int) int {
	result := 1
	for _, v := range m {
		result *= v
	}
	return result
}

func part1() {
	Loop:
	for k:=0; k<1000; k++ {
		totalLow += 1
		dest := modules["broadcaster"].getDest()
		isPulseHigh := []bool{}
		source := []string{}
		for i:=0; i<len(dest); i++ {
			isPulseHigh = append(isPulseHigh, false)
			source = append(source, "broadcaster")
		}

		var newDest []string
		var newIsPulseHigh []bool
		var newSource []string
		
		for {
			if len(dest) == 0 {
				continue Loop
			}
			newDest = []string{}
			newIsPulseHigh = []bool{}
			newSource = []string{}
			for i, moduleName := range dest {
				thisDest, thisIsPulseHigh := sendPulse(isPulseHigh[i], moduleName, source[i])
				if len(thisDest) > 0 {
					newDest = append(newDest, thisDest...)
					for j:=0; j<len(thisDest); j++ {
						newIsPulseHigh = append(newIsPulseHigh, thisIsPulseHigh)
						newSource = append(newSource, moduleName)
					}
				}
			}
			dest = newDest
			isPulseHigh = newIsPulseHigh
			source = newSource
		}
	}
	fmt.Printf("Part1 ans: {totalLow: %d, totalHigh: %d} = %d\n", totalLow, totalHigh, totalLow*totalHigh)
}

//return list of dest, is new pulse high
func sendPulse(isPulseHigh bool, moduleName, source string) ([]string, bool) {
	
	if isPulseHigh {
		totalHigh += 1
	} else {
		if val, ok := rxSourceInputs[moduleName]; ok {
			if val<0 {
				rxSourceInputs[moduleName] = buttonPresses
			}
		}
		totalLow += 1
	}
	thisModule, ok := modules[moduleName]
	if !ok {
		return []string{}, false
	}
	if !isPulseHigh && thisModule.getType() == "flipflop" {
		ffModule := thisModule.(flipflop)
		ffModule.isOn = !ffModule.isOn
		modules[moduleName] = ffModule
	}
	if thisModule.getType() == "conjunction" {
		cModule := thisModule.(conjunction)
		if _, ok := cModule.inputs[source]; ok {
			cModule.inputs[source] = isPulseHigh
		}
		modules[moduleName] = cModule
	}
	if isNewPulseHigh, ok := modules[moduleName].isResultingPulseHigh(isPulseHigh); ok {
		return modules[moduleName].getDest(), isNewPulseHigh
	}
	return []string{}, false
}

func convertInput(input string) {
	for _, row := range strings.Split(input, "\n") {
		parts := strings.Split(strings.ReplaceAll(row, " ", ""), "->")
		left := parts[0]
		dest := strings.Split(parts[1], ",")
		if left == "broadcaster" {
			modules[left] = broadcaster{dest: dest}
		} else {
			typeStr := string(left[0])
			moduleName := string(left[1:])
			if typeStr == "%" {
				modules[moduleName] = flipflop{dest: dest}
			} else {
				modules[moduleName] = conjunction{dest: dest, inputs: map[string]bool{}}
			}
		}
	}

	for moduleName, module := range modules {
		for _, dest := range module.getDest() {
			if mod, ok := modules[dest]; ok && mod.getType()=="conjunction" {
				addModuleAsInput(dest, moduleName)
			}
		}
	}
}

func addModuleAsInput(conjunctionName, inputName string) {
	res := modules[conjunctionName].(conjunction)
	res.inputs[inputName] = false
}

type module interface {
	// returns ifHigh, isSending
	isResultingPulseHigh(bool) (bool, bool)
	getDest() []string
	getType() string
}

type broadcaster struct {
	dest []string
}

func (broadcaster) isResultingPulseHigh(isReceivingPulseHighbool bool) (bool, bool) {
	return isReceivingPulseHighbool, true
}

func (b broadcaster) getDest() []string {
	return b.dest
}

func (broadcaster) getType() string {
	return "broadcaster"
}

type flipflop struct {
	isOn bool
	dest []string
}

func (f flipflop) isResultingPulseHigh(isReceivingPulseHigh bool) (bool, bool) {
	if isReceivingPulseHigh {
		return false, false
	}
	return f.isOn, true
}

func (f flipflop) getDest() []string {
	return f.dest
}

func (flipflop) getType() string {
	return "flipflop"
}

type conjunction struct {
	inputs map[string]bool
	dest []string
}

func (c conjunction) isResultingPulseHigh(isReceivingPulseHigh bool) (bool, bool) {
	for _, isHigh := range c.inputs {
		if !isHigh {
			return true, true
		}
	}
	return false, true
}

func (c conjunction) getDest() []string {
	return c.dest
}

func (conjunction) getType() string {
	return "conjunction"
}