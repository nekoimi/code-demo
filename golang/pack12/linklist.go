package pack12

import (
	"errors"
	"fmt"
)

// 线性数据结构
// 第一个节点为头节点，并不真实保存数据，头节点基本代表了整个链表
// 单链表的操作 正确移动链表的指针
/**
	- Get
	- Insert
	- Delete
	- Traverse
 */


// 节点
type Element int

// link node
type LinkNode struct {
	Data Element
	Next *LinkNode
}

// Get New Link
func NewLink() *LinkNode {
	return &LinkNode{
		0,
		nil,
	}
}

/**
	获取链表中的第i个元素
	从链表头一直遍历到链表i的位置
	算法时间复杂度 O(n)
 */
func (head *LinkNode) Get(i int) (Element, error) {
	p := head.Next
	for j := 1; j < i ; j++  {
		if p == nil {
			return -1, errors.New("not found")
		}
	}
	return p.Data, nil
}

/**
	在链表i的位置插入节点e
	先找到i - 1的位置
    判断是不是到达表尾
    更改链接指向
 */
func (head *LinkNode) Insert(i int, e Element) bool  {
	p := head
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	// 在i-1处插入i
	// 断开旧的链接 指向新的链接
	s := &LinkNode{Data: e}
	s.Next = p.Next
	p.Next = s
	return true
}


/**
	删除i元素
 */
func (head *LinkNode) Delete(i int) bool  {
	p := head.Next
	j := 1
	// 移动到链表的i-1处
	for p != nil && j < i {
		p = p.Next
		j++
	}

	if nil == p || j > i {
		return false
	}

	p.Next = p.Next.Next
	return true
}

/**
	遍历链表
 */
func (head *LinkNode) Traverse()  {
	p := head.Next
	for nil != p {
		fmt.Println(p.Data)
		p = p.Next
	}
}

func LinkMain()  {
	linkList := NewLink()
	linkList.Insert(1,23)
	linkList.Insert(1,233)
	linkList.Insert(1,2333)
	linkList.Insert(1,23333)
	linkList.Insert(1,233333)
	linkList.Insert(1,2333333)
	linkList.Insert(2,23)
	linkList.Insert(2,233)
	linkList.Insert(2,2333)
	linkList.Insert(2,23333)
	linkList.Insert(2,233333)
	linkList.Insert(2,2333333)
	linkList.Traverse()

	linkList.Delete(7)

	e, err := linkList.Get(9)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("LinkList 9 is :", e)
}

