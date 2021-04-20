// https://leetcode.com/problems/unique-paths/
// Time Complexity:		O(`m` * `n`)
// Space Complexity:	O(`n`)
//
// References:
//	https://youtu.be/BxblkIz6TZc?t=518
package lc0062

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			dp[c] = dp[c] + dp[c-1]
		}
	}

	return dp[n-1]
}
