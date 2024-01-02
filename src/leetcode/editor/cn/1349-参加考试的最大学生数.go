//给你一个 m * n 的矩阵 seats 表示教室中的座位分布。如果座位是坏的（不可用），就用 '#' 表示；否则，用 '.' 表示。 
//
// 学生可以看到左侧、右侧、左上、右上这四个方向上紧邻他的学生的答卷，但是看不到直接坐在他前面或者后面的学生的答卷。请你计算并返回该考场可以容纳的同时参加考试
//且无法作弊的 最大 学生人数。 
//
// 学生必须坐在状况良好的座位上。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：seats = [["#",".","#","#",".","#"],
//              [".","#","#","#","#","."],
//              ["#",".","#","#",".","#"]]
//输出：4
//解释：教师可以让 4 个学生坐在可用的座位上，这样他们就无法在考试中作弊。 
// 
//
// 示例 2： 
//
// 
//输入：seats = [[".","#"],
//              ["#","#"],
//              ["#","."],
//              ["#","#"],
//              [".","#"]]
//输出：3
//解释：让所有学生坐在可用的座位上。
// 
//
// 示例 3： 
//
// 
//输入：seats = [["#",".",".",".","#"],
//              [".","#",".","#","."],
//              [".",".","#",".","."],
//              [".","#",".","#","."],
//              ["#",".",".",".","#"]]
//输出：10
//解释：让学生坐在第 1、3 和 5 列的可用座位上。
// 
//
// 
//
// 提示： 
//
// 
// seats 只包含字符 '.' 和'#' 
// m == seats.length 
// n == seats[i].length 
// 1 <= m <= 8 
// 1 <= n <= 8 
// 
//
// Related Topics 位运算 数组 动态规划 状态压缩 矩阵 👍 206 👎 0

package main

import (
	"math"
	"math/bits"
)

func main() {

}
//leetcode submit region begin(Prohibit modification and deletion)
func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	memo := make(map[int]int)

	isSingleRowCompliant := func(status, row int) bool {
		n := len(seats[0])
		for j := 0; j < n; j++ {
			if (status >> j) & 1 == 1 {
				if seats[row][j] == '#' {
					return false
				}
				if j > 0 && (status >> (j - 1)) & 1 == 1 {
					return false
				}
			}
		}
		return true
	}

	isCrossRowsCompliant := func(status, upperRowStatus int) bool {
		n := len(seats[0])
		for j := 0; j < n; j++ {
			if (status >> j) & 1 == 1 {
				if j > 0 && (upperRowStatus >> (j - 1)) & 1 == 1 {
					return false
				}
				if j < n - 1 && (upperRowStatus >> (j + 1)) & 1 == 1 {
					return false
				}
			}
		}
		return true
	}

	var dp func(int, int) int
	dp = func(row, status int) int {
		n := len(seats[0])
		key := (row << n) + status
		if _, ok := memo[key]; !ok {
			if !isSingleRowCompliant(status, row) {
				memo[key] = math.MinInt32
				return math.MinInt32
			}
			students := bits.OnesCount(uint(status))
			if row == 0 {
				memo[key] = students
				return students
			}
			mx := 0
			for upperRowStatus := 0; upperRowStatus < 1 << n; upperRowStatus++ {
				if isCrossRowsCompliant(status, upperRowStatus) {
					mx = max(mx, dp(row - 1, upperRowStatus))
				}
			}
			memo[key] = students + mx
		}
		return memo[key]
	}

	mx := 0
	for i := 0; i < (1 << n); i++ {
		mx = max(mx, dp(m - 1, i))
	}
	return mx
}

//leetcode submit region end(Prohibit modification and deletion)
