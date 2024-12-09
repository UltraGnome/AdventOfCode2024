package main

import (
	"bufio"
	"fmt"
	"github.com/UltraGnome/AdventOfCode2024/pkg/sti"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("Day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0
	var rules = make(map[int][]int)

	var intPages [][]int

	intPages = readInput(scanner, rules)
	for _, pgSet := range intPages {

		if !isCompliant(pgSet, rules) {
			sortUpdate(&pgSet, rules)
			total += pgSet[len(pgSet)/2]
			continue
		}
	}

	fmt.Println(total)
}

func isCompliant(pgSet []int, rules map[int][]int) bool {
	result := true

	for i, pg := range pgSet {
		for j := i + 1; j < len(pgSet); j++ {
			// If there exists a later value that the current value depends
			// on, this is invalid as it needs to be to the left.
			if slices.Contains(rules[pg], pgSet[j]) {
				result = false
			}
		}
	}
	return result
}

func sortUpdate(update *[]int, rules map[int][]int) {
	sort.Slice(*update, func(i, j int) bool {
		n1 := (*update)[i]
		n2 := (*update)[j]
		rules1, exists := (rules)[n1]
		if exists && slices.Contains(rules1, n2) {
			return true
		}
		rules2, exists := (rules)[n2]
		if exists && slices.Contains(rules2, n1) {
			return false
		}
		return true
	})
}

func readInput(scanner *bufio.Scanner, rules map[int][]int) [][]int {
	stringPages := []string{}
	intPages := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			if _, ok := rules[sti.Sti(rule[1])]; ok {
				rules[sti.Sti(rule[1])] = append(rules[sti.Sti(rule[1])], sti.Sti(rule[0]))
			} else {
				myPages := []int{sti.Sti(rule[0])}
				rules[sti.Sti(rule[1])] = myPages
			}

		}

		if strings.Contains(line, ",") {
			stringPages = append(stringPages, line)
		}
	}
	for _, strPage := range stringPages {
		sp := strings.Split(strPage, ",")
		myPg := []int{}
		for _, v := range sp {
			myPg = append(myPg, sti.Sti(v))
		}
		intPages = append(intPages, myPg)
	}
	return intPages
}
