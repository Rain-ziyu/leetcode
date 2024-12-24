//给你一个字符串数组 words 和一个字符串 target。
//
// 如果字符串 x 是 words 中 任意 字符串的 前缀，则认为 x 是一个 有效 字符串。
//
// 现计划通过 连接 有效字符串形成 target ，请你计算并返回需要连接的 最少 字符串数量。如果无法通过这种方式形成 target，则返回 -1。
//
//
//
// 示例 1：
//
//
// 输入： words = ["abc","aaaaa","bcdef"], target = "aabcdabc"
//
//
// 输出： 3
//
// 解释：
//
// target 字符串可以通过连接以下有效字符串形成：
//
//
// words[1] 的长度为 2 的前缀，即 "aa"。
// words[2] 的长度为 3 的前缀，即 "bcd"。
// words[0] 的长度为 3 的前缀，即 "abc"。
//
//
// 示例 2：
//
//
// 输入： words = ["abababab","ab"], target = "ababaababa"
//
//
// 输出： 2
//
// 解释：
//
// target 字符串可以通过连接以下有效字符串形成：
//
//
// words[0] 的长度为 5 的前缀，即 "ababa"。
// words[0] 的长度为 5 的前缀，即 "ababa"。
//
//
// 示例 3：
//
//
// 输入： words = ["abcdef"], target = "xyz"
//
//
// 输出： -1
//
//
//
// 提示：
//
//
// 1 <= words.length <= 100
// 1 <= words[i].length <= 5 * 10⁴
// 输入确保 sum(words[i].length) <= 10⁵.
// words[i] 只包含小写英文字母。
// 1 <= target.length <= 5 * 10⁴
// target 只包含小写英文字母。
//
//
// Related Topics 线段树 数组 字符串 二分查找 动态规划 字符串匹配 哈希函数 滚动哈希 👍 32 👎 0

package main

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minValidStrings(words []string, target string) int {
	prefixFunction := func(word, target string) []int {
		s := word + "#" + target
		n := len(s)
		pi := make([]int, n)
		for i := 1; i < n; i++ {
			j := pi[i-1]
			for j > 0 && s[i] != s[j] {
				j = pi[j-1]
			}
			if s[i] == s[j] {
				j++
			}
			pi[i] = j
		}
		return pi
	}

	n := len(target)
	back := make([]int, n)
	for _, word := range words {
		pi := prefixFunction(word, target)
		fmt.Print(pi)
		m := len(word)
		for i := 0; i < n; i++ {
			back[i] = int(math.Max(float64(back[i]), float64(pi[m+1+i])))
		}
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = int(1e9)
	}
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i+1-back[i]] + 1
		if dp[i+1] > n {
			return -1
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
