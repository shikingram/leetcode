package binarysearch

import (
	"fmt"
	"testing"
)

/*
	binarySearch 二分查找递归解法
	l : 左边下标
	r : 右边下标
	target： 要找到的数
*/
func binarySearch(array []int, target int, l, r int) int {
	if l > r {
		// 出界了，找不到
		return -1
	}
	// 从中间开始找
	mid := (l + r) / 2
	middleNum := array[mid]
	if middleNum == target {
		return mid // 找到了
	} else if middleNum > target {
		// 中间的数比目标还大，从左边找
		return binarySearch(array, target, 1, mid-1)
	} else {
		// 中间的数比目标还小，从右边找
		return binarySearch(array, target, mid+1, r)
	}
}

func TestBinarySearch(t *testing.T) {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 89
	result := binarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)
}
