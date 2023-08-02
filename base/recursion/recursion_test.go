package recursion

import (
	"fmt"
	"testing"
)

/*
	求阶乘 1 * 2 * 3 * 4 * ... * n
*/
func recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion(n-1)
}

func TestRecursion(t *testing.T) {
	fmt.Println(recursion(5))
}

/*
	尾部调用，每次调用时把值计算后传入，减少函数过程中的运算
*/
func recursionTail(n, a int) int {
	if n == 0 {
		return a
	}
	return recursionTail(n-1, a*n)
}

func TestRecursionTail(t *testing.T) {
	fmt.Println(recursionTail(5, 1))
}
