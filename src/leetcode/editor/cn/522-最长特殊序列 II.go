//给定字符串列表 strs ，返回其中 最长的特殊序列 的长度。如果最长特殊序列不存在，返回 -1 。
//
// 特殊序列 定义如下：该序列为某字符串 独有的子序列（即不能是其他字符串的子序列）。
//
// s 的 子序列可以通过删去字符串 s 中的某些字符实现。
//
//
// 例如，"abc" 是 "aebdc" 的子序列，因为您可以删除"aebdc"中的下划线字符来得到 "abc" 。"aebdc"的子序列还包括
//"aebdc"、 "aeb" 和 "" (空字符串)。
//
//
//
//
// 示例 1：
//
//
//输入: strs = ["aba","cdc","eae"]
//输出: 3
//
//
// 示例 2:
//
//
//输入: strs = ["aaa","aaa","aa"]
//输出: -1
//
//
//
//
// 提示:
//
//
// 2 <= strs.length <= 50
// 1 <= strs[i].length <= 10
// strs[i] 只包含小写英文字母
//
//
// Related Topics 数组 哈希表 双指针 字符串 排序 👍 222 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	for i, n := 0, len(strs); i < n; i++ {
		tmp := true
		for i1 := 0; i1 < len(strs); i1++ {
			if i1 == i {
				continue
			}
			fmt.Println(!isZiXuLie(strs[i], strs[i1]))
			if len(strs[i1]) >= len(strs[i]) && !isZiXuLie(strs[i], strs[i1]) {
			} else if len(strs[i1]) < len(strs[i]) {
				break
			} else {
				tmp = false
				break
			}
		}
		if tmp {
			return len(strs[i])
		}
	}
	return -1
}
func isZiXuLie(str1 string, str2 string) bool {
	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}
	return i == len(str1)
}

//leetcode submit region end(Prohibit modification and deletion)
