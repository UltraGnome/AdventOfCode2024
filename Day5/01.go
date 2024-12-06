package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	file, err := os.Open("Day5/test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules := map[string]string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			rules[rule[0]] = rule[1]
		}
		if strings.TrimSpace(line) == "" {
			continue
		}

	}

}
