////给你一个正整数 n。 
////
//// 如果一个二进制字符串 x 的所有长度为 2 的子字符串中包含 至少 一个 "1"，则称 x 是一个 有效 字符串。 
////
//// 返回所有长度为 n 的 有效 字符串，可以以任意顺序排列。 
////
//// 
////
//// 示例 1： 
////
//// 
//// 输入： n = 3 
//// 
////
//// 输出： ["010","011","101","110","111"] 
////
//// 解释： 
////
//// 长度为 3 的有效字符串有："010"、"011"、"101"、"110" 和 "111"。 
////
//// 示例 2： 
////
//// 
//// 输入： n = 1 
//// 
////
//// 输出： ["0","1"] 
////
//// 解释： 
////
//// 长度为 1 的有效字符串有："0" 和 "1"。 
////
//// 
////
//// 提示： 
////
//// 
//// 1 <= n <= 18 
//// 
////
//// Related Topics 位运算 字符串 回溯 👍 30 👎 0
//

package main
//leetcode submit region begin(Prohibit modification and deletion)
func validStrings(n int) []string {
	tmp := make(map[int][]string)
	tmp[0] = []string{"0"}
	tmp[1] = []string{"1"}
	for i := 2; i <= n; i++ {
		tmp1 := make(map[int][]string)
		for i2 := range tmp[0] {
			tmp1[1] = append(tmp1[1],tmp[0][i2]+"1")
		}
		for i2 := range tmp[1] {
			tmp1[1] = append(tmp1[1],tmp[1][i2]+"1")
			tmp1[0] = append(tmp1[0],tmp[1][i2]+"0")
		}
		tmp = tmp1
	}
	return append(tmp[0],tmp[1]...)
}
//leetcode submit region end(Prohibit modification and deletion)
