//给定一个由 n 个节点组成的网络，用 n x n 个邻接矩阵 graph 表示。在节点网络中，只有当 graph[i][j] = 1 时，节点 i 能够直接
//连接到另一个节点 j。 
//
// 一些节点 initial 最初被恶意软件感染。只要两个节点直接连接，且其中至少一个节点受到恶意软件的感染，那么两个节点都将被恶意软件感染。这种恶意软件的传
//播将继续，直到没有更多的节点可以被这种方式感染。 
//
// 假设 M(initial) 是在恶意软件停止传播之后，整个网络中感染恶意软件的最终节点数。 
//
// 我们可以从 initial 中删除一个节点，并完全移除该节点以及从该节点到任何其他节点的任何连接。 
//
// 请返回移除后能够使 M(initial) 最小化的节点。如果有多个节点满足条件，返回索引 最小的节点 。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：graph = [[1,1,0],[1,1,0],[0,0,1]], initial = [0,1]
//输出：0
// 
//
// 示例 2： 
//
// 
//输入：graph = [[1,1,0],[1,1,1],[0,1,1]], initial = [0,1]
//输出：1
// 
//
// 示例 3： 
//
// 
//输入：graph = [[1,1,0,0],[1,1,1,0],[0,1,1,1],[0,0,1,1]], initial = [0,1]
//输出：1
// 
//
// 
//
// 提示： 
// 
//
// 
// n == graph.length 
// n == graph[i].length 
// 2 <= n <= 300 
// graph[i][j] 是 0 或 1. 
// graph[i][j] == graph[j][i] 
// graph[i][i] == 1 
// 1 <= initial.length < n 
// 0 <= initial[i] <= n - 1 
// initial 中每个整数都不同 
// 
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 哈希表 👍 89 👎 0

package main

func main() {

}
//leetcode submit region begin(Prohibit modification and deletion)
func dfs(graph [][]int, initialSet, infectedSet []int, v int) {
	n := len(graph)
	for u := 0; u < n; u++ {
		if graph[v][u] == 0 || initialSet[u] == 1 || infectedSet[u] == 1 {
			continue
		}
		infectedSet[u] = 1
		dfs(graph, initialSet, infectedSet, u)
	}
}

func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)
	initialSet := make([]int, n)
	for _, v := range initial {
		initialSet[v] = 1
	}
	infectedBy := make([][]int, n)
	for _, v := range initial {
		infectedSet := make([]int, n)
		dfs(graph, initialSet, infectedSet, v)
		for u := 0; u < n; u++ {
			if infectedSet[u] == 1 {
				infectedBy[u] = append(infectedBy[u], v)
			}
		}
	}
	count := make([]int, n)
	for u := 0; u < n; u++ {
		if len(infectedBy[u]) == 1 {
			count[infectedBy[u][0]]++
		}
	}
	res := initial[0]
	for _, v := range initial {
		if count[v] > count[res] || count[v] == count[res] && v < res {
			res = v
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
