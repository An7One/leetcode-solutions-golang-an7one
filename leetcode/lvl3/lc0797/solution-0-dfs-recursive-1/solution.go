package lc0797

// https://leetcode.com/problems/all-paths-from-source-to-target/
// Time Complexity:		O()
// Space Complexity:	O()
// Reference
// 	https://leetcode.com/problems/all-paths-from-source-to-target/discuss/582435/golang-clean-recursive-backtracking
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	var ans [][]int
	// visited := make([]bool, n)
	// visited[n - 1] = true

	var rec func(path []int, cur int)
	rec = func(path []int, cur int) {
		// if !visited[cur]{
		for _, next := range graph[cur] {
			if next == n-1 {
				copied := make([]int, len(path)+1)
				copy(copied, path)
				copied[len(copied)-1] = n - 1
				ans = append(ans, copied)
			} else {
				// visited[cur] = true
				rec(append(path, next), next)
				// visited[cur] = false
			}
		}
		// }
	}

	rec([]int{0}, 0)
	return ans
}
