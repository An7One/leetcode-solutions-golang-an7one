// https://leetcode.com/problems/unique-paths/
// Time Complexity:		O(`m` * `n`)
// Space Complexity:	O(`m` * `n`)
//
// References:
//	https://youtu.be/BxblkIz6TZc?t=518
package lc0062

func uniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for r := range memo {
		memo[r] = make([]int, n)
	}

	return dfs(m-1, n-1, memo)
}

func dfs(r int, c int, memo [][]int) int {
	if r < 0 || c < 0 {
		return 0
	}

	if r == 0 && c == 0 {
		return 1
	}

	if memo[r][c] > 0 {
		return memo[r][c]
	}

	cnt := dfs(r-1, c, memo) + dfs(r, c-1, memo)
	memo[r][c] = cnt
	return cnt
}
