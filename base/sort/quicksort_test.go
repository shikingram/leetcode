package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/shikingram/leetcode/base/sort/swap"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 5, 3, 3, 3, 5, 5, 2, 2, 7, 7, 8, 8, 8, 4, 4, 4, 1, 1, 1}
	quickSort(arr)
	fmt.Println("result:", arr)
}

/*
	快排原理：
		1.荷兰国旗问题
			给一个数列，和一个数num,找出小于num的数放到数列左边，大于num的数放到数列右边，等于num的数放中间
		2.快排原理
			首先随机下标取一个数，和数列最后一个位置的数字做交换，然后利用荷兰国旗问题，求出小于区间和大于区间
			再对两个区间做递归循环如此处理
*/

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSortProcess(arr, 0, len(arr)-1)
}

func quickSortProcess(arr []int, L, R int) {
	if L < R {
		// 随机取值放到R位置
		rand.Seed(time.Now().UnixNano())
		// 0 ~ len(arr) -1
		swap.Swap(arr, L+rand.Intn(R-L+1), R)
		arrP := partition(arr, L, R)
		fmt.Println("arrP", arrP)
		quickSortProcess(arr, L, arrP[0]-1)
		quickSortProcess(arr, arrP[1]+1, R)
	}
}

/*
荷兰国旗问题求解思路：

	三个指针分别记录小于区间下标(p1)，大于区间下标(p2)，以及当前遍历下标(i)
	1）当前下标的值target 小于 指定num 则，target和p1+1下标的值交换,p1 ++ ,i ++
	2）target == num ,  i ++
	3) target > num ,target 和 p2-1下标的值交换，p2 --,i 不变
*/
func partition(arr []int, L, R int) []int {
	fmt.Printf("L:%d,R:%d,arr[L]:%d,arr[R]:%d \n", L, R, arr[L], arr[R])
	i := L
	p1 := L - 1
	// p2 的位置不要从R +1 开始，因为R的位置不需要动
	p2 := R

	for {
		if i >= p2 {
			break
		}

		if arr[i] < arr[R] {
			p1++
			swap.Swap(arr, p1, i)
			i++
		} else if arr[i] > arr[R] {
			p2--
			swap.Swap(arr, p2, i)
		} else {
			i++
		}
	}

	swap.Swap(arr, p2, R)

	return []int{p1 + 1, p2}
}
