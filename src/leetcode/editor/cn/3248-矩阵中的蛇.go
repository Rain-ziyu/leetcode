//大小为 n x n 的矩阵 grid 中有一条蛇。蛇可以朝 四个可能的方向 移动。矩阵中的每个单元格都使用位置进行标识： grid[i][j] = (i *
// n) + j。
//
// 蛇从单元格 0 开始，并遵循一系列命令移动。
//
// 给你一个整数 n 表示 grid 的大小，另给你一个字符串数组 commands，其中包括 "UP"、"RIGHT"、"DOWN" 和 "LEFT"。题目
//测评数据保证蛇在整个移动过程中将始终位于 grid 边界内。
//
// 返回执行 commands 后蛇所停留的最终单元格的位置。
//
//
//
// 示例 1：
//
//
// 输入：n = 2, commands = ["RIGHT","DOWN"]
//
//
// 输出：3
//
// 解释：
//
//
//
//
//
// 0
// 1
//
//
// 2
// 3
//
//
//
//
//
//
//
//
// 0
// 1
//
//
// 2
// 3
//
//
//
//
//
//
//
// 0
// 1
//
//
// 2
// 3
//
//
//
//
// 示例 2：
//
//
// 输入：n = 3, commands = ["DOWN","RIGHT","UP"]
//
//
// 输出：1
//
// 解释：
//
//
//
//
//
// 0
// 1
// 2
//
//
// 3
// 4
// 5
//
//
// 6
// 7
// 8
//
//
//
//
//
//
//
//
// 0
// 1
// 2
//
//
// 3
// 4
// 5
//
//
// 6
// 7
// 8
//
//
//
//
//
//
//
// 0
// 1
// 2
//
//
// 3
// 4
// 5
//
//
// 6
// 7
// 8
//
//
//
//
//
//
//
// 0
// 1
// 2
//
//
// 3
// 4
// 5
//
//
// 6
// 7
// 8
//
//
//
//
//
//
// 提示：
//
//
// 2 <= n <= 10
// 1 <= commands.length <= 100
// commands 仅由 "UP"、"RIGHT"、"DOWN" 和 "LEFT" 组成。
// 生成的测评数据确保蛇不会移动到矩阵的边界外。
//
//
// Related Topics 数组 字符串 模拟 👍 22 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func finalPositionOfSnake(n int, commands []string) int {
	dis := make(map[string][]int)
	dis["UP"] = []int{0, -1}
	dis["RIGHT"] = []int{1,0}
	dis["DOWN"] = []int{0,1}
	dis["LEFT"] = []int{-1,0}
	first := []int{0,0}
	for _,com := range commands {
		first = []int{first[0]+dis[com][0],first[1]+dis[com][1]}
	}
	return first[0] +first[1]*n
}

//leetcode submit region end(Prohibit modification and deletion)
