package main

func maxSlidingWindow(nums []int, k int) []int {
	const value_range int = 1e4 + 7

	var len_sliding_window int = len(nums) - k + 1

	var ans = make([]int, len_sliding_window)
	var deque = make([]int, len_sliding_window+1)
	for i := range deque {
		deque[i] = -value_range
	}

	var idx_head, idx_tail int = 0, -1

	for idx, val := range nums {
		if idx_head <= idx_tail && idx-k+1 > deque[idx_head] {
			idx_head++
		}

		for idx_head <= idx_tail && nums[deque[idx_tail]] <= val {
			idx_tail--
		}

		idx_tail++
		deque[idx_tail] = idx

		if idx >= k-1 {
			ans[idx-k+1] = nums[deque[idx_head]]
		}
	}

	return ans
}
