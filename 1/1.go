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

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := int(len(a) / 2)
	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])
	return merge(left, right)
}

func merge(l []int, r []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	res = append(res, l[i:]...)
	res = append(res, r[j:]...)
	return res
}

func getSplitListsFromLinesAsInt(a []string, delimiter string) ([]int, []int) {
	var leftList []int
	var rightList []int
	for i := range a {
		split := strings.Split(a[i], delimiter)
		leftValue, err := strconv.Atoi(strings.TrimSpace(split[0]))
		check(err)
		leftList = append(leftList, leftValue)
		rightValue, err := strconv.Atoi(strings.TrimSpace(split[1]))
		check(err)
		rightList = append(rightList, rightValue)
	}
	return leftList, rightList
}

func getSortedLists(lines []string) ([]int, []int) {
	leftList, rightList := getSplitListsFromLinesAsInt(lines, "   ")
	leftList = mergeSort(leftList)
	rightList = mergeSort(rightList)
	return leftList, rightList
}

func getNumUniqValueInArray(a []int, n int) int {
	var occurrences = 0
	for i := range a {
		if a[i] == n {
			occurrences++
		}
	}
	return occurrences
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	leftList, rightList := getSortedLists(readLines(os.Args[1]))

	// part 1
	difference := 0
	for i := range leftList {
		difference += abs(leftList[i] - rightList[i])
	}
	fmt.Println(difference)

	// part 2
	similarity := 0
	for i := range leftList {
		similarity += leftList[i] * getNumUniqValueInArray(rightList, leftList[i])
	}
	fmt.Println(similarity)
}
