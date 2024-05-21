//给你一个字符串 s 。请返回 s 中最长的 超赞子字符串 的长度。
//
// 「超赞子字符串」需满足满足下述两个条件：
//
//
// 该字符串是 s 的一个非空子字符串
// 进行任意次数的字符交换后，该字符串可以变成一个回文字符串
//
//
//
//
// 示例 1：
//
// 输入：s = "3242415"
//输出：5
//解释："24241" 是最长的超赞子字符串，交换其中的字符后，可以得到回文 "24142"
//
//
// 示例 2：
//
// 输入：s = "12345678"
//输出：1
//
//
// 示例 3：
//
// 输入：s = "213123"
//输出：6
//解释："213123" 是最长的超赞子字符串，交换其中的字符后，可以得到回文 "231132"
//
//
// 示例 4：
//
// 输入：s = "00"
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10^5
// s 仅由数字组成
//
//
// Related Topics 位运算 哈希表 字符串 👍 129 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func longestAwesome(s string) int {
	qianzhui := make([][]int, len(s)+1)
	qianzhui[0] = make([]int, 10)
	for i := range s {

		qianzhui[i+1] = make([]int, 10)
		copy(qianzhui[i+1], qianzhui[i])

		qianzhui[i+1][int(s[i]-'0')]++
	}
	res := 1
	for i := range qianzhui {
		end := len(s) - 1
		for end >= i {
			if check(qianzhui, i, end+1) {
				if end-i+1 > res {
					res = end - i + 1
				}
			}
			end--
		}
	}
	return res
}

func check(qianzhui [][]int, start int, end int) bool {
	res := 0
	for i := 0; i < 10; i++ {
		if (qianzhui[end][i]-qianzhui[start][i])%2 != 0 {
			res++
		}
	}
	return res <= 1
}

//leetcode submit region end(Prohibit modification and deletion)
