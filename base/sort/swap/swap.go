package swap

func Swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
