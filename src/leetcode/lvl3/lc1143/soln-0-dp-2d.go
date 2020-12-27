package lc1143

func longestCommonSubsequence(text1 string, text2 string) int {
	var len1, len2 = len(text1), len(text2)

	var dp = make([][]int, 1+len1)
	for idx := range dp {
		dp[idx] = make([]int, 1+len2)
	}

	for idx1 := 0; idx1 < len1; idx1++ {
		for idx2 := 0; idx2 < len2; idx2++ {
			if text1[idx1] == text2[idx2] {
				dp[idx1+1][idx2+1] = 1 + dp[idx1][idx2]
			} else {
				if dp[idx1+1][idx2] > dp[idx1][idx2+1] {
					dp[idx1+1][idx2+1] = dp[idx1+1][idx2]
				} else {
					dp[idx1+1][idx2+1] = dp[idx1][idx2+1]
				}
			}
		}
	}

	return dp[len1][len2]
}
