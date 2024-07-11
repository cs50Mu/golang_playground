package bst

func binarySearch(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2

	if nums[mid] == target {
		return mid
	} else if target > nums[mid] {
		return binarySearch(nums, target, mid+1, end)
	} else {
		return binarySearch(nums, target, start, mid-1)
	}
}

func binarySearchIter(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
