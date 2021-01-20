// https://leetcode.com/problems/longest-palindromic-substring/
//
// Time Complexity:		O(`len(s)` ^ 2)
// Space Complexity:	O(`len(s)` ^ 2)
//
// References:
//	https://leetcode.com/problems/longest-palindromic-substring/discuss/2921/Share-my-Java-solution-using-dynamic-programming/3570
//	https://leetcode.com/problems/longest-palindromic-substring/discuss/2921/Share-my-Java-solution-using-dynamic-programming/223961
package lc0005

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	idxStart, idxEnd := 0, 0

	for lo := len(s) - 1; lo >= 0; lo-- {
		for hi := lo; hi < len(s); hi++ {
			dp[lo][hi] = s[lo] == s[hi] && (hi-lo < 2 || dp[lo+1][hi-1])

			if dp[lo][hi] && hi-lo > idxEnd-idxStart {
				idxStart, idxEnd = lo, hi
			}
		}
	}

	return s[idxStart : idxEnd+1]
}
