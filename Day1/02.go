package main

import (
	"embed"
	_ "embed"
	"fmt"
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
	//h.Expect(1, 31)
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
	//sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })

	//sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })

	itemCountmap := map[int64]int64{}

	for _, nums := range left {
		itemCountmap[nums] = 0
	}
	for _, num := range right {
		if _, ok := itemCountmap[num]; ok {
			itemCountmap[num]++
		}
	}
	for key, v := range itemCountmap {

		total = total + (key * v)

	}

	fmt.Println(total)

	return total
}
