package pack01


type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode  {
	/**
		初始化
	 */
	head := ListNode{
		0,
		nil,
	}
	p := l1
	q := l2
	curr := head
	currArrSum := 0
	for p != nil || q != nil {
		x := 0
		y := 0
		if p != nil{
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := currArrSum + x + y
		currArrSum = sum / 10
		curr.Next = &ListNode{
			sum % 10,
			nil,
		}
		curr = *curr.Next
		if p != nil{
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if currArrSum > 0 {
		curr.Next = &ListNode{
			currArrSum,
			nil,
		}
	}
	return head.Next
}



func TwoSum1(nums [] int, target int) []int  {
	var returns [] int
	for i := 0; i < len(nums); i++  {
		for j := 0; j < len(nums); j++  {
			if nums[i] + nums[j] == target && i != j {
				returns = append(returns, i)
				returns = append(returns, j)
				goto STOP
			}
		}
	}
	STOP:
	return returns
}



func findMap(m map[int]int, index int) bool  {
	if _, err := m[index]; !err{
		return false
	}
	return true
}


func findValueMap(m map[int]int, value int) bool  {
	var result bool = false
	for _, v := range m{
		if value == v {
			result = true
			goto STOP
		}
	}
	STOP:
	return result
}

func addMap(m map[int]int, k int, v int) {
	if !findMap(m, v) {
		m[k] = v
	}
}

func deleteMap(m map[int]int, k int)  {
	if findMap(m, k) {
		delete(m, k)
	}
}
