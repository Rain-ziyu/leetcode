//给你两个字符串 coordinate1 和 coordinate2，代表 8 x 8 国际象棋棋盘上的两个方格的坐标。
//
// 以下是棋盘的参考图。
//
//
//
// 如果这两个方格颜色相同，返回 true，否则返回 false。
//
// 坐标总是表示有效的棋盘方格。坐标的格式总是先字母（表示列），再数字（表示行）。
//
//
//
// 示例 1：
//
//
// 输入： coordinate1 = "a1", coordinate2 = "c3"
//
//
// 输出： true
//
// 解释：
//
// 两个方格均为黑色。
//
// 示例 2：
//
//
// 输入： coordinate1 = "a1", coordinate2 = "h3"
//
//
// 输出： false
//
// 解释：
//
// 方格 "a1" 是黑色，而 "h3" 是白色。
//
//
//
// 提示：
//
//
// coordinate1.length == coordinate2.length == 2
// 'a' <= coordinate1[0], coordinate2[0] <= 'h'
// '1' <= coordinate1[1], coordinate2[1] <= '8'
//
//
// Related Topics 数学 字符串 👍 23 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	u := coordinate1[0] - 'a'
	u2 := coordinate1[1] - '0'
	u3 := coordinate2[0] - 'a'
	u4 := coordinate2[1] - '0'
	return (u+u2)%2 == (u3+u4)%2
}

//leetcode submit region end(Prohibit modification and deletion)
