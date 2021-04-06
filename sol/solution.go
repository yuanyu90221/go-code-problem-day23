package sol

func FindUnsortedArrayLen(nums []int) int {
	result := 0
	MAX := -1
	MAX_IDX := -1
	left := -1
	for idx, num := range nums {
		if idx == 0 {
			MAX = num
			continue
		}
		if MAX > num {
			if left == -1 {
				left = MAX_IDX
			}
			result = idx - left + 1
		} else {
			MAX = num
			MAX_IDX = idx
		}
	}
	return result
}
