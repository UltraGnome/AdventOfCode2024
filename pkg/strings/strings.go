package strings

import (
	"github.com/UltraGnome/AdventOfCode2024/pkg/buffers"
	"github.com/UltraGnome/AdventOfCode2024/pkg/chars"
)

func Ints(s string) []int {
	var ints []int
	buffer := buffers.Number(0)
	set := false
	neg := false
	for _, char := range s {
		if chars.IsNum(char) {
			buffer.Add(chars.NumVal(char))
			set = true
			continue
		}
		if char == '-' {
			neg = true
			continue
		}
		if set {
			b := buffer.Clear()
			if neg {
				b = -b
			}
			neg = false
			ints = append(ints, b)
			set = false
		}
	}
	if set {
		b := buffer.Clear()
		if neg {
			b = -b
		}
		ints = append(ints, b)
	}
	return ints
}
