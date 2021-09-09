package main

import "fmt"

/*
	给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

	请你将两个数相加，并以相同形式返回一个表示和的链表。

	你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	// 哨兵
	result := list
	temp := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}

		list.Next = &ListNode{
			Val:  temp % 10,
			Next: nil,
		}

		// 进位
		temp = temp / 10

		list = list.Next
	}
	return result.Next
}

func main() {
	var a = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	var b = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	c := addTwoNumbers(a, b)
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}
