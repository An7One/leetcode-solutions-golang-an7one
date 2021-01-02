// https://leetcode.com/problems/unique-paths/
// Time Complexity:		O(`m` * `n`)
// Space Complexity:	O(`m` * `n`)
//
// References:
//	https://youtu.be/BxblkIz6TZc?t=518

package lc0062

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		dp[i][0] = 1
	}

	for i := range dp[0] {
		dp[0][i] = 1
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}

	return dp[m-1][n-1]
}
