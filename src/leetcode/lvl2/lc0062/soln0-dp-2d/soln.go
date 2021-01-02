package lc0062

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	nr, nc := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, nr)
	for r := range dp {
		dp[r] = make([]int, nc)
	}

	for r := range dp {
		if obstacleGrid[r][0] == 1 {
			break
		}

		dp[r][0] = 1
	}

	for c := range dp[0] {
		if obstacleGrid[0][c] == 1 {
			break
		}

		dp[0][c] = 1
	}

	for r := 1; r < nr; r++ {
		for c := 1; c < nc; c++ {
			if obstacleGrid[r][c] == 1 {
				continue
			}

			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}

	return dp[nr-1][nc-1]
}
