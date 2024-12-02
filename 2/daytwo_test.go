package main

import (
	"testing"
)

func TestIsLevelUnsafe(t *testing.T) {
	tests := []struct {
		name      string
		direction string
		cur       int
		prev      int
		want      bool
	}{
		{"increasingWithGoodDirectionAndSafeDifference", "increasing", 4, 1, false},
		{"increasingWithGoodDirectionAndUnsafeDifference", "increasing", 5, 1, true},
		{"increasingWithBadDirectionAndSafeDifference", "increasing", 1, 4, true},
		{"increasingWithBadDirectionAndUnsafeDifference", "increasing", 1, 5, true},
		{"decreasingWithGoodDirectionAndSafeDifference", "decreasing", 1, 4, false},
		{"decreasingWithBadDirectionAndSafeDifference", "decreasing", 4, 1, true},
		{"decreasingWithBadDirectionAndUnsafeDifference", "decreasing", 5, 1, true},
		{"decreasingWithGoodDirectionAndUnsafeDifference", "decreasing", 1, 5, true},
		{"decreasingWithNoDirectionAndSafeDifference", "", 2, 5, true},
		{"decreasingWithNoDirectionAndUnsafeDifference", "", 1, 5, true},
		{"increasingWithNoDirectionAndSafeDifference", "", 5, 2, true},
		{"increasingWithNoDirectionAndUnsafeDifference", "", 5, 1, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isLevelUnsafe(test.direction, test.cur, test.prev)
			if got != test.want {
				t.Fatalf("isLevelUnsafe(%s, %d, %d), want '%t' got %t", test.direction, test.cur, test.prev, test.want, got)
			}
		})
	}
}
