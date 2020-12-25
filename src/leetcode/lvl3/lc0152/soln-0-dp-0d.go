package lc0152

import "math"

func maxProduct(nums []int) int {
	var prev_min, prev_max int = 1, 1
	var largest = math.MinInt8

	for _, num := range nums {
		var cur_min = minOf(num, prev_min*num, prev_max*num)
		var cur_max = maxOf(num, prev_min*num, prev_max*num)

		if cur_max > largest {
			largest = cur_max
		}

		prev_min = cur_min
		prev_max = cur_max
	}

	return largest
}

func minOf(vars ...int) int {
	min := vars[0]

	for _, val := range vars {
		if val < min {
			min = val
		}
	}

	return min
}

func maxOf(vars ...int) int {
	max := vars[0]

	for _, val := range vars {
		if val > max {
			max = val
		}
	}

	return max
}
