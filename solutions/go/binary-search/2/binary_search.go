package binarysearch

func SearchInts(list []int, key int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + (high-low)/2
		switch {
		case list[mid] == key:
			return mid
		case key < list[mid]:
			high = mid - 1
		default:
			low = mid + 1
		}
	}

	return -1
}
