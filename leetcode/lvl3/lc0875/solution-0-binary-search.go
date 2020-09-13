package lc0875

func minEatingSpeed(piles []int, H int) int {
	const Range int = 1e9

	var lo, hi int = 1, Range

	for lo < hi {
		var totalHours int = 0

		var mid int = lo + (hi-lo)/2

		for _, pile := range piles {
			totalHours += (pile + mid - 1) / mid
		}

		if totalHours > H {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
