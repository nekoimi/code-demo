package pack15

type ListNode struct {
	Val int
	Next *ListNode
}

/**
	合并两个有序链表
	两个
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

//func merge(lists []*ListNode, left int, right int)  *ListNode {
//	if left == right {
//		return lists[left]
//	}
//
//	mid := left + (right - left)
//	l1 := merge(lists, left, mid)
//	l2 := merge(lists, mid + 1, right)
//	return mergeTwoLists(l1, l2)
//}
//
//
//func mergeKLists(lists []*ListNode) *ListNode {
//	if lists == nil {
//		return nil
//	}
//
//	length := len(lists)
//	return merge(lists, 0, length - 1)
//}

func Pack15Main()  {

}
