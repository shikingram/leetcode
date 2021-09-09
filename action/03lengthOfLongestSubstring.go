package main

import "fmt"

/*
	Given a string s, find the length of the longest substring without repeating characters.
*/

func lengthOfLongestSubstring(s string) int {
	// 取最大值
	var max = func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	// 利用map的key不可重复
	m := make(map[byte]int)

	n := len(s)

	// rk表示右指针，ants表示长度
	rk, ants := -1, 0

	//遍历s
	for i := range s {

		// 每次剔除最左边的字符
		if i != 0 {
			delete(m, s[i-1])
		}

		// 右移指针1位，直到出现重复的字符停止
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		// 保存最长的长度
		ants = max(ants, rk-i+1)

	}
	return ants
}

func main() {
	s := "abcabcbb"

	fmt.Println(lengthOfLongestSubstring(s))
}
