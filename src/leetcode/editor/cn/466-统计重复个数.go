//定义 str = [s, n] 表示 str 由 n 个字符串 s 连接构成。
//
//
// 例如，str == ["abc", 3] =="abcabcabc" 。
//
//
// 如果可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。
//
//
// 例如，根据定义，s1 = "abc" 可以从 s2 = "abdbec" 获得，仅需要删除加粗且用斜体标识的字符。
//
//
// 现在给你两个字符串 s1 和 s2 和两个整数 n1 和 n2 。由此构造得到两个字符串，其中 str1 = [s1, n1]、str2 = [s2,
//n2] 。
//
// 请你找出一个最大整数 m ，以满足 str = [str2, m] 可以从 str1 获得。
//
//
//
// 示例 1：
//
//
//输入：s1 = "acb", n1 = 4, s2 = "ab", n2 = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：s1 = "acb", n1 = 1, s2 = "acb", n2 = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 100
// s1 和 s2 由小写英文字母组成
// 1 <= n1, n2 <= 10⁶
//
//
// Related Topics 字符串 动态规划 👍 249 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0 // 注意此处直接使用 Ra Rb 的下标，不取模

	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}

	map1, map2 := make(map[int]int), make(map[int]int)
	ans := 0 // 注意，此处存储的是 Ra 中 Sb 的个数，而非 Ra 中 Rb 的个数

	for index1/len1 < n1 { // 遍历整个 Ra
		if index1%len1 == len1-1 { //在 Sa 末尾
			if val, ok := map1[index2%len2]; ok { // 出现了循环，进行快进
				cycleLen := index1/len1 - val/len1                 // 每个循环占多少个 Sa
				cycleNum := (n1 - 1 - index1/len1) / cycleLen      // 还有多少个循环
				cycleS2Num := index2/len2 - map2[index2%len2]/len2 // 每个循环含有多少个 Sb

				index1 += cycleNum * cycleLen * len1 // 将 Ra 快进到相应的位置
				ans += cycleNum * cycleS2Num         // 把快进部分的答案数量加上
			} else { // 第一次，注意存储的是未取模的
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}

		}

		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				ans += 1
			}
			index2 += 1
		}
		index1 += 1
	}
	return ans / n2
}

//leetcode submit region end(Prohibit modification and deletion)
