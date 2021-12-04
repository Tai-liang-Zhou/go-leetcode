package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}
	if digits[0] == 0 {
		NewDigits := make([]int, len(digits)+1)
		NewDigits[0] = 1
		for i := range digits {
			NewDigits[i+1] = digits[i]
		}
		return NewDigits
	}
	return digits
}

// func main() {
// 	// test1 := []int{1, 2, 3}
// 	test2 := []int{9, 9, 9}
// 	fmt.Println(plusOne(test2))
// }
