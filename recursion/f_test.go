package recursion

import (
	"fmt"
	"testing"
)

/*
	求斐波那契数列中第n的值
	1 1 2 3 5 8 13 21 ... N-1 N 2N-1
*/
func f(n, a1, a2 int) int {
	if n == 1 {
		return a1
	}
	return f(n-1, a2, a1+a2)
}

func TestF(t *testing.T) {
	fmt.Println(f(5, 1, 1))
}
