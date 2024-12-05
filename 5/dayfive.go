package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(filename string, delimiter string) []string {
	dat, err := os.ReadFile(filename)
	check(err)
	return strings.Split(strings.Replace(strings.TrimSpace(string(dat)), "\r", "", -1), delimiter)
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	check(err)
	return i
}

func keysContain(searchMap map[string][]string, term string) bool {
	for k := range searchMap {
		if term == k {
			return true
		}
	}
	return false
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func getAllDependentNumbers(arr []string, num string) []string {
	var dependentNumbers []string
	for i := range arr {
		number := strings.Split(arr[i], "|")[0]
		if num == number {
			dependentNumbers = append(dependentNumbers, strings.Split(arr[i], "|")[1])
		}
	}
	return dependentNumbers
}

// getDependencies returns a map where each unique number in the ruleset
// is the key, and the value is a list of strings of each number which must
// come after it.
func getDependencies(numbers []string) map[string][]string {
	dependencies := make(map[string][]string)
	for _, v := range numbers {
		number := strings.Split(v, "|")[0]
		if !keysContain(dependencies, number) {
			dependencies[number] = getAllDependentNumbers(numbers, number)
		}
	}
	return dependencies
}

// reorderRow starts at the end of the row and works backwards to find
// any entry which is present in the list of dependencies, and swaps it
// with the current number if it is present, and recursively orders
// the new row until the row meets the dependency criteria.
func reorderRow(dependencies map[string][]string, row []string) []string {
	for j := range row {
		index := len(row) - j - 1
		currentNum := row[index]
		for k := index; k >= 0; k-- {
			if contains(dependencies[currentNum], row[k]) {
				row[index] = row[k]
				row[k] = currentNum
				return reorderRow(dependencies, row)
			}
		}
	}
	return row
}

func main() {
	sections := readLines(os.Args[1], "\n\n")
	pairs := sections[0]
	updates := strings.Split(sections[1], "\n")
	dependencies := getDependencies(strings.Split(pairs, "\n"))

	total := 0
	totalp2 := 0
	start := time.Now()
	for _, v := range updates {
		row := strings.Split(v, ",")
		rowCopy := make([]string, len(row))
		copy(rowCopy, row)
		orderedRow := reorderRow(dependencies, rowCopy)
		if slices.Equal(row, orderedRow) {
			total += convertStringToInt(row[len(row)/2])
		} else {
			totalp2 += convertStringToInt(orderedRow[len(row)/2])
		}
	}
	fmt.Println("Solved in", time.Since(start))
	fmt.Println(total)
	fmt.Println(totalp2)
}
