//给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素
// 。
//
// 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
//。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//
//
// 示例 2:
//
//
//输入: nums = [1,2,3,4,3]
//输出: [2,3,4,-1,4]
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 栈 数组 单调栈 👍 966 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	stack := []int{}
	stack = append(stack, 0)
	res := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		// 栈不为空 且 当前元素大于栈顶元素 说明需要取出栈顶元素
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := len(stack) - 1; i >= 0; i-- {
		for i2 := 0; i2 < len(nums); i2++ {
			if len(stack) == 0 {
				break
			}
			current := (stack[i] + i2 + 1) % len(nums)
			for len(stack) > 0 && nums[current] > nums[stack[len(stack)-1]] {
				res[stack[len(stack)-1]] = nums[current]
				stack = stack[:len(stack)-1]
				i--
			}
			if len(stack) > 0 && (stack[i]+i2+1)-len(nums) == stack[len(stack)-1] {
				res[stack[len(stack)-1]] = -1
				stack = stack[:len(stack)-1]
			}

		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
