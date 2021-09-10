package main

import "sort"

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针遍历
func mergePro(nums1 []int, m int, nums2 []int, n int) {
	// 声明一个变量，存储排好序的数组
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		// nums1遍历结束
		if p1 == m {
			// 把nums2的剩余元素接过来
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		// nums2遍历结束
		if p2 == n {
			// 把nums1的剩余元素接过来
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		// 保存较小的值
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}

	copy(nums1, sorted)

}
