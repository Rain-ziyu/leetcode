//给你一个下标从 0 开始的二维整数矩阵 grid，大小为 n * n ，其中的值在 [1, n²] 范围内。除了 a 出现 两次，b 缺失 之外，每个整数都
// 恰好出现一次 。 
//
// 任务是找出重复的数字a 和缺失的数字 b 。 
//
// 返回一个下标从 0 开始、长度为 2 的整数数组 ans ，其中 ans[0] 等于 a ，ans[1] 等于 b 。 
//
// 
//
// 示例 1： 
//
// 
//输入：grid = [[1,3],[2,2]]
//输出：[2,4]
//解释：数字 2 重复，数字 4 缺失，所以答案是 [2,4] 。
// 
//
// 示例 2： 
//
// 
//输入：grid = [[9,1,7],[8,9,2],[3,4,6]]
//输出：[9,5]
//解释：数字 9 重复，数字 5 缺失，所以答案是 [9,5] 。
// 
//
// 
//
// 提示： 
//
// 
// 2 <= n == grid.length == grid[i].length <= 50 
// 1 <= grid[i][j] <= n * n 
// 对于所有满足1 <= x <= n * n 的 x ，恰好存在一个 x 与矩阵中的任何成员都不相等。 
// 对于所有满足1 <= x <= n * n 的 x ，恰好存在一个 x 与矩阵中的两个成员相等。 
// 除上述的两个之外，对于所有满足1 <= x <= n * n 的 x ，都恰好存在一对 i, j 满足 0 <= i, j <= n - 1 且 
//grid[i][j] == x 。 
// 
//
// Related Topics 数组 哈希表 数学 矩阵 👍 15 👎 0

package main

import "fmt"

func main() {

}
//leetcode submit region begin(Prohibit modification and deletion)
func findMissingAndRepeatedValues(grid [][]int) []int {
	bo := make([]bool, len(grid)*len(grid)+1)
	res1:=0
	res2:=(len(grid)*len(grid)+1)*len(grid)*len(grid)/2
	fmt.Println(res2)
	for i := range grid {
		for i2 := range grid[i] {
			res2 -= grid[i][i2]
			if bo[grid[i][i2]] {
				res1 = grid[i][i2]
				res2 += res1
			}
			bo[grid[i][i2]] = true
		}
	}
	return []int{res1,res2}
}
//leetcode submit region end(Prohibit modification and deletion)
