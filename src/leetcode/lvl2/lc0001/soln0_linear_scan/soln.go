package lc0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		expNum := target - num
		if expIdx, ok := m[expNum]; ok {
			return []int{idx, expIdx}
		}
		m[num] = idx
	}

	return nil
}
