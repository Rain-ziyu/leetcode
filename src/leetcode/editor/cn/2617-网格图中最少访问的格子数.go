//给你一个下标从 0 开始的 m x n 整数矩阵 grid 。你一开始的位置在 左上角 格子 (0, 0) 。
//
// 当你在格子 (i, j) 的时候，你可以移动到以下格子之一：
//
//
// 满足 j < k <= grid[i][j] + j 的格子 (i, k) （向右移动），或者
// 满足 i < k <= grid[i][j] + i 的格子 (k, j) （向下移动）。
//
//
// 请你返回到达 右下角 格子 (m - 1, n - 1) 需要经过的最少移动格子数，如果无法到达右下角格子，请你返回 -1 。
//
//
//
// 示例 1：
//
//
//
// 输入：grid = [[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]]
//输出：4
//解释：上图展示了到达右下角格子经过的 4 个格子。
//
//
// 示例 2：
//
//
//
// 输入：grid = [[3,4,2,1],[4,2,1,1],[2,1,1,0],[3,4,1,0]]
//输出：3
//解释：上图展示了到达右下角格子经过的 3 个格子。
//
//
// 示例 3：
//
//
//
// 输入：grid = [[2,1,0],[1,0,0]]
//输出：-1
//解释：无法到达右下角格子。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10⁵
// 1 <= m * n <= 10⁵
// 0 <= grid[i][j] < m * n
// grid[m - 1][n - 1] == 0
//
//
// Related Topics 栈 广度优先搜索 并查集 数组 动态规划 矩阵 单调栈 堆（优先队列） 👍 52 👎 0

package main

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)
func minimumVisitedCells(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[0][0] = 1
	row := make([]PriorityQueue, m)
	col := make([]PriorityQueue, n)

	update := func(x *int, y int) {
		if *x == -1 || y < *x {
			*x = y
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for len(row[i]) > 0 && row[i][0].second+grid[i][row[i][0].second] < j {
				heap.Pop(&row[i])
			}
			if len(row[i]) > 0 {
				update(&dist[i][j], dist[i][row[i][0].second]+1)
			}

			for len(col[j]) > 0 && col[j][0].second+grid[col[j][0].second][j] < i {
				heap.Pop(&col[j])
			}
			if len(col[j]) > 0 {
				update(&dist[i][j], dist[col[j][0].second][j]+1)
			}
			if dist[i][j] != -1 {
				heap.Push(&row[i], Pair{dist[i][j], j})
				heap.Push(&col[j], Pair{dist[i][j], i})
			}
		}
	}
	return dist[m-1][n-1]
}

type Pair struct {
	first  int
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
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
