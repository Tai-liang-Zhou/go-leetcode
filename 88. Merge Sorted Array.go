package main

import "fmt"

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	tempslice := make([]int, m+n)
// 	i, j := 0, 0

// 	for i < m && j < n {
// 		// fmt.Println("int len :", i, j)
// 		if nums1[i] <= nums2[j] {
// 			tempslice[i+j] = nums1[i]
// 			i++
// 		} else {
// 			tempslice[i+j] = nums2[j]
// 			j++
// 		}
// 	}

// 	for i < m {
// 		// fmt.Println("int len :", i, j)
// 		tempslice[i+j] = nums1[i]
// 		i++
// 	}

// 	for j < n {
// 		// fmt.Println("int len :", i, j)
// 		tempslice[i+j] = nums2[j]
// 		j++
// 	}
// 	for i := range nums1 {
// 		nums1[i] = tempslice[i]
// 	}

// 	fmt.Println(nums1)

// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m == 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}

	fmt.Println(nums1)
}
// func main() {
// 	num1 := []int{0, 0, 3, 0, 0, 0, 0, 0, 0}
// 	var m int = 3
// 	num2 := []int{-1, 1, 1, 1, 2, 3}
// 	var n int = 6

// 	merge(num1, m, num2, n)

// }
