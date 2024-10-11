//给你一个数组 nums 和一个整数 k 。你需要找到 nums 的一个 子数组 ，满足子数组中所有元素按位或运算 OR 的值与 k 的 绝对差 尽可能 小
//。换言之，你需要选择一个子数组 nums[l..r] 满足 |k - (nums[l] OR nums[l + 1] ... OR nums[r])| 最小。
//
//
// 请你返回 最小 的绝对差值。
//
// 子数组 是数组中连续的 非空 元素序列。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,4,5], k = 3
//
//
// 输出：0
//
// 解释：
//
// 子数组 nums[0..1] 的按位 OR 运算值为 3 ，得到最小差值 |3 - 3| = 0 。
//
// 示例 2：
//
//
// 输入：nums = [1,3,1,3], k = 2
//
//
// 输出：1
//
// 解释：
//
// 子数组 nums[1..1] 的按位 OR 运算值为 3 ，得到最小差值 |3 - 2| = 1 。
//
// 示例 3：
//
//
// 输入：nums = [1], k = 10
//
//
// 输出：9
//
// 解释：
//
// 只有一个子数组，按位 OR 运算值为 1 ，得到最小差值 |10 - 1| = 9 。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
// 1 <= k <= 10⁹
//
//
// Related Topics 位运算 线段树 数组 二分查找 👍 38 👎 0

package main

import (
	"math"
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minimumDifference(nums []int, k int) int {
	n := len(nums)
	bitsMaxPos := make([]int, 31)
	for i := range bitsMaxPos {
		bitsMaxPos[i] = -1
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		for j := 0; j <= 30; j++ {
			if nums[i]>>j&1 == 1 {
				bitsMaxPos[j] = i
			}
		}

		posToBit := make([][2]int, 0)
		for j := 0; j <= 30; j++ {
			if bitsMaxPos[j] != -1 {
				posToBit = append(posToBit, [2]int{bitsMaxPos[j], j})
			}
		}
		sort.Slice(posToBit, func(a, b int) bool {
			return posToBit[a][0] > posToBit[b][0]
		})

		val := 0
		for j, p := 0, 0; j < len(posToBit); p = j {
			for j < len(posToBit) && posToBit[j][0] == posToBit[p][0] {
				val |= 1 << posToBit[j][1]
				j++
			}
			res = min(res, int(math.Abs(float64(val-k))))
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
