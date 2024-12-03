package main

import (
	"embed"
	_ "embed"
	"github.com/UltraGnome/AdventOfCode2024/pkg/math"
	"strconv"
	"strings"

	"github.com/UltraGnome/AdventOfCode2024/pkg/harness"
	"github.com/UltraGnome/AdventOfCode2024/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 2)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var total int = 0
	for _, line := range s {
		reportStrings := strings.Split(line, " ")
		reportInts := make([]int, len(reportStrings))
		for i, report := range reportStrings {
			reportInts[i], _ = strconv.Atoi(report)

		}
		if EvaluateReport(reportInts) {
			total++
		}

	}

	return total
}

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func EvaluateReport(report []int) bool {
	levelDirection := "" //down,up
	safe := true

	safe, levelDirection = initializeReport(report, safe, levelDirection)

	if !safe {
		return safe
	}

	for i := 0; i < len(report); i++ {
		if i == len(report)-1 {
			break
		}
		if report[i] == report[i+1] {
			safe = false
			return safe
		} else if levelDirection == "up" && report[i] > report[i+1] {
			safe = false
			return safe
		} else if levelDirection == "down" && report[i] < report[i+1] {
			safe = false
			return safe
		}
		if math.Abs(report[i]-report[i+1]) > 3 {
			safe = false
			return safe
		}
	}
	return safe

}

func initializeReport(report []int, safe bool, levelDirection string) (bool, string) {
	if report[0] == report[1] {
		safe = false
		return false, ""
	}
	if report[0] > report[1] {
		levelDirection = "down"
	} else if report[0] < report[1] {
		levelDirection = "up"
	}
	return safe, levelDirection
}
