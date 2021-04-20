// https://leetcode.com/problems/maximum-product-subarray/
//
// Time Complexity:		O()
// Space Complexity:	O()
package lc0152

import "math"

func maxProduct(nums []int) int {
	var prevMin, prevMax int = 1, 1
	var largest = math.MinInt8

	for _, num := range nums {
		curMin := minOf(num, prevMin*num, prevMax*num)
		var curMax = maxOf(num, prevMin*num, prevMax*num)

		if curMax > largest {
			largest = curMax
		}

		prevMin = curMin
		prevMax = curMax
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
