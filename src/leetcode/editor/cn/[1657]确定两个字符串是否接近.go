// 如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：
//
// 操作 1：交换任意两个 现有 字符。
//
// 例如，abcde -> aecdb
//
// 操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
//
// 例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
//
// 你可以根据需要对任意一个字符串多次使用这两种操作。
//
// 给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：word1 = "abc", word2 = "bca"
// 输出：true
// 解释：2 次操作从 word1 获得 word2 。
// 执行操作 1："abc" -> "acb"
// 执行操作 1："acb" -> "bca"
//
// 示例 2：
//
// 输入：word1 = "a", word2 = "aa"
// 输出：false
// 解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
//
// 示例 3：
//
// 输入：word1 = "cabbba", word2 = "abbccc"
// 输出：true
// 解释：3 次操作从 word1 获得 word2 。
// 执行操作 1："cabbba" -> "caabbb"
// 执行操作 2："caabbb" -> "baaccc"
// 执行操作 2："baaccc" -> "abbccc"
//
// 示例 4：
//
// 输入：word1 = "cabbba", word2 = "aabbss"
// 输出：false
// 解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
//
// 提示：
//
// 1 <= word1.length, word2.length <= 10⁵
// word1 和 word2 仅包含小写英文字母
//
// Related Topics 哈希表 字符串 排序 👍 124 👎 0
package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	ans := make([]int, 26)
	ans1 := make([]int, 26)
	for _, i2 := range word1 {
		ans[i2-'a']++
	}
	for _, i := range word2 {
		ans1[i-'a']++
	}
	for i := 0; i < 26; i++ {
		if (ans[i] > 0 && ans1[i] == 0) || (ans[i] == 0 && ans1[i] > 0) {
			return false
		}
	}
	sort.Ints(ans)
	sort.Ints(ans1)
	for i := 0; i < len(ans); i++ {
		if ans[i] != ans1[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
