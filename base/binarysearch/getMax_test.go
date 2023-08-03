package binarysearch

import (
	"fmt"
	"math"
	"testing"
)

func TestGetMax(t *testing.T) {
	arr := []int{1, 3, 5, 67, 34, 55, 33}
	max := getMax(arr)
	fmt.Println(max)
}

func getMax(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	M := L + ((R - L) >> 1)

	leftMax := process(arr, L, M)
	rightMax := process(arr, M+1, R)
	return int(math.Max(float64(leftMax), float64(rightMax)))
}
