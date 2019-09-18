package pack23

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

// New List
func newList() *ListNode  {
	return &ListNode{}
}

// Insert List
func (head *ListNode) insertOfIndex (index int, value int) {
	if head == nil {
		return
	}
	p := head
	j := 0
	for p != nil && j < index{
		p = p.Next
		j++
	}
	if p == nil || j < index {
		fmt.Println(index, "超出链表的长度, 无法插入")
		return
	}
	n := &ListNode{Val:value}
	n.Next = p.Next
	p.Next = n
}


// 遍历链表
func (head *ListNode) traverse () {
	p := head
	for p != nil {
		fmt.Print(&p)
		fmt.Println("   ", p.Val)
		p = p.Next
	}
}

// 反转链表 - 就地反转
func (head *ListNode) reverse () {

}

func Pack23Main () {
	l1 := newList()
	l1.Val = 0
	for i := 20;i > 0 ;i--  {
		insertVal := i * 10
		l1.insertOfIndex(0, insertVal)
	}
	l1.traverse()

	l1.reverse()

	l1.traverse()
}
