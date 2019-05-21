func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	count := 1
	pre := 0
	for cur := 1; cur < len(nums); cur += 1 {
		if nums[pre] == nums[cur] && count == 0 {
			continue
		}

		if nums[pre] == nums[cur] {
			count -= 1
		} else {
			count = 1
		}
		pre += 1
		if pre != cur {
			nums[pre] = nums[cur]
		}
	}
	return pre + 1
}
