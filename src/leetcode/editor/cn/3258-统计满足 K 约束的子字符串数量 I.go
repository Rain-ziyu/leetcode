//给你一个 二进制 字符串 s 和一个整数 k。 
//
// 如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束： 
//
// 
// 字符串中 0 的数量最多为 k。 
// 字符串中 1 的数量最多为 k。 
// 
//
// 返回一个整数，表示 s 的所有满足 k 约束 的子字符串的数量。 
//
// 
//
// 示例 1： 
//
// 
// 输入：s = "10101", k = 1 
// 
//
// 输出：12 
//
// 解释： 
//
// s 的所有子字符串中，除了 "1010"、"10101" 和 "0101" 外，其余子字符串都满足 k 约束。 
//
// 示例 2： 
//
// 
// 输入：s = "1010101", k = 2 
// 
//
// 输出：25 
//
// 解释： 
//
// s 的所有子字符串中，除了长度大于 5 的子字符串外，其余子字符串都满足 k 约束。 
//
// 示例 3： 
//
// 
// 输入：s = "11111", k = 1 
// 
//
// 输出：15 
//
// 解释： 
//
// s 的所有子字符串都满足 k 约束。 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 50 
// 1 <= k <= s.length 
// s[i] 是 '0' 或 '1'。 
// 
//
// Related Topics 字符串 滑动窗口 👍 33 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func countKConstraintSubstrings(s string, k int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		zero := 0
		one := 0
		for j := i; j < len(s); j++ {
			if s[j] == '0'{
				zero++
			}else {
				one ++
			}
			if zero<=k||one<=k{
				res++
			}else {
				break
			}
		}
	}
	return  res
}
//leetcode submit region end(Prohibit modification and deletion)
