package main

import (
	"bufio"
	"fmt"
	"github.com/UltraGnome/AdventOfCode2024/pkg/sti"
	"log"
	"os"
	"strconv"
	"strings"
)

type Calculation struct {
	targetAnswer int64
	components   []int64
}

func main() {

	//candidateCalculations := readInput()
	//validCalculations := []Calculation{}
	//
	//for _, candidate := range candidateCalculations {
	//	if isSolvable(candidate.targetAnswer, candidate.components) {
	//		validCalculations = append(validCalculations, candidate)
	//	}
	//}
	//
	//var total int64 = 0
	//for _, validCalculation := range validCalculations {
	//	total += validCalculation.targetAnswer
	//}
	//
	//fmt.Println(total)
	day7_2(readInputAsString())

}

func isSolvable(target int64, components []int64) bool {
	lastItemIdx := len(components) - 1
	result := false
	if len(components) == 1 {
		result = (target == components[0])
		return result
	}
	if target%components[lastItemIdx] == 0 && isSolvable((target/components[lastItemIdx]), components[:lastItemIdx]) {
		result = true
		return result
	}
	if target > components[lastItemIdx] && isSolvable((target-components[lastItemIdx]), components[:lastItemIdx]) {
		result = true
		return result
	}
	targetString := strconv.FormatInt(target, 10)
	lastComponentString := strconv.FormatInt(components[lastItemIdx], 10)
	concatenatedString := targetString + lastComponentString
	concatInt64, _ := strconv.ParseInt(concatenatedString, 10, 64)
	if strings.HasSuffix(targetString, lastComponentString) && len(targetString) > len(lastComponentString) && isSolvable(concatInt64, components[:lastItemIdx]) {
		result = true
		return result
	}
	return result
}

// borrowed because concatenation beat me
func day7_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	var evaluate func(target int, numbers []int, index int, currentResult int) bool
	evaluate = func(target int, numbers []int, index int, currentResult int) bool {
		if index >= len(numbers) {
			return currentResult == target
		}

		currentNumber := numbers[index]

		newResult := currentResult + currentNumber
		if newResult <= target {
			if evaluate(target, numbers, index+1, newResult) {
				return true
			}
		}

		newResult = currentResult * currentNumber
		if newResult <= target {
			if evaluate(target, numbers, index+1, newResult) {
				return true
			}
		}

		concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentResult, currentNumber))
		if concatenated <= target {
			correct := evaluate(target, numbers, index+1, concatenated)
			return correct
		}

		return false
	}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(parts[0])
		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))

		for i, numString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numString)
		}

		if evaluate(target, numbers, 0, 0) {
			output += target
		}
	}

	fmt.Println("Output Day 7 Part 2", output)
}

func readInputAsString() string {
	if _, err := os.Stat("Day7/input.txt"); err == nil {
		content, _ := os.ReadFile("Day7/input.txt")
		return string(content)
	}
	return ""
}

func readInput() []Calculation {
	candidateCalculations := []Calculation{}
	file, err := os.Open("Day7/test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		first, second := strings.Split(line, ": ")[0], strings.Split(line, ": ")[1]
		targetValue := int64(sti.Sti(first))
		componentString := strings.Split(second, " ")
		componentInts := make([]int64, len(componentString))

		for i, s := range componentString {
			componentInts[i] = int64(sti.Sti(s))
		}
		candidateCalculations = append(candidateCalculations, Calculation{targetValue, componentInts})

	}
	return candidateCalculations
}
