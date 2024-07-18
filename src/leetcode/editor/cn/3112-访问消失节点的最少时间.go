//给你一个二维数组 edges 表示一个 n 个点的无向图，其中 edges[i] = [ui, vi, lengthi] 表示节点 ui 和节点 vi 之间
//有一条需要 lengthi 单位时间通过的无向边。
//
// 同时给你一个数组 disappear ，其中 disappear[i] 表示节点 i 从图中消失的时间点，在那一刻及以后，你无法再访问这个节点。
//
// 注意，图有可能一开始是不连通的，两个节点之间也可能有多条边。
//
// 请你返回数组 answer ，answer[i] 表示从节点 0 到节点 i 需要的 最少 单位时间。如果从节点 0 出发 无法 到达节点 i ，那么
//answer[i] 为 -1 。
//
//
//
// 示例 1：
//
//
//
//
// 输入：n = 3, edges = [[0,1,2],[1,2,1],[0,2,4]], disappear = [1,1,5]
//
//
// 输出：[0,-1,4]
//
// 解释：
//
// 我们从节点 0 出发，目的是用最少的时间在其他节点消失之前到达它们。
//
//
// 对于节点 0 ，我们不需要任何时间，因为它就是我们的起点。
// 对于节点 1 ，我们需要至少 2 单位时间，通过 edges[0] 到达。但当我们到达的时候，它已经消失了，所以我们无法到达它。
// 对于节点 2 ，我们需要至少 4 单位时间，通过 edges[2] 到达。
//
//
// 示例 2：
//
//
//
//
// 输入：n = 3, edges = [[0,1,2],[1,2,1],[0,2,4]], disappear = [1,3,5]
//
//
// 输出：[0,2,3]
//
// 解释：
//
// 我们从节点 0 出发，目的是用最少的时间在其他节点消失之前到达它们。
//
//
// 对于节点 0 ，我们不需要任何时间，因为它就是我们的起点。
// 对于节点 1 ，我们需要至少 2 单位时间，通过 edges[0] 到达。
// 对于节点 2 ，我们需要至少 3 单位时间，通过 edges[0] 和 edges[1] 到达。
//
//
// 示例 3：
//
//
// 输入：n = 2, edges = [[0,1,1]], disappear = [1,1]
//
//
// 输出：[0,-1]
//
// 解释：
//
// 当我们到达节点 1 的时候，它恰好消失，所以我们无法到达节点 1 。
//
//
//
// 提示：
//
//
// 1 <= n <= 5 * 10⁴
// 0 <= edges.length <= 10⁵
// edges[i] == [ui, vi, lengthi]
// 0 <= ui, vi <= n - 1
// 1 <= lengthi <= 10⁵
// disappear.length == n
// 1 <= disappear[i] <= 10⁵
//
//
// Related Topics 图 数组 最短路 堆（优先队列） 👍 37 👎 0

package main

import "container/heap"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumTime(n int, edges [][]int, disappear []int) []int {
	adj := make([][]struct{ v, length int }, n)
	for _, edge := range edges {
		u, v, length := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], struct{ v, length int }{v, length})
		adj[v] = append(adj[v], struct{ v, length int }{u, length})
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{0, 0})
	answer := make([]int, n)
	for i := range answer {
		answer[i] = -1
	}
	answer[0] = 0
	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		t, u := item.priority, item.value
		if t != answer[u] {
			continue
		}
		for _, edge := range adj[u] {
			v, length := edge.v, edge.length
			if t+length < disappear[v] && (answer[v] == -1 || t+length < answer[v]) {
				heap.Push(pq, Item{t + length, v})
				answer[v] = t + length
			}
		}
	}
	return answer
}

type Item struct {
	priority, value int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

//leetcode submit region end(Prohibit modification and deletion)
