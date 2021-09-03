package pack03

import "fmt"

func ArrayFind(arr [][] int, target int) bool  {
	if arr == nil || len(arr) == 0 || len(arr[0]) == 0 {
		return false
	}

	rows := len(arr)
	cols := len(arr[0])

	row := 0
	col := cols - 1

	for row <= rows -1 && col >= 0 {

		fmt.Printf("查找坐标轨迹: (%d, %d) Value : %d \n", row, col, arr[row][col])

		if target == arr[row][col] {
			return true
		} else if target > arr[row][col] {
			row ++
		} else {
			col --
		}
	}

	return false;
}
