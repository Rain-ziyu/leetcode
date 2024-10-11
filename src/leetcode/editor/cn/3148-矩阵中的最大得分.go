//给你一个由 正整数 组成、大小为 m x n 的矩阵 grid。你可以从矩阵中的任一单元格移动到另一个位于正下方或正右侧的任意单元格（不必相邻）。从值为 
//c1 的单元格移动到值为 c2 的单元格的得分为 c2 - c1 。 
//
// 你可以从 任一 单元格开始，并且必须至少移动一次。 
//
// 返回你能得到的 最大 总得分。 
//
// 
//
// 示例 1： 
// 
// 
// 输入：grid = [[9,5,7,3],[8,9,6,1],[6,7,14,3],[2,5,3,1]] 
// 
//
// 输出：9 
//
// 解释：从单元格 (0, 1) 开始，并执行以下移动： - 从单元格 (0, 1) 移动到 (2, 1)，得分为 7 - 5 = 2 。 - 从单元格 (2
//, 1) 移动到 (2, 2)，得分为 14 - 7 = 7 。 总得分为 2 + 7 = 9 。 
//
// 示例 2： 
//
// 
//
// 
// 输入：grid = [[4,3,2],[3,2,1]] 
// 
//
// 输出：-1 
//
// 解释：从单元格 (0, 0) 开始，执行一次移动：从 (0, 0) 到 (0, 1) 。得分为 3 - 4 = -1 。 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 2 <= m, n <= 1000 
// 4 <= m * n <= 10⁵ 
// 1 <= grid[i][j] <= 10⁵ 
// 
//
// Related Topics 数组 动态规划 矩阵 👍 45 👎 0

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	prerow := make([][]int, m)
	precol := make([][]int, m)
	f := make([][]int, m)

	for i := range prerow {
		prerow[i] = make([]int, n)
		precol[i] = make([]int, n)
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MinInt32
		}
	}

	ans := math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				f[i][j] = max(f[i][j], grid[i][j] + precol[i - 1][j])
			}
			if j > 0 {
				f[i][j] = max(f[i][j], grid[i][j] + prerow[i][j - 1])
			}
			ans = max(ans, f[i][j])
			prerow[i][j] = max(f[i][j], 0) - grid[i][j]
			precol[i][j] = prerow[i][j]
			if i > 0 {
				precol[i][j] = max(precol[i][j], precol[i - 1][j])
			}
			if j > 0 {
				prerow[i][j] = max(prerow[i][j], prerow[i][j - 1])
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
