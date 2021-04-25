// https://leetcode.com/problems/rotate-image/
//
// Time Complexity:		O(`nRows` * `nCols`)
// Space Complexity:	O(1)
package lc0048

func rotate(matrix [][]int) {
	// not used
	// nRows := len(matrix)
	reverseRows(matrix)

	for r := range matrix {
		nCols := len(matrix[r])
		for c := 1 + r; c < nCols; c++ {
			temp := matrix[r][c]
			matrix[r][c] = matrix[c][r]
			matrix[c][r] = temp
		}
	}
}

func reverseRows(matrix [][]int) {
	nRows := len(matrix)
	lo, hi := 0, nRows-1

	for lo < hi {
		temp := matrix[lo]
		matrix[lo] = matrix[hi]
		matrix[hi] = temp

		lo++
		hi--
	}
}
