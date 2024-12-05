package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	file, err := os.Open("Day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var full_text []string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		full_text = append(full_text, line)
		i += 1
	}

	count_horizontal := checkHorizontal(full_text) // check in straight lines
	pivoted_slice := pivotSlice(full_text)
	count_vertical := checkHorizontal(pivoted_slice) // turn vertical lines horizontal, re-use the function

	counter_diagonal := checkDiagonal(full_text)
	fmt.Println(count_horizontal, count_vertical, counter_diagonal)
	fmt.Println(count_horizontal + count_vertical + counter_diagonal)
	//day two
	xmas_count := getXmasFromDiagonals(full_text)
	fmt.Println(xmas_count)
}

func checkHorizontal(xmasList []string) int {
	var current_window string
	counter := 0
	row_length := len(xmasList[0])
	no_of_rows := len(xmasList)
	fmt.Println(row_length, no_of_rows)
	for i := 0; i < no_of_rows; i++ {
		for j := 0; j < row_length-3; j++ {
			current_window = xmasList[i][j : j+4]
			if current_window == "XMAS" || current_window == "SAMX" {
				counter += 1
			}
		}
	}
	return counter
}
func pivotSlice(xmasList []string) []string {
	var return_slice []string
	var new_string = ""
	row_length := len(xmasList[0])
	no_of_rows := len(xmasList)
	for i := 0; i < row_length; i++ {
		for j := 0; j < no_of_rows; j++ {
			new_string += string(xmasList[j][i])
		}
		return_slice = append(return_slice, new_string)
		new_string = ""
	}
	return return_slice
}

func checkDiagonal(xmasList []string) int {
	counter := 0
	row_length := len(xmasList[0])
	no_of_rows := len(xmasList)
	for i := 0; i < row_length-3; i++ {
		for j := 0; j < no_of_rows-3; j++ {
			diagonal_1 := string(xmasList[i][j]) + string(xmasList[i+1][j+1]) + string(xmasList[i+2][j+2]) + string(xmasList[i+3][j+3])
			diagonal_2 := string(xmasList[i][j+3]) + string(xmasList[i+1][j+2]) + string(xmasList[i+2][j+1]) + string(xmasList[i+3][j])

			if diagonal_1 == "XMAS" || diagonal_1 == "SAMX" {
				counter += 1
			}
			if diagonal_2 == "XMAS" || diagonal_2 == "SAMX" {
				counter += 1
			}
		}
	}
	return counter
}

func getXmasFromDiagonals(xmasList []string) int {
	counter := 0
	row_length := len(xmasList[0])
	no_of_rows := len(xmasList)
	for i := 0; i < row_length-2; i++ {
		for j := 0; j < no_of_rows-2; j++ {
			cross_1 := string(xmasList[i][j]) + string(xmasList[i+1][j+1]) + string(xmasList[i+2][j+2])
			cross_2 := string(xmasList[i][j+2]) + string(xmasList[i+1][j+1]) + string(xmasList[i+2][j])
			if cross_1 != "MAS" && cross_1 != "SAM" {
				continue
			}
			if cross_2 == "MAS" || cross_2 == "SAM" {
				counter += 1
			}
		}
	}
	return counter
}
