package sort

import (
	"fmt"
	"testing"

	"github.com/shikingram/leetcode/base/sort/swap"
)

func TestSelectSort(t *testing.T) {
	var arr = []int{4, 2, 6, 2, 3, 8, 5}
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(arr []int) {
	/*
		选择排序的原理：
			拿每个元素和所有元素进行对比，如果比他大，就交换
	*/

	for i := 0; i < len(arr)-1; i++ { // 从第0个元素开始，到第n-2个元素结束
		var maxIndex = 0
		// 从i+1 下标往后找出最大元素下标
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				maxIndex = j
			} else {
				maxIndex = i
			}
			swap.Swap(arr, i, maxIndex)
		}
	}
}
