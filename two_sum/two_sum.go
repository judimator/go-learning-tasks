package two_sum

func twoSum(nums []int, target int) []int {
	callback := func(nums []int, target int) []int {
		result := make([]int, 2)
		for k, num := range nums {
			for i, el := range nums {
				if k == i {
					continue
				}

				if (num + el) == target {
					result[0] = k
					result[1] = i

					return result
				}
			}
		}
		return result
	}

	return callback(nums, target)
}
