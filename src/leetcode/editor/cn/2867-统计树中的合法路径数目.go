//给你一棵 n 个节点的无向树，节点编号为 1 到 n 。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[i] =
// [ui, vi] 表示节点 ui 和 vi 在树中有一条边。
//
// 请你返回树中的 合法路径数目 。
//
// 如果在节点 a 到节点 b 之间 恰好有一个 节点的编号是质数，那么我们称路径 (a, b) 是 合法的 。
//
// 注意：
//
//
// 路径 (a, b) 指的是一条从节点 a 开始到节点 b 结束的一个节点序列，序列中的节点 互不相同 ，且相邻节点之间在树上有一条边。
// 路径 (a, b) 和路径 (b, a) 视为 同一条 路径，且只计入答案 一次 。
//
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 5, edges = [[1,2],[1,3],[2,4],[2,5]]
//输出：4
//解释：恰好有一个质数编号的节点路径有：
//- (1, 2) 因为路径 1 到 2 只包含一个质数 2 。
//- (1, 3) 因为路径 1 到 3 只包含一个质数 3 。
//- (1, 4) 因为路径 1 到 4 只包含一个质数 2 。
//- (2, 4) 因为路径 2 到 4 只包含一个质数 2 。
//只有 4 条合法路径。
//
//
// 示例 2：
//
//
//
//
//输入：n = 6, edges = [[1,2],[1,3],[2,4],[3,5],[3,6]]
//输出：6
//解释：恰好有一个质数编号的节点路径有：
//- (1, 2) 因为路径 1 到 2 只包含一个质数 2 。
//- (1, 3) 因为路径 1 到 3 只包含一个质数 3 。
//- (1, 4) 因为路径 1 到 4 只包含一个质数 2 。
//- (1, 6) 因为路径 1 到 6 只包含一个质数 3 。
//- (2, 4) 因为路径 2 到 4 只包含一个质数 2 。
//- (3, 6) 因为路径 3 到 6 只包含一个质数 3 。
//只有 6 条合法路径。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
// edges.length == n - 1
// edges[i].length == 2
// 1 <= ui, vi <= n
// 输入保证 edges 形成一棵合法的树。
//
//
// Related Topics 树 深度优先搜索 数学 动态规划 数论 👍 55 👎 0

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	countPaths(5, [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}})
}

// leetcode submit region begin(Prohibit modification and deletion)
func countPaths(n int, edges [][]int) int64 {
	//	构造无向树的邻接矩阵
	adj := make([][]int, n+1)
	for _, value := range edges {
		adj[value[1]] = append(adj[value[1]], value[0])
		adj[value[0]] = append(adj[value[0]], value[1])
	}
	fmt.Println("最终的邻接表是：", adj)
	var dfs func(int, int, []int, int) []int
	dfs = func(prevalue int, preprevalue int, next []int, precount int) []int {
		res := []int{}
		if precount == 1 {
			for i := range next {
				if next[i] != preprevalue && !is(next[i]) {
					res = append(res, next[i])
					strings := dfs(next[i], prevalue, adj[next[i]], 1)
					res = append(res, strings...)
				}
			}
		} else {
			for i := range next {
				if next[i] != preprevalue && is(next[i]) {
					res = append(res, next[i])
					strings := dfs(next[i], prevalue, adj[next[i]], 1)
					res = append(res, strings...)
				} else if next[i] != preprevalue {
					strings := dfs(next[i], prevalue, adj[next[i]], 0)
					res = append(res, strings...)
				}
			}
		}
		return res
	}
	var res = make(map[string]bool)
	for i := 1; i < n+1; i++ {
		b := is(i)
		ispre := 0
		if b {
			ispre = 1
		}
		strings := dfs(i, 0, adj[i], ispre)
		for i2 := range strings {
			res[strconv.Itoa(i)+":"+strconv.Itoa(strings[i2])] = true
		}
	}
	return int64(len(res) / 2)
}
func is(number int) bool {
	if number <= 1 {
		return false
	}

	// 2是唯一的偶数质数
	if number == 2 {
		return true
	}

	// 排除其他偶数
	if number%2 == 0 {
		return false
	}

	// 只需检查奇数的可能因子
	maxDivisor := int(math.Sqrt(float64(number)))

	for i := 3; i <= maxDivisor; i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
