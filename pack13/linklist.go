package pack13

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}

	cache := 0

	var l1V int
	var l2V int
	l3 := result

	for l1 != nil || l2 != nil || cache > 0 {

		if l1 != nil {
			l1V = l1.Val
		} else {
			l1V = 0
		}

		if l2 != nil {
			l2V = l2.Val
		} else {
			l2V = 0
		}

		l3V := l1V + l2V + cache
		cache = 0

		if l3V > 9 {
			cache = 1
			l3V = l3V - 10
		}

		fmt.Println("l1V : " , l1V)
		fmt.Println("l2V : " , l2V)
		fmt.Println("l3V : " , l3V)
		fmt.Println("===================================")

		l3.Next = &ListNode{Val: l3V}

		l3 = l3.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}

	return result.Next
}

func ListMain()  {
	list1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	resultList := addTwoNumbers(list1, list2)
	fmt.Println(resultList.Val, resultList.Next.Val, resultList.Next.Next.Val)
}
