package main

import (
	"embed"
	_ "embed"

	"github.com/UltraGnome/AdventOfCode2024/pkg/chars"
	"github.com/UltraGnome/AdventOfCode2024/pkg/harness"
	"github.com/UltraGnome/AdventOfCode2024/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 142)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	for _, line := range s {
		firstDigit := -1
		lastDigit := 0

		for _, char := range line {
			if !chars.IsNum(char) {
				continue
			}

			if firstDigit == -1 {
				firstDigit = chars.NumVal(char)
			}
			lastDigit = chars.NumVal(char)
		}
		total += (firstDigit * 10) + lastDigit
	}
	return total
}
