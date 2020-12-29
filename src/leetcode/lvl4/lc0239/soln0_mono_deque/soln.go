package lc0239

const valueRange int = 1e4 + 7

func maxSlidingWindow(nums []int, k int) []int {
	lenSlidingWindow := len(nums) - k + 1

	var ans = make([]int, lenSlidingWindow)
	var deque = make([]int, lenSlidingWindow+1)
	for i := range deque {
		deque[i] = -valueRange
	}

	idxHead, idxTail := 0, -1

	for idx, val := range nums {
		if idxHead <= idxTail && idx-k+1 > deque[idxHead] {
			idxHead++
		}

		for idxHead <= idxTail && nums[deque[idxTail]] <= val {
			idxTail--
		}

		idxTail++
		deque[idxTail] = idx

		if idx >= k-1 {
			ans[idx-k+1] = nums[deque[idxHead]]
		}
	}

	return ans
}
