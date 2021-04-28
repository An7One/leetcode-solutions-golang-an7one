// https://leetcode.com/problems/power-of-three/
//
// Time Compleixty:		O(`n`)
// Space Complexity:	O(1)
//
// Reference:
//	https://leetcode.com/problems/power-of-three/discuss/77876/**-A-summary-of-all-solutions-(new-method-included-at-15:30pm-Jan-8th)/82291
package lc0326

func isPowerOfThree(n int) bool {
	if n > 1 {
		for n%3 == 0 {
			n /= 3
		}
	}

	return n == 1
}
