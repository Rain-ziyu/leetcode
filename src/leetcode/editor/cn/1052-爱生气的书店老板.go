//有一个书店老板，他的书店开了 n 分钟。每分钟都有一些顾客进入这家商店。给定一个长度为 n 的整数数组 customers ，其中 customers[i]
// 是在第 i 分钟开始时进入商店的顾客数量，所有这些顾客在第 i 分钟结束后离开。
//
// 在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。
//
// 当书店老板生气时，那一分钟的顾客就会不满意，若老板不生气则顾客是满意的。
//
// 书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 minutes 分钟不生气，但却只能使用一次。
//
// 请你返回 这一天营业下来，最多有多少客户能够感到满意 。
//
// 示例 1：
//
//
//输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], minutes = 3
//输出：16
//解释：书店老板在最后 3 分钟保持冷静。
//感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.
//
//
// 示例 2：
//
//
//输入：customers = [1], grumpy = [0], minutes = 1
//输出：1
//
//
//
// 提示：
//
//
// n == customers.length == grumpy.length
// 1 <= minutes <= n <= 2 * 10⁴
// 0 <= customers[i] <= 1000
// grumpy[i] == 0 or 1
//
//
// Related Topics 数组 滑动窗口 👍 316 👎 0

package main

import "fmt"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	total := 0
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	increase := 0
	for i := 0; i < minutes; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase
	for i := minutes; i < n; i++ {
		increase = increase - customers[i-minutes]*grumpy[i-minutes] + customers[i]*grumpy[i]
		maxIncrease = max(maxIncrease, increase)
	}
	return total + maxIncrease
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
