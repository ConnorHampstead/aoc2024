package main

import (
	"fmt"
	"os"
	"strconv"
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
	return strings.Split(string(dat), "\n")
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	check(err)
	return i
}

func convertStringArrayToInt(stringArray []string) []int {
	var integerArray []int

	for i := range stringArray {
		integerArray = append(integerArray, convertStringToInt(stringArray[i]))
	}

	return integerArray
}

func getDirection(cur int, prev int) string {
	if cur > prev {
		return "increasing"
	} else if cur < prev {
		return "decreasing"
	}
	return ""
}

func isDirectionContinuous(direction string, cur int, prev int) bool {
	if direction == "increasing" {
		return cur > prev
	} else if direction == "decreasing" {
		return cur < prev
	}
	return false
}

func isDifferenceWithinSafeRange(cur int, prev int) bool {
	difference := abs(cur - prev)
	return difference >= 1 && difference <= 3
}

func isLevelUnsafe(direction string, cur int, prev int) bool {
	return !(isDirectionContinuous(direction, cur, prev) && isDifferenceWithinSafeRange(cur, prev))
}

func isReportSafe(report []int) bool {
	var direction string
	for i := 1; i < len(report); i++ {
		current_level := report[i]
		previous_level := report[i-1]
		if i == 1 {
			direction = getDirection(current_level, previous_level)
		}
		if isLevelUnsafe(direction, current_level, previous_level) {
			return false
		}
	}
	return true
}

func isReportSafeWithRemovedLevel(report []int) bool {
	for i := range report {
		var tempReport []int
		for level := range report {
			if level != i {
				tempReport = append(tempReport, report[level])
			}
		}
		if isReportSafe(tempReport) {
			return true
		}
	}
	return false
}

func main() {
	reports := readLines(os.Args[1])
	var numberSafeReportsPartOne int = 0
	var numberSafeReportsPartTwo int = 0

	for i := range reports {
		report := convertStringArrayToInt(strings.Split(reports[i], " "))
		if isReportSafe(report) {
			numberSafeReportsPartOne++
			numberSafeReportsPartTwo++
		} else {
			if isReportSafeWithRemovedLevel(report) {
				numberSafeReportsPartTwo++
			}
		}
	}

	fmt.Println(numberSafeReportsPartOne)
	fmt.Println(numberSafeReportsPartTwo)
}
