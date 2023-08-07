package sort

import (
	"container/heap"
	"testing"
)

func TestSorted(t *testing.T) {
	arr := []int{2, 1, 3, 5, 4, 7, 6, 8}

	sortedArrayDistanceLessK(arr, 1)

	t.Log(arr)
}

func TestHeap(t *testing.T) {
	var arrHeap IntHeap
	heap.Init(&arrHeap)
	heap.Push(&arrHeap, 2)
	heap.Push(&arrHeap, 4)
	heap.Push(&arrHeap, 1)
	res := heap.Pop(&arrHeap)
	t.Log(res)
}

/*
对几乎有序的数组排序

几乎有序的数组：
	如果把数组排好序的话，每个元素的移动距离不超过K，并且K相对数组较小
*/

func sortedArrayDistanceLessK(arr []int, k int) {
	var arrHeap IntHeap
	heap.Init(&arrHeap)
	// （1）先把0~k 位置上的数字构建小根堆

	var i int
	for ; i < k+1; i++ {
		heap.Push(&arrHeap, arr[i])
	}

	// (2) 设置index，表示排好序的数字下标
	var index int
	for ; index < len(arr) && i < len(arr); index++ {
		heap.Push(&arrHeap, arr[i])
		var temp = heap.Pop(&arrHeap)
		arr[index] = temp.(int)
		i++
	}

	// (3)剩下的数字放入数组中
	for {
		if arrHeap.Len() > 0 {
			var temp = heap.Pop(&arrHeap)
			arr[index] = temp.(int)
			index++
		} else {
			break
		}
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
