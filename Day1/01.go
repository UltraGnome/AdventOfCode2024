package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/UltraGnome/AdventOfCode2024/pkg/math"
	"sort"
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
	h.Expect(1, 11)
	h.Solve()
}

func solve(input string) int64 {
	s := utils.ParseInput(input)

	left := []int64{}
	right := []int64{}

	var total int64 = 0
	for _, line := range s {
		pair := strings.Split(line, "   ")
		n, err := strconv.Atoi(pair[0])
		if err == nil {
		}
		left = append(left, int64(n))

		m, e := strconv.Atoi(pair[1])
		if e == nil {
		}
		right = append(right, int64(m))
	}
	sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })

	sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })

	for i := 0; i < len(left); i++ {
		fmt.Println("abs diff:" + strconv.FormatInt(math.Abs(left[i]-right[i]), 10))
		total += math.Abs(left[i] - right[i])
	}
	fmt.Println("total is: " + strconv.FormatInt(total, 10))
	return total
}
