package main

import (
	"sort"
)

func isCovered(ranges [][]int, left int, right int) bool {
	// sort the ranges 由小到大
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// mark last position covered
	for _, r := range ranges {
		if left >= r[0] && left <= r[1] {
			left = r[1] + 1
		}
	}

	// check it all ranges cover
	return left >= right+1

}

// func main() {
// 	twoDimSlice := [][]int{
// 		{1, 2},
// 		{3, 4},
// 		{5, 6},
// 	}

// 	result := isCovered(twoDimSlice, 2, 5)
// 	fmt.Println(result)
// }
