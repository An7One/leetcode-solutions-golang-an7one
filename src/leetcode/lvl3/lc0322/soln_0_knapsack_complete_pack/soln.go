// https://leetcode.com/problems/coin-change/
//
// Time Complexity:		O()
// Space Complexity:	O()
package lc0322

func coinChange(coins []int, amount int) int {
	var dp = make([]int, 1+amount)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1 + amount
	}

	for _, coin := range coins {
		for faceValue := coin; faceValue <= amount; faceValue++ {
			if dp[faceValue] > 1+dp[faceValue-coin] {
				dp[faceValue] = 1 + dp[faceValue-coin]
			}
		}
	}

	if dp[amount] >= 1+amount {
		return -1
	}

	return dp[amount]
}
