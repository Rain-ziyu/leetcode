//给你一个字符串 word ，你可以向其中任何位置插入 "a"、"b" 或 "c" 任意次，返回使 word 有效 需要插入的最少字母数。
//
// 如果字符串可以由 "abc" 串联多次得到，则认为该字符串 有效 。
//
//
//
// 示例 1：
//
// 输入：word = "b"
//输出：2
//解释：在 "b" 之前插入 "a" ，在 "b" 之后插入 "c" 可以得到有效字符串 "abc" 。
//
//
// 示例 2：
//
// 输入：word = "aaa"
//输出：6
//解释：在每个 "a" 之后依次插入 "b" 和 "c" 可以得到有效字符串 "abcabcabc" 。
//
//
// 示例 3：
//
// 输入：word = "abc"
//输出：0
//解释：word 已经是有效字符串，不需要进行修改。
//
//
//
//
// 提示：
//
//
// 1 <= word.length <= 50
// word 仅由字母 "a"、"b" 和 "c" 组成。
//
//
// Related Topics 栈 贪心 字符串 动态规划 👍 68 👎 0

package main

import "fmt"

func main() {
	fmt.Println(addMinimum("b"))
}

// leetcode submit region begin(Prohibit modification and deletion)
func addMinimum(word string) int {
	dp := make([]uint8, len(word)+1)
	pre := uint8('c')
	for i := 0; i < len(word); i++ {
		if word[i] == pre+1 || (word[i] == 'a' && pre == 'c') {
			dp[i+1] = dp[i]
		} else if word[i] == pre {
			dp[i+1] = dp[i] + 2
		} else {
			dp[i+1] = dp[i] + 1
		}
		pre = word[i]
	}
	if word[len(word)-1] == 'c' {
		return int(dp[len(dp)-1])
	}
	if word[len(word)-1] == 'b' {
		return int(dp[len(dp)-1] + 1)
	}
	return int(dp[len(dp)-1] + 2)

}

//leetcode submit region end(Prohibit modification and deletion)
