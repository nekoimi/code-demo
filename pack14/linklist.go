package pack14

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

/**
遍历链表
*/
func (head *ListNode) Traverse()  {
	p := head
	for nil != p {
		fmt.Println(p.Val)
		p = p.Next
	}
}

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

	给定一个链表: 1->2->3->4->5, 和 n = 2.
	当删除了倒数第二个节点后，链表变为 1->2->3->5.

	时间复杂度 O(n)
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	// 计算链表的长度
	p := head
	listLength := 0
	for p != nil {
		listLength ++
		p = p.Next
	}

	if n > listLength {
		return head
	}

	tmp := head
	// 倒数第n个节点 ===> 正数第 (length - n) + 1 个节点
	// 移动到 (length -n)个节点 删除Next
	// 如果 (length - n) == 0  删除链表的最后一个元素
	// 如果删除的是头节点
	if listLength == n {
		if n == 1 {
			head = nil
		} else { // 删除头节点
			head = head.Next
		}
	} else {
		for j := 1; j < (listLength - n); j++ {
			tmp = tmp.Next
		}

		tmp.Next = tmp.Next.Next
	}

	return head
}

// Main
func Pack14Man()  {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	res := removeNthFromEnd(list, 2)
	res.Traverse()
}
