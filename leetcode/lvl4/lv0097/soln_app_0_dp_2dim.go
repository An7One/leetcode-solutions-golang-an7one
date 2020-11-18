package lv0097

func isInterleave(s1 string, s2 string, s3 string) bool {
	var len1, len2, len3 = len(s1), len(s2), len(s3)

	if len1+len2 != len3 {
		return false
	}

	dp := make([][]bool, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]bool, len2+1)
	}

	dp[0][0] = true
	for idx := 1; idx <= len1; idx++ {
		dp[idx][0] = dp[idx-1][0] && s1[idx-1] == s3[idx-1]
	}

	for idx := 1; idx <= len2; idx++ {
		dp[0][idx] = dp[0][idx-1] && s2[idx-1] == s3[idx-1]
	}

	for idx1 := 1; idx1 <= len1; idx1++ {
		for idx2 := 1; idx2 <= len2; idx2++ {
			var idx3 = idx1 + idx2

			dp[idx1][idx2] = (dp[idx1-1][idx2] && s1[idx1-1] == s3[idx3-1]) || (dp[idx1][idx2-1] && s2[idx2-1] == s3[idx3-1])
		}
	}

	return dp[len1][len2]
}
