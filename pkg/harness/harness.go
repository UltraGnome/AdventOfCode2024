package harness

import (
	"embed"
	"fmt"

	"github.com/UltraGnome/AdventOfCode2024/pkg/benchmark"
	"github.com/UltraGnome/AdventOfCode2024/pkg/test"
)

type Harness[T comparable] struct {
	solve func(string) T
	input string
	test  test.Tester[T]
}

func New[T comparable](solve func(string) T, input string, tests embed.FS) *Harness[T] {
	h := Harness[T]{
		solve: solve,
		input: input,
	}
	h.test = test.Register(tests, h.solve)
	return &h
}

func (h *Harness[T]) Expect(n int, value T) {
	h.test.Expect(n, value)
	fmt.Printf("Test %d passed\n", n)
}

func (h *Harness[T]) Solve() {
	fmt.Println(h.solve(h.input))
}

func (h *Harness[T]) Benchmark(cond benchmark.Condition) {
	benchmark.Run(func() {
		h.solve(h.input)
	}, cond)
}
