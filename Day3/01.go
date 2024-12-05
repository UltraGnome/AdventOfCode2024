package main

//import (
//	_ "embed"
//)
////
////go:embed input.txt
//var input string
//
////go:embed test*.txt
//var tests embed.FS
//
//func main() {
//	h := harness.New(solve, input, tests, harness.Nothing())
//	h.Tester.Expect(1, 161)
//	h.Run()
//
//}
//
//func solve(input string) int {
//	total := 0
//
//	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
//	matches := r.FindAllStringSubmatch(input, -1)
//
//	for _, match := range matches {
//		total += sti.Sti(match[1]) * sti.Sti(match[2])
//	}
//
//	return total
//}
