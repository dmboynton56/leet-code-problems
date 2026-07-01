package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= merged[len(merged)-1][1] {
			intervals[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

func main() {
	fmt.Println("test 1:", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
