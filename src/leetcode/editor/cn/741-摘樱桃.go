//给你一个 n x n 的网格 grid ，代表一块樱桃地，每个格子由以下三种数字的一种来表示：
//
//
// 0 表示这个格子是空的，所以你可以穿过它。
// 1 表示这个格子里装着一个樱桃，你可以摘到樱桃然后穿过它。
// -1 表示这个格子里有荆棘，挡着你的路。
//
//
// 请你统计并返回：在遵守下列规则的情况下，能摘到的最多樱桃数：
//
//
// 从位置 (0, 0) 出发，最后到达 (n - 1, n - 1) ，只能向下或向右走，并且只能穿越有效的格子（即只可以穿过值为 0 或者 1 的格子）；
//
// 当到达 (n - 1, n - 1) 后，你要继续走，直到返回到 (0, 0) ，只能向上或向左走，并且只能穿越有效的格子；
// 当你经过一个格子且这个格子包含一个樱桃时，你将摘到樱桃并且这个格子会变成空的（值变为 0 ）；
// 如果在 (0, 0) 和 (n - 1, n - 1) 之间不存在一条可经过的路径，则无法摘到任何一个樱桃。
//
//
//
//
// 示例 1：
//
//
//输入：grid = [[0,1,-1],[1,0,-1],[1,1,1]]
//输出：5
//解释：玩家从 (0, 0) 出发：向下、向下、向右、向右移动至 (2, 2) 。
//在这一次行程中捡到 4 个樱桃，矩阵变成 [[0,1,-1],[0,0,-1],[0,0,0]] 。
//然后，玩家向左、向上、向上、向左返回起点，再捡到 1 个樱桃。
//总共捡到 5 个樱桃，这是最大可能值。
//
//
// 示例 2：
//
//
//输入：grid = [[1,1,-1],[1,-1,1],[-1,1,1]]
//输出：0
//
//
//
//
// 提示：
//
//
// n == grid.length
// n == grid[i].length
// 1 <= n <= 50
// grid[i][j] 为 -1、0 或 1
// grid[0][0] != -1
// grid[n - 1][n - 1] != -1
//
//
// Related Topics 数组 动态规划 矩阵 👍 429 👎 0

package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func cherryPickup(grid [][]int) int {
	n := len(grid)
	f := make([][][]int, n*2-1)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt32
			}
		}
	}
	f[0][0][0] = grid[0][0]
	for k := 1; k < n*2-1; k++ {
		for x1 := max(k-n+1, 0); x1 <= min(k, n-1); x1++ {
			y1 := k - x1
			if grid[x1][y1] == -1 {
				continue
			}
			for x2 := x1; x2 <= min(k, n-1); x2++ {
				y2 := k - x2
				if grid[x2][y2] == -1 {
					continue
				}
				res := f[k-1][x1][x2] // 都往右
				if x1 > 0 {
					res = max(res, f[k-1][x1-1][x2]) // 往下，往右
				}
				if x2 > 0 {
					res = max(res, f[k-1][x1][x2-1]) // 往右，往下
				}
				if x1 > 0 && x2 > 0 {
					res = max(res, f[k-1][x1-1][x2-1]) // 都往下
				}
				res += grid[x1][y1]
				if x2 != x1 { // 避免重复摘同一个樱桃
					res += grid[x2][y2]
				}
				f[k][x1][x2] = res
			}
		}
	}
	return max(f[n*2-2][n-1][n-1], 0)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
