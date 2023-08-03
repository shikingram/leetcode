package sort

import (
	"fmt"
	"testing"

	"github.com/shikingram/leetcode/base/sort/swap"
)

func TestBubbleSort(t *testing.T) {
	var arr = []int{1, 5, 4, 5, 3, 7, 2}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	/*
		冒泡的原理是：
			相邻两个元素比较大小，交换位置，大的往后移动,每次比较完了后，最后一个元素就不用动了
			- 1.依次比较 0~n-1 相邻元素大小
			- 2.依次比较 0~n-2 相邻元素大小
			- 3.依次比较 0~n-3 相邻元素大小
	*/

	for e := len(arr) - 1; e > 0; e-- { // 表示的是0 - e 的轮次
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				swap.Swap(arr, i, i+1)
			}
		}
	}
}
