// https://leetcode.com/problems/meeting-scheduler/
//
// Time Complexity:		O(`lens1` * lg(`lens1`)) + O(`lens2` * lg(`lens2`)) ~ O(lens & lg(lens))
// Space Complexity:	O(1)
//
// Reference:
//	https://leetcode.com/problems/meeting-scheduler/discuss/408491/JavaC%2B%2B-Two-Pointer-Clean-code-O(NlogN)-for-Sorting-Beat-100
// 	https://leetcode.com/problems/meeting-scheduler/discuss/865697/Simple-Golang-solution-Beats-100
package lc1229

import "sort"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	lens1, lens2 := len(slots1), len(slots2)

	sort.Slice(slots1, func(a, b int) bool { return slots1[a][0] < slots1[b][0] })
	sort.Slice(slots2, func(a, b int) bool { return slots2[a][0] < slots2[b][0] })

	idx1, idx2 := 0, 0

	for idx1 < lens1 && idx2 < lens2 {
		startIntersect := maxOf(slots1[idx1][0], slots2[idx2][0])
		endIntersect := minOf(slots1[idx1][1], slots2[idx2][1])

		if startIntersect+duration <= endIntersect {
			return []int{startIntersect, startIntersect + duration}
		} else if slots1[idx1][1] < slots2[idx2][1] {
			idx1++
		} else {
			idx2++
		}
	}

	return []int{}
}

func minOf(num int, nums ...int) int {
	min := num

	for _, value := range nums {
		if min > value {
			min = value
		}
	}

	return min
}

func maxOf(num int, nums ...int) int {
	max := num

	for _, value := range nums {
		if max < value {
			max = value
		}
	}

	return max
}
