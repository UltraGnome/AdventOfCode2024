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

	part2 := func() {
		bad := 0 // Track the num of bad setups

		for y, line := range lines {
		nextObs:
			for x, _ := range line {
				// Reset the guard
				gPosY = gStartY
				gPosX = gStartX
				moveIdx := 0

				// Make a new copy of the original map
				newLines := make([]string, len(lines))
				copy(newLines, lines[:])

				// Stick the new obstruction somewhere...
				newLine := []byte(newLines[y])
				newLine[x] = byte(int('#'))
				newLines[y] = string(newLine)

				// New history
				stepHistory := make(map[string]int)

				// The key difference from part1 -- we store the direction in addition to position.
				// The direction and position should be enough to detect that we're on the same path as before.
				newKey := fmt.Sprintf("%d,%d,%d,%d", gPosY, gPosX, moves[moveIdx][0], moves[moveIdx][1])
				stepHistory[newKey] = 0

				for gPosY > 0 && gPosY < len(newLines)-1 && gPosX > 0 && gPosX < len(newLines[0])-1 {
					nextPosY := gPosY + moves[moveIdx][0] // Peek at next step
					nextPosX := gPosX + moves[moveIdx][1]

					for string(newLines[nextPosY][nextPosX]) == "#" { // If next step is #
						moveIdx = (moveIdx + 1) % 4          // Turn
						nextPosY = gPosY + moves[moveIdx][0] // Peek
						nextPosX = gPosX + moves[moveIdx][1]
					}

					gPosY = nextPosY // Step forward
					gPosX = nextPosX
					newKey = fmt.Sprintf("%d,%d,%d,%d", gPosY, gPosX, moves[moveIdx][0], moves[moveIdx][1])
					_, ok := stepHistory[newKey]
					if !ok {
						stepHistory[newKey] = 0
					} else {
						bad += 1
						// log.Println("FOUND:", bad)
						continue nextObs
					}
				}
			}
		}

		log.Println("Part2:", bad)

	}
	part2()
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
