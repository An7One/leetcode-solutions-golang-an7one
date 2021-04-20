// @author: Leon
// https://leetcode.com/problems/fibonacci-number/
//
// Time Complexity:     O(`n`)
// Space Complexity:    O(1)
package lc0509

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	prev2 := 0
	prev1 := 1

	sum := 0

	for i := 2; i <= n; i++ {
		sum = prev2 + prev1
		prev2 = prev1
		prev1 = sum
	}

	return sum
}
