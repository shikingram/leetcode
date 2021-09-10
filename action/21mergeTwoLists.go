package main

/*
	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 哨兵，指向结果链表
	prehead := &ListNode{}
	result := prehead

	// 如果l1或者l2有一个为nil，则退出循环
	for l1 != nil && l2 != nil {
		// 结果链表指向更小的
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}

	// 指向剩余的链表
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}
