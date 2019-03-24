package pack04


func BinarySearch(arr []int, target int) int  {
	mid := len(arr) >> 1
	if target == arr[mid] {
		return mid
	}
	start_index := 0
	end_index := len(arr) - 1
	for start_index <= end_index {
		mid = (start_index + end_index) >> 1
		if target < arr[mid] {
			end_index = mid - 1
		} else if target > arr[mid] {
			start_index = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
