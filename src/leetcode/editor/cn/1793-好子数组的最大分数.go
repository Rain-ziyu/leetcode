//给你一个整数数组 nums （下标从 0 开始）和一个整数 k 。
//
// 一个子数组 (i, j) 的 分数 定义为 min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1) 。一个
// 好 子数组的两个端点下标需要满足 i <= k <= j 。
//
// 请你返回 好 子数组的最大可能 分数 。
//
//
//
// 示例 1：
//
// 输入：nums = [1,4,3,7,4,5], k = 3
//输出：15
//解释：最优子数组的左右端点下标是 (1, 5) ，分数为 min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15 。
//
//
// 示例 2：
//
// 输入：nums = [5,5,4,5,4,1,1,1], k = 0
//输出：20
//解释：最优子数组的左右端点下标是 (0, 4) ，分数为 min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 2 * 10⁴
// 0 <= k < nums.length
//
//
// Related Topics 栈 数组 双指针 二分查找 单调栈 👍 116 👎 0

package main

import (
	"fmt"
	"log"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumScore(nums []int, k int) int {
	left := k-1
	right := k+1
	res := nums[k]
	for i := nums[k]; ; i-- {
		for left >= 0 && nums[left] >= i {
			left--
		}
		for right < len(nums) && nums[right] >= i {
			right++
		}
		res = max(res, i*(right-left-1))
		if left < 0 && right >= len(nums) {
			break
		}
	}

	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
