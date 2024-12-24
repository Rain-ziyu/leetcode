//给你一个字符串 s ，它由某个字符串 t 和若干 t 的 同位字符串 连接而成。
//
// 请你返回字符串 t 的 最小 可能长度。
//
// 同位字符串 指的是重新排列一个单词得到的另外一个字符串，原来字符串中的每个字符在新字符串中都恰好只使用一次。
//
//
//
// 示例 1：
//
//
// 输入：s = "abba"
//
//
// 输出：2
//
// 解释：
//
// 一个可能的字符串 t 为 "ba" 。
//
// 示例 2：
//
//
// 输入：s = "cdef"
//
//
// 输出：4
//
// 解释：
//
// 一个可能的字符串 t 为 "cdef" ，注意 t 可能等于 s 。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s 只包含小写英文字母。
//
//
// Related Topics 哈希表 字符串 计数 👍 35 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func minAnagramLength(s string) int {
	n := len(s)
	check := func(m int) bool {
		var count0 [26]int
		for j := 0; j < n; j += m {
			var count1 [26]int
			for k := j; k < j+m; k++ {
				count1[s[k]-'a']++
			}
			if j > 0 && count0 != count1 {
				return false
			}
			count0 = count1
		}
		return true
	}
	for i := 1; i < n; i++ {
		if n%i != 0 {
			continue
		}
		if check(i) {
			return i
		}
	}
	return n
}

//leetcode submit region end(Prohibit modification and deletion)
