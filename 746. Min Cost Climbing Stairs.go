package main

import (
	"math"
)

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += int(math.Min(float64(cost[i-1]), float64(cost[i-2])))
	}

	return int(math.Min(float64(cost[len(cost)-1]), float64(cost[len(cost)-2])))
}

// func main() {
// 	test1 := []int{10, 15, 20}
// 	results := minCostClimbingStairs(test1)
// 	fmt.Println(results)
// }
