//给你一个有 n 个节点的 有向带权 图，节点编号为 0 到 n - 1 。图中的初始边用数组 edges 表示，其中 edges[i] = [fromi, 
//toi, edgeCosti] 表示从 fromi 到 toi 有一条代价为 edgeCosti 的边。 
//
// 请你实现一个 Graph 类： 
//
// 
// Graph(int n, int[][] edges) 初始化图有 n 个节点，并输入初始边。 
// addEdge(int[] edge) 向边集中添加一条边，其中 edge = [from, to, edgeCost] 。数据保证添加这条边之前对应的两
//个节点之间没有有向边。 
// int shortestPath(int node1, int node2) 返回从节点 node1 到 node2 的路径 最小 代价。如果路径不存在，
//返回 -1 。一条路径的代价是路径中所有边代价之和。 
// 
//
// 
//
// 示例 1： 
//
// 
//
// 输入：
//["Graph", "shortestPath", "shortestPath", "addEdge", "shortestPath"]
//[[4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]], [3, 2], [0, 3], [[1, 3, 4]
//], [0, 3]]
//输出：
//[null, 6, -1, null, 6]
//
//解释：
//Graph g = new Graph(4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]);
//g.shortestPath(3, 2); // 返回 6 。从 3 到 2 的最短路径如第一幅图所示：3 -> 0 -> 1 -> 2 ，总代价为 3 +
// 2 + 1 = 6 。
//g.shortestPath(0, 3); // 返回 -1 。没有从 0 到 3 的路径。
//g.addEdge([1, 3, 4]); // 添加一条节点 1 到节点 3 的边，得到第二幅图。
//g.shortestPath(0, 3); // 返回 6 。从 0 到 3 的最短路径为 0 -> 1 -> 3 ，总代价为 2 + 4 = 6 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 100 
// 0 <= edges.length <= n * (n - 1) 
// edges[i].length == edge.length == 3 
// 0 <= fromi, toi, from, to, node1, node2 <= n - 1 
// 1 <= edgeCosti, edgeCost <= 10⁶ 
// 图中任何时候都不会有重边和自环。 
// 调用 addEdge 至多 100 次。 
// 调用 shortestPath 至多 100 次。 
// 
//
// Related Topics 图 设计 最短路 堆（优先队列） 👍 43 👎 0

package main

import (
	"container/heap"
	"math"
)

func main() {

}
//leetcode submit region begin(Prohibit modification and deletion)
type Graph struct {
	graph [][]Pair
}

func Constructor(n int, edges [][]int) Graph {
	g := make([][]Pair, n)
	for i := 0; i < n; i++ {
		g[i] = []Pair{}
	}
	for _, e := range edges {
		x, y, cost := e[0], e[1], e[2]
		g[x] = append(g[x], Pair{y, cost})
	}
	return Graph{g}
}

func (this *Graph) AddEdge(edge []int)  {
	x, y, cost := edge[0], edge[1], edge[2]
	(*this).graph[x] = append((*this).graph[x], Pair{y, cost})
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	pq := PriorityQueue{}
	dist := make([]int, len((*this).graph))
	for i := 0; i < len((*this).graph); i++ {
		dist[i] = math.MaxInt32
	}
	dist[node1] = 0
	heap.Push(&pq, Pair{0, node1})
	for len(pq) > 0 {
		p := heap.Pop(&pq).(Pair)
		cost, cur := p.first, p.second
		if cur == node2 {
			return cost
		}
		for _, e := range (*this).graph[cur] {
			next, ncost := e.first, e.second
			if dist[next] > cost + ncost {
				dist[next] = cost + ncost
				heap.Push(&pq, Pair{cost + ncost, next})
			}
		}
	}
	return -1;
}

type Pair struct {
	first int
	second int
}

type PriorityQueue []Pair

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].first < pq[j].first
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Pair))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n - 1]
	*pq = (*pq)[:n - 1]
	return x
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
//leetcode submit region end(Prohibit modification and deletion)
