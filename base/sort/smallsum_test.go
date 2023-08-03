package sort

import (
	"testing"
)

func TestSmallSum(t *testing.T) {
	arr := []int{1, 3, 4, 2, 5}
	res := smallSum(arr)
	t.Log(res)
}

/*
	小和问题：
		什么是小和？
		例如： 1,3,4,2，5
		对每个元素找到左边比他小的值加一起
		1  无
		3  1
		4  3 + 1
		2  1
		5  1+3 +4 +2
		小和 = 16

		代码中可以逆向思维转化为如下问题：
			找到右边元素中：
			比1 大的元素的个数 * 1 
			比3 大的元素个数 * 3
			比4 大的元素个数 * 4 
			比2 大的元素个数 * 2 
			比5 大的元素个数 * 5
		代入分别为：
			4 * 1
			2 * 3 
			1 * 4 
			1 * 2 
			无
			小和 = 16
*/

func smallSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	return smallSumProcess(arr, 0, len(arr)-1)
}

func smallSumProcess(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	M := L + ((R - L) >> 1)

	return smallSumProcess(arr, L, M) + smallSumProcess(arr, M+1, R) + smallSumMerge(arr, L, M, R)

}

func smallSumMerge(arr []int, L, M, R int) int {
	var temp []int
	var p1 = L
	var p2 = M + 1
	var res int
	for {
		if p1 > M || p2 > R {
			break
		}

		// Note : 此处不能加=号，相等时必须先拷贝右边数组
		if arr[p1] < arr[p2] {
			temp = append(temp, arr[p1])
			res += (R - p2 + 1) * arr[p1]
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

	for i := 0; i < len(temp); i++ {
		arr[L+i] = temp[i]
	}

	return res
}
