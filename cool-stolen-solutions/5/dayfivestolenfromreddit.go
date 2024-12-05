package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func readLines(filename string, delimiter string) []string {
	dat, _ := os.ReadFile(filename)
	return strings.Split(strings.Replace(strings.TrimSpace(string(dat)), "\r", "", -1), delimiter)
}

func main() {
	split := readLines(os.Args[1], "\n\n")

	start := time.Now()
	cmp := func(a, b string) int {
		for _, s := range strings.Split(split[0], "\n") {
			if s := strings.Split(s, "|"); s[0] == a && s[1] == b {
				return -1
			}
		}
		return 0
	}

	run := func(sorted bool) (r int) {
		for _, s := range strings.Split(split[1], "\n") {
			if s := strings.Split(s, ","); slices.IsSortedFunc(s, cmp) == sorted {
				slices.SortFunc(s, cmp)
				n, _ := strconv.Atoi(s[len(s)/2])
				r += n
			}
		}
		return r
	}

	fmt.Println("Solved in", time.Since(start))
	fmt.Println(run(true))
	fmt.Println(run(false))
}
