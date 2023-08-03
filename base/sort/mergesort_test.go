package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 3, 2, 7, 8, 4, 1}
	mergeSort(arr)
	fmt.Println("result:", arr)
}

/*
	归并排序的原理：
		数列一分为2 两边各排好序  然后合并
		合并的时候  用两个指针分别移动两个数列  比较2个数列指针指向的元素
		较小的元素拷贝到新数列中，相等的情况 优先拷贝左边
		拷贝完一个数列，另一个数列剩下的元素全部拷贝到新数列尾部 就完成了整体排序
*/

func mergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) {
	if L == R {
		return
	}
	//求中间元素下标
	M := L + ((R - L) >> 1)
	process(arr, L, M)
	process(arr, M+1, R)
	merge(arr, L, M, R)
}

func merge(arr []int, L, M, R int) {
	fmt.Printf("L:%d,M:%d,R:%d \n", L, M, R)
	// 临时切片
	var temp []int
	// 左侧数组下标
	p1 := L
	// 右侧数组下标
	p2 := M + 1

	for {
		if p1 > M || p2 > R {
			break
		}

		if arr[p1] <= arr[p2] {
			temp = append(temp, arr[p1])
			p1++
		} else {
			temp = append(temp, arr[p2])
			p2++
		}
	}

	for ; p1 <= M; p1++ {
		temp = append(temp, arr[p1])
	}

	for ; p2 <= R; p2++ {
		temp = append(temp, arr[p2])
	}

	fmt.Println("temp", temp)

	// 	temp数列赋值给arr数列
	// Note:
	// 	注意arr是从L开始 + i
	for i := 0; i < len(temp); i++ {
		arr[L+i] = temp[i]
	}
}
