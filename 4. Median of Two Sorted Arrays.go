package main

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		v1 := findKth(nums1, 0, nums2, 0, l/2)
		v2 := findKth(nums1, 0, nums2, 0, (l/2)+1)
		return (float64(v1) + float64(v2)) / 2.0
	} else {
		return float64(findKth(nums1, 0, nums2, 0, (l+1)/2))
	}
}

func findKth(A []int, pa int, B []int, pb int, k int) int {
	m, n := len(A), len(B)
	if pa >= m {
		return B[pb+k-1]
	}
	if pb >= n {
		return A[pa+k-1]
	}

	if k == 1 {
		return int(math.Min(float64(A[pa]), float64(B[pb])))
	}

	va := math.MaxInt64
	if pa+k/2-1 < m {
		va = A[pa+k/2-1]
	}

	vb := math.MaxInt64
	if pb+k/2-1 < n {
		vb = B[pb+k/2-1]
	}

	if va <= vb {
		return findKth(A, pa+k/2, B, pb, k-k/2)
	} else {
		return findKth(A, pa, B, pb+k/2, k-k/2)
	}
}

// func main() {
// 	sliceA := []int{1}
// 	sliceB := []int{2, 3, 4, 5, 6}
// 	result_mid := findMedianSortedArrays(sliceA, sliceB)
// 	fmt.Println(result_mid)
// 	fmt.Println(3 / 2)
// }
