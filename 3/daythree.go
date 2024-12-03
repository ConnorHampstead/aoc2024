package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) string {
	dat, err := os.ReadFile(filename)
	check(err)
	return string(dat)
}

func filterByPattern(target string, pattern string) []string {
	r := regexp.MustCompile(pattern)
	return r.FindAllString(target, -1)
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	check(err)
	return i
}

func main() {
	reports := readFile(os.Args[1])
	partOne := filterByPattern(reports, `mul\(\d+,\d+\)`)
	partOneTotal := 0
	for i := range partOne {
		numbers := filterByPattern(partOne[i], `\d+`)
		partOneTotal += convertStringToInt(numbers[0]) * convertStringToInt(numbers[1])
	}
	fmt.Println(partOneTotal)

	partTwo := filterByPattern(reports, `(do|don't)\(\)|mul\(\d+,\d+\)`)
	partTwoTotal := 0
	count := true
	for i := range partTwo {
		curr := partTwo[i]
		if curr == "don't()" {
			count = false
			continue
		} else if curr == "do()" {
			count = true
			continue
		}
		if !count {
			continue
		}
		numbers := filterByPattern(partTwo[i], `\d+`)
		partTwoTotal += convertStringToInt(numbers[0]) * convertStringToInt(numbers[1])

	}
	fmt.Println(partTwoTotal)
}
