package foursum

import "sort"

func twoSum(nums []int, start int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int

	left := start
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum < target || (left > start && nums[left] == nums[left-1]) {
			left++
		} else if sum > target || (right < len(nums)-1 && nums[right] == nums[right+1]) {
			right--
		} else {
			result = append(result, []int{nums[left], nums[right]})
			left++
			right--
		}
	}

	return result
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	return kSum(nums, target, 0, 4)
}

func kSum(nums []int, target int, start int, k int) [][]int {
	var result [][]int

	if start == len(nums) {
		return result
	}

	if k == 2 {
		return twoSum(nums, start, target)
	}

	for i := start; i < len(nums); i++ {
		if i == start || nums[i-1] != nums[i] {
			for _, subset := range kSum(nums, target-nums[i], i+1, k-1) {
				result = append(result, append([]int{nums[i]}, subset...))
			}
		}
	}

	return result
}
