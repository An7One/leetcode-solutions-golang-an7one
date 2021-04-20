// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/
//
// Time Complexity:		O(`nRows` * (`nCols` ^ 2))
// Space Complexity:	O(`nRows` * `nCols`)
//
// Reference:
//  https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/discuss/303750/javacpython-find-the-subarray-with-target-sum/910389
package lc1074

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	// nRows := len(matrix)
	nCols := len(matrix[0])

	for row, _ := range matrix {
		for col := 1; col < nCols; col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}

	count := 0
	for lo := range matrix[0] {
		for hi := lo; hi < nCols; hi++ {
			sumToFreq := make(map[int]int)
			sumToFreq[0] = 1

			sum := 0
			for row := range matrix {
				sum += matrix[row][hi]
				if lo > 0 {
					sum -= matrix[row][lo-1]
				}

				expected := sum - target
				if freq, found := sumToFreq[expected]; found {
					count += freq
				}

				sumToFreq[sum]++
			}
		}
	}

	return count
}
