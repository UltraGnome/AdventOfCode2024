package algorithms

import (
	"github.com/UltraGnome/AdventOfCode2024/pkg/math"
)

func Shoelace[N math.Number](coords [][2]N) N {
	sum := N(0)
	last := coords[len(coords)-1]
	for _, coord := range coords {
		sum += coord[0]*last[1] - coord[1]*last[0]
		last = coord
	}
	return math.Abs(sum) / 2
}
