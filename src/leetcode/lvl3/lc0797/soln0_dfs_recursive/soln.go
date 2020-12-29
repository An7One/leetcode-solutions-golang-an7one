package lc0797_soln0BFS

// https://leetcode.com/problems/all-paths-from-source-to-target/
// Time Complexity:		O()
// Space Complexity:	O()
// Reference
// 	https://leetcode.com/problems/all-paths-from-source-to-target/discuss/615535/Golang-DFS
func allPathsSourceTarget(graph [][]int) [][]int {
	paths := [][]int{}

	dfs(0, []int{0}, len(graph)-1, graph, &paths)

	return paths
}

func dfs(cur int,
	path []int,
	destination int,
	graph [][]int,
	paths *[][]int) {

	if cur == destination {
		copied := make([]int, len(path))
		copy(copied, path)
		*paths = append(*paths, copied)
		return
	}

	for _, next := range graph[cur] {
		path = append(path, next)
		dfs(next, path, destination, graph, paths)
		path = path[:len(path)-1]
	}
}
