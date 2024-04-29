//给你一个整数 n ，以二进制字符串的形式返回该整数的 负二进制（base -2）表示。
//
// 注意，除非字符串就是 "0"，否则返回的字符串中不能含有前导零。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出："110"
//解释：(-2)² + (-2)¹ = 2
//
//
// 示例 2：
//
//
//输入：n = 3
//输出："111"
//解释：(-2)² + (-2)¹ + (-2)⁰ = 3
//
//
// 示例 3：
//
//
//输入：n = 4
//输出："100"
//解释：(-2)² = 4
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁹
//  00 10 11  01
//  0  -2 -1  1
// Related Topics 数学 👍 205 👎 0

package main

import (
	"math"
	"strconv"
)

func main() {
	baseNeg2(2)
}

// leetcode submit region begin(Prohibit modification and deletion)
func baseNeg2(n int) string {
	res := 0
	i := 0
	for n > 0 {
		if (i+1)%2 == 0 {
			res = res  + (n%2)*int(math.Pow(float64(2), float64(i)))
			n = (n / 2) + (n%2)
		} else {
			res = res + (n%2)*int(math.Pow(float64(2), float64(i)))
			n = n / 2
		}
		i++
	}
	finalStr := ""
	for i := res; i > 0; {
		finalStr = strconv.Itoa(i%2) + finalStr
		i = i / 2
	}
	if res == 0 {
		return "0"
	}
	return finalStr
}

//leetcode submit region end(Prohibit modification and deletion)
