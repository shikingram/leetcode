package sort

import (
	"fmt"
	"testing"

	"github.com/shikingram/leetcode/base/sort/swap"
)

func TestInsertSort(t *testing.T) {
	var arr = []int{1, 5, 4, 3, 7, 2}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(arr []int) {
	/*
		插入排序原理：
			依次把0~i 做有序了
			- 0~1有序，只需要arr[0] < arr[1]
			- 0~2 有序，需要arr[1] < arr[2]
	*/
	for i := 0; i < len(arr); i++ { // 0~i 做到有序
		for j := i - 1; j > 0 && arr[j+1] < arr[j]; j-- {
			swap.Swap(arr, j+1, j)
		}
	}
}
