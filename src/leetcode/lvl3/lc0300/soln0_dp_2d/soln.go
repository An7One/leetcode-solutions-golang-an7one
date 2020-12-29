package lc0300

func lengthOfLIS(nums []int) int {
	var longest int = 1

	dp := make([]int, len(nums))
	for idx := range dp {
		dp[idx] = 1
	}

	for hi := range dp {
		for lo := 0; lo < hi; lo++ {
			if nums[lo] < nums[hi] {
				if 1+dp[lo] > dp[hi] {
					dp[hi] = 1 + dp[lo]
				}

				if dp[hi] > longest {
					longest = dp[hi]
				}
			}
		}
	}

	return longest
}
