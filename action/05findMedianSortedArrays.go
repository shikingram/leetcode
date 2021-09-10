package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//已知两个正序数组
	sorted := make([]int, 0, len(nums1)+len(nums2))
	sorted = append(sorted, nums1...)
	sorted = append(sorted, nums2...)

	sort.Ints(sorted)

	var avg = func(x, y int) float64 {
		fmt.Println(x, y)
		return float64(x+y) / 2
	}

	// 偶数返回两数平均数
	if len(sorted)%2 == 0 {
		a := len(sorted) / 2
		return avg(sorted[a-1], sorted[a])
	} else {
		a := len(sorted) / 2
		return float64(sorted[a])
	}
}

func main() {
	f := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(f)
}
