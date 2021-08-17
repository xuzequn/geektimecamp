package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 指针递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1.Next = mergeTwoLists(l1.Next, l2)
			return l1
		} else {
			l2.Next = mergeTwoLists(l1, l2.Next)
			return l2
		}
	} else {
		if l1 == nil {
			return l2
		} else {
			return l1
		}
	}

}
