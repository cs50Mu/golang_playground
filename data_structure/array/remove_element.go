package array

// 利用快慢指针
// 快指针负责找不等于val的，找到后直接跟放在慢指针所在处即可，不用担心覆盖的问题
// 因为慢指针所在处一定是等于val的或者已经存在的不等于val的元素
// 之所以可以这么做是因为，本题是要删除相同的元素，而且只检查前 k 个不等的元素
func removeElementSlowFast(nums []int, val int) int {
	slowIdx := 0
	for fastIdx := 0; fastIdx < len(nums); fastIdx++ {
		if nums[fastIdx] != val {
			nums[slowIdx] = nums[fastIdx]
			slowIdx++
		}
	}

	return slowIdx
}

// 利用首尾指针
func removeElement(nums []int, val int) int {
	if len(nums) == 1 {
		if nums[0] != val {
			return 1
		} else {
			return 0
		}
	}

	i := 0
	j := len(nums) - 1

	var res int
	// 双指针：一个在首，一个在尾
	for i < j {
		// 注意避免越界
		for i <= j && nums[i] != val {
			i++
			res++
		}
		for i <= j && nums[j] == val {
			j--
		}
		if i >= j {
			break
		}
		// swap
		nums[i], nums[j] = nums[j], nums[i]
	}

	return res
}
