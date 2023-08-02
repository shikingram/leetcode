package tower

import (
	"fmt"
	"testing"
)

var total = 0

/*
	第一步，把n-1个盘子从A杆子上移动到B杆子上
	第二步，把第n个盘子移动到C杆子上
	第三步，把n-1个盘子从B杆子上移动到C杆子上
*/
func tower(n int, a, b, c string) {
	//第二步，把第n个盘子移动到C杆子上
	if n == 1 {
		total = total + 1
		fmt.Println(a, "->", c)
		return
	}
	//第一步，把n-1个盘子从A杆子上移动到B杆子上
	tower(n-1, a, c, b)
	total = total + 1
	fmt.Println(a, "->", c)
	//第三步，把n-1个盘子从B杆子上移动到C杆子上
	tower(n-1, b, a, c)
	return
}

func TestTower(t *testing.T) {
	n := 3   // 3 个盘子
	a := "a" // 杆子A
	b := "b" // 杆子B
	c := "c" // 杆子C
	tower(n, a, b, c)
	fmt.Println("total:", total)
}
