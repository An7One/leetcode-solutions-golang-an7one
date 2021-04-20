// https://leetcode.com/problems/unique-paths/
//
// Time Complexity:		O(len)
// Space Complexity:	O(1)
package lc1816

func truncateSentence(s string, k int) string {
	cnt := 0
	ans := ""

	for _, ch := range s {
		if ch == ' ' {
			cnt++
			if cnt == k {
				break
			}
		}

		// inefficient
		ans += string(ch)
	}

	return ans
}
