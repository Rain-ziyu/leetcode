//给你一个整数数组 nums。
//
// 返回两个（不一定不同的）质数在 nums 中 下标 的 最大距离。
//
//
//
// 示例 1：
//
//
// 输入： nums = [4,2,9,5,3]
//
//
// 输出： 3
//
// 解释： nums[1]、nums[3] 和 nums[4] 是质数。因此答案是 |4 - 1| = 3。
//
// 示例 2：
//
//
// 输入： nums = [4,8,2,8]
//
//
// 输出： 0
//
// 解释： nums[2] 是质数。因为只有一个质数，所以答案是 |2 - 2| = 0。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁵
// 1 <= nums[i] <= 100
// 输入保证 nums 中至少有一个质数。
//
//
// Related Topics 数组 数学 数论 👍 24 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumPrimeDifference(nums []int) int {
	zhishumap := make(map[int]bool)
	//放入100以内的质数
	zhishu := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97}
	for _, v := range zhishu {
		zhishumap[v] = true
	}
	for i := 0; i < len(nums); i++ {
		if zhishumap[nums[i]] {
			for j := len(nums)-1; j >=i ; j-- {
				if zhishumap[nums[j]] {
					return j - i
				}
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
