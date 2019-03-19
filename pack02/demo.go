package pack02

func ListFind1(list []int, value int) int  {
	index := 0
	for i := 0; i < len(list) ; i++ {
		if list[i] == value {
			index = i
			break
		}
	}
	return index
}

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		pivot := list[0]
		var less []int
		for i := range list{
			if i <= pivot {
				less = append(less, i)
			}
		}
		var greater []int
		for j := range list{
			if j > pivot {
				greater = append(greater, j)
			}
		}
		tmp := append(quickSort(less), []int{pivot}...)
		return append(tmp, quickSort(greater)...)
	}
}

func ListFind2(list []int, value int) int  {

	list = quickSort(list)
	println(list)

	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == value {
			return mid
		}
		if guess > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
