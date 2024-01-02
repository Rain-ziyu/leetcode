//我们定义 arr 是 山形数组 当且仅当它满足：
//
//
// arr.length >= 3
// 存在某个下标 i （从 0 开始） 满足 0 < i < arr.length - 1 且：
//
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
//
//
//
// 给你整数数组 nums ，请你返回将 nums 变成 山形状数组 的 最少 删除次数。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,1]
//输出：0
//解释：数组本身就是山形数组，所以我们不需要删除任何元素。
//
//
// 示例 2：
//
//
//输入：nums = [2,1,1,5,6,2,3,1]
//输出：3
//解释：一种方法是将下标为 0，1 和 5 的元素删除，剩余元素为 [1,5,6,3,1] ，是山形数组。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// 1 <= nums[i] <= 10⁹
// 题目保证 nums 删除一些元素后一定能得到山形数组。
//
//
// Related Topics 贪心 数组 二分查找 动态规划 👍 87 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumMountainRemovals(nums []int) int {
	length := len(nums)
	zheng := getMaxArray(nums)
	fan := getMaxArray(reverse(nums))
	fan = reverse(fan)
	res := 0
	for i := 0; i < length; i++ {
		//题目要求0 < i < arr.length - 1因此不能是左右两边开始的因此必须大于1
		if zheng[i] > 1 && fan[i] > 1 {
			res = max(res, zheng[i]+fan[i]-1)
		}
	}
	return length - res
}

func reverse(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = nums[length-i-1]
	}
	return res
}

/*
计算当前数组的最长递增子序列 最短也是1 因为只有它本身
*/
func getMaxArray(nums []int) []int {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		maxValue := 1
		for cunrrent := 0; cunrrent < i; cunrrent++ {
			if nums[cunrrent] < nums[i] {
				maxValue = max(maxValue, dp[cunrrent]+1)
			}
		}
		dp[i] = maxValue
	}
	return dp
}

//leetcode submit region end(Prohibit modification and deletion)
