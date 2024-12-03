package main

import (
	"embed"
	"github.com/UltraGnome/AdventOfCode2024/pkg/harness"
	"github.com/UltraGnome/AdventOfCode2024/pkg/math"
	"slices"
)

func solve(input [][]int) int {
	safeReports := 0

	for _, ints := range input {

		for i := range ints {
			if isSafe(slices.Concat(ints[:i], ints[i+1:])) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func isSafe(ints []int) bool {
	upOrDown := ints[1] > ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] == ints[i-1] {
			return false
		}
		if upOrDown != (ints[i] > ints[i-1]) {
			return false
		}
		if math.Abs(ints[i]-ints[i-1]) > 3 {
			return false
		}
	}
	return true
}

func main() {
	h := harness.New(solve, input, tests, harness.SplitNewlinesWithInts())
	h.Tester.Expect(1, 4)
	h.Run()
}

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS
