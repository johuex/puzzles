package sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return [][]int{}
	}

	sort.Ints(nums) // sort values
	res := [][]int{}

	// loop on two base pointer, found another two in [j+1; n-1]
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate by values
		}
		// pruning because another values lead to greater sum
		if int64(nums[i])+int64(nums[i+1])+int64(nums[i+2])+int64(nums[i+3]) > int64(target) {
			break
		}
		// continue because nums[i] is too small for target sum
		if int64(nums[i])+int64(nums[n-3])+int64(nums[n-2])+int64(nums[n-1]) < int64(target) {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue // skip duplicate by values
			}
			// pruning because another values lead to greater sum
			if int64(nums[i])+int64(nums[j])+int64(nums[j+1])+int64(nums[j+2]) > int64(target) {
				break
			}
			// continue because nums[i] & nums[j] is too small for target sum
			if int64(nums[i])+int64(nums[j])+int64(nums[n-2])+int64(nums[n-1]) < int64(target) {
				continue
			}

			// а дальше двумя другими указателями по краям ищем таргет
			left, right := j+1, n-1
			for left < right {
				sum := int64(nums[i]) + int64(nums[j]) + int64(nums[left]) + int64(nums[right])
				if sum == int64(target) {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++ // skip duplicate
					}
					for left < right && nums[right] == nums[right-1] {
						right-- // skip duplicate
					}
					// просто сужаем окно если сумма равна
					left++
					right--
				} else if sum < int64(target) {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}
