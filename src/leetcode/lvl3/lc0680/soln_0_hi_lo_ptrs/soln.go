// https://leetcode.com/problems/valid-palindrome-ii/
//
// Time Complexity:		O(`len(s)`)
// Space Complexity:	O(1)
//
// Reference:
//	https://leetcode.com/problems/valid-palindrome-ii/discuss/107716/Java-O(n)-Time-O(1)-Space
package lc0680

func validPalindrome(s string) bool {
	lo, hi := 0, len(s)-1
	rns := []rune(s)

	for lo < hi {
		if rns[lo] != rns[hi] {
			return isPalindrome(lo+1, hi, rns) || isPalindrome(lo, hi-1, rns)
		}

		lo++
		hi--
	}

	return true
}

func isPalindrome(lo int, hi int, rns []rune) bool {
	for lo < hi {
		if rns[lo] != rns[hi] {
			return false
		}

		lo++
		hi--
	}

	return true
}
