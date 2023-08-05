package sort

import (
	"fmt"
	"testing"

	"github.com/shikingram/leetcode/base/sort/swap"
)

func TestHeapSort(t *testing.T) {
	var arr = []int{1, 5, 4, 5, 3, 7, 2}
	heapSort(arr)
	fmt.Println(arr)
}

/*
	堆排序的原理：
		// 构建大根堆
		// 拿掉大根堆最大值，heapsize --
		// 进行heapify 调整堆结构为大根堆
		// 周而复始
}
*/
func heapSort(arr []int) {
	if len(arr) < 2 {
		return 
	}
	for i := 0 ;i < len(arr);i ++ {
		heapInsert(arr,i)
	}
	heapsize := len(arr)
	swap.Swap(arr,0,heapsize-1)

	for {
		if heapsize <= 0 {
			break
		}

		heapify(arr,0,heapsize)
		heapsize -- 
		swap.Swap(arr,0,heapsize)
	}
}


func heapInsert(arr []int,index int) {
	for{
		if arr[index] <= arr[(index-1)/2] {
			break
		}else{
			swap.Swap(arr,index ,(index -1) /2 )
			index = (index-1) /2
		}
	}
}

func heapify(arr []int ,index ,heapsize  int) {
	left := 2 * index +1 
	for {
		if left >= heapsize {
			break
		}

		var largest int
		if left +1 < heapsize && arr[left+1] > arr[left] {
			largest = left +1
		} else {
			largest  = left 
		}

		if arr[largest] <= arr[index] {
			largest = index 
		}

		if largest == index {
			break
		}
		swap.Swap(arr,largest,index)
		index = largest
		left = 2 * index + 1 
	}
}