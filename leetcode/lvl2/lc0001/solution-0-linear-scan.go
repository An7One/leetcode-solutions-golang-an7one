package lc0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		if key, ok := m[target-val]; ok {
			return []int{idx, key}
		}
		m[val] = idx
	}

	return nil
}
