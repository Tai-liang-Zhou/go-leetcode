package main

func maxSubArray(nums []int) int {
	biggest := nums[0]
	cur := nums[0]
	for p := 1; p < len(nums); p++ {
		if cur >= 0 {
			cur += nums[p]
		} else {
			cur = nums[p]
		}
		biggest = max(biggest, cur)

	}
	return biggest
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// func main() {
// 	test1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
// 	maxSubArray(test1)
// 	// fmt.Println(test1)
// }
