//给你一个下标从 1 开始、大小为 m x n 的整数矩阵 mat，你可以选择任一单元格作为 起始单元格 。
//
// 从起始单元格出发，你可以移动到 同一行或同一列 中的任何其他单元格，但前提是目标单元格的值 严格大于 当前单元格的值。
//
// 你可以多次重复这一过程，从一个单元格移动到另一个单元格，直到无法再进行任何移动。
//
// 请你找出从某个单元开始访问矩阵所能访问的 单元格的最大数量 。
//
// 返回一个表示可访问单元格最大数量的整数。
//
//
//
// 示例 1：
//
//
//
// 输入：mat = [[3,1],[3,4]]
//输出：2
//解释：上图展示了从第 1 行、第 2 列的单元格开始，可以访问 2 个单元格。可以证明，无论从哪个单元格开始，最多只能访问 2 个单元格，因此答案是 2 。
//
//
//
// 示例 2：
//
//
//
// 输入：mat = [[1,1],[1,1]]
//输出：1
//解释：由于目标单元格必须严格大于当前单元格，在本示例中只能访问 1 个单元格。
//
//
// 示例 3：
//
//
//
// 输入：mat = [[3,1,6],[-9,5,7]]
//输出：4
//解释：上图展示了从第 2 行、第 1 列的单元格开始，可以访问 4 个单元格。可以证明，无论从哪个单元格开始，最多只能访问 4 个单元格，因此答案是 4 。
//
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 10⁵
// 1 <= m * n <= 10⁵
// -10⁵ <= mat[i][j] <= 10⁵
//
//
// Related Topics 记忆化搜索 数组 哈希表 二分查找 动态规划 矩阵 有序集合 排序 👍 71 👎 0

package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxIncreasingCells(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	mp := make(map[int][][2]int)
	row := make([]int, m)
	col := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[mat[i][j]] = append(mp[mat[i][j]], [2]int{i, j})
		}
	}
	keys := []int{}
	for key, _ := range mp {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		pos := mp[key]
		res := []int{} // 存放相同数值的答案，便于后续更新 row 和 col
		for _, arr := range pos {
			res = append(res, max(row[arr[0]], col[arr[1]])+1)
		}
		for i := 0; i < len(pos); i++ {
			arr, d := pos[i], res[i]
			row[arr[0]] = max(row[arr[0]], d)
			col[arr[1]] = max(col[arr[1]], d)
		}
	}

	return maxSlice(row)
}

func maxSlice(slice []int) int {
	maxVal := slice[0]
	for _, val := range slice {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

//leetcode submit region end(Prohibit modification and deletion)
