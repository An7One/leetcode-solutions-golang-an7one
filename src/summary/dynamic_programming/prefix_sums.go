package summary

// CreatePrefixSums to create prefixSums array based on the nums given
func CreatePrefixSums(nums []int) []int {
	var prefixSums = make([]int, 1+len(nums))
	for i, val := range nums {
		prefixSums[i+1] = prefixSums[i] + val
	}

	return prefixSums
}

// GetSubarraySum to query the sum of the subarray
func GetSubarraySum(lo int, hi int, prefixSums []int) int {
	return prefixSums[hi] - prefixSums[lo-1]
}
