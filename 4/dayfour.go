package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(filename string) []string {
	dat, err := os.ReadFile(filename)
	check(err)
	return strings.Split(strings.Replace(strings.TrimSpace(string(dat)), "\r", "", -1), "\n")
}

func checkOutOfBounds(puzzle []string, row int, col int) bool {
	return row < 0 || col < 0 || row >= len(puzzle) || col >= len(puzzle[0])
}

func checkDirection(puzzle []string, row int, col int, direction []int, searchTerm string) int {
	for letter := 1; letter < len(searchTerm); letter++ {
		nextRow := row + direction[0]*letter
		nextCol := col + direction[1]*letter
		if checkOutOfBounds(puzzle, nextRow, nextCol) {
			return 0
		}
		next := puzzle[row+direction[0]*letter][col+direction[1]*letter]
		if next != searchTerm[letter] {
			return 0
		}
	}
	return 1
}

func partOne(searchTerm string) int {
	puzzle := readLines(os.Args[1])
	var directions = [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	total := 0
	for row := range puzzle {
		for col := range puzzle[row] {
			if puzzle[row][col] == searchTerm[0] {
				for i := 0; i < 8; i++ {
					direction := directions[i]
					total += checkDirection(puzzle, row, col, direction, searchTerm)
				}
			}
		}
	}
	return total
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func partTwo(searchTerm string) int {
	puzzle := readLines(os.Args[1])
	var validHits = []string{
		string(searchTerm[0] + searchTerm[2]),
		string(searchTerm[2] + searchTerm[0]),
	}
	total := 0
	for row := 1; row < len(puzzle)-1; row++ {
		for col := 1; col < len(puzzle[0])-1; col++ {
			if puzzle[row][col] == searchTerm[1] {
				if contains(validHits, string(puzzle[row+1][col-1]+puzzle[row-1][col+1])) &&
					contains(validHits, string(puzzle[row-1][col-1]+puzzle[row+1][col+1])) {
					total++
				}
			}
		}
	}
	return total
}

func main() {
	fmt.Println(partOne("XMAS"))
	fmt.Println(partTwo("MAS"))
}
