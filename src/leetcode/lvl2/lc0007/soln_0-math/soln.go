package lc0007

import "math"

func reverse(x int) int {
	var ans int64 = 0

	for x != 0 {
		ans = ans*10 + int64(x%10)
		x = x / 10
	}

	if ans < math.MinInt32 || ans > math.MaxInt32 {
		return 0
	}

	return int(ans)
}
