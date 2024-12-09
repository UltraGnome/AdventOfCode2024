package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	gStartY, gStartX := getStartPoint(lines)

	gPosY := gStartY
	gPosX := gStartX

	moves := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	num := 0
	part1 := func() {
		moveIdx := 0

		stepHistory := make(map[string]int)
		newKey := fmt.Sprintf("%d,%d", gPosY, gPosX)
		stepHistory[newKey] = 0

		for gPosY > 0 && gPosY < len(lines)-1 && gPosX > 0 && gPosX < len(lines[0])-1 {
			nextPosY := gPosY + moves[moveIdx][0] // Peek at next step
			nextPosX := gPosX + moves[moveIdx][1]

			if string(lines[nextPosY][nextPosX]) == "#" { // If next step is #
				moveIdx = (moveIdx + 1) % 4          // Turn
				nextPosY = gPosY + moves[moveIdx][0] // Peek
				nextPosX = gPosX + moves[moveIdx][1]
				num += 1
			}

			gPosY = nextPosY // Step forward
			gPosX = nextPosX
			newKey = fmt.Sprintf("%d,%d", gPosY, gPosX)
			_, ok := stepHistory[newKey]
			if !ok {
				stepHistory[newKey] = 0
			}
		}

		log.Println("Part1:", len(stepHistory), num)
	}
	part1()
}

func getStartPoint(lines []string) (int, int) {
	gStartY := 0
	gStartX := 0

	for y, line := range lines {
		for x, char := range line {
			if char == '^' {
				gStartY = y
				gStartX = x
			}
		}
	}
	return gStartY, gStartX
}
