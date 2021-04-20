// https://leetcode.com/problems/all-paths-from-source-to-target/
//
// Time Complexity:		O()
// Space Complexity:	O()
//
// Reference
// 	https://leetcode.com/problems/all-paths-from-source-to-target/discuss/752384/Golang-Simple-BFS
package lc0797

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := [][]int{}
	ans := [][]int{}
	paths = append(paths, []int{0})

	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]
		// last element in the `path`
		cur := path[len(path)-1]
		// to check if the last element is the end
		if cur == len(graph)-1 {
			ans = append(ans, path)
			continue
		}

		for _, next := range graph[cur] {
			copied := make([]int, len(path)+1)
			copy(copied, path)
			copied[len(path)] = next
			paths = append(paths, copied)
		}
	}

	return ans
}
