//给你一个下标从 0 开始的数组 nums ，数组长度为 n 。 
//
// nums 的 不同元素数目差 数组可以用一个长度为 n 的数组 diff 表示，其中 diff[i] 等于前缀 nums[0, ..., i] 中不同元素
//的数目 减去 后缀 nums[i + 1, ..., n - 1] 中不同元素的数目。 
//
// 返回 nums 的 不同元素数目差 数组。 
//
// 注意 nums[i, ..., j] 表示 nums 的一个从下标 i 开始到下标 j 结束的子数组（包含下标 i 和 j 对应元素）。特别需要说明的是，
//如果 i > j ，则 nums[i, ..., j] 表示一个空子数组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3,4,5]
//输出：[-3,-1,1,3,5]
//解释：
//对于 i = 0，前缀中有 1 个不同的元素，而在后缀中有 4 个不同的元素。因此，diff[0] = 1 - 4 = -3 。
//对于 i = 1，前缀中有 2 个不同的元素，而在后缀中有 3 个不同的元素。因此，diff[1] = 2 - 3 = -1 。
//对于 i = 2，前缀中有 3 个不同的元素，而在后缀中有 2 个不同的元素。因此，diff[2] = 3 - 2 = 1 。
//对于 i = 3，前缀中有 4 个不同的元素，而在后缀中有 1 个不同的元素。因此，diff[3] = 4 - 1 = 3 。
//对于 i = 4，前缀中有 5 个不同的元素，而在后缀中有 0 个不同的元素。因此，diff[4] = 5 - 0 = 5 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [3,2,3,4,2]
//输出：[-2,-1,0,2,3]
//解释：
//对于 i = 0，前缀中有 1 个不同的元素，而在后缀中有 3 个不同的元素。因此，diff[0] = 1 - 3 = -2 。
//对于 i = 1，前缀中有 2 个不同的元素，而在后缀中有 3 个不同的元素。因此，diff[1] = 2 - 3 = -1 。
//对于 i = 2，前缀中有 2 个不同的元素，而在后缀中有 2 个不同的元素。因此，diff[2] = 2 - 2 = 0 。
//对于 i = 3，前缀中有 3 个不同的元素，而在后缀中有 1 个不同的元素。因此，diff[3] = 3 - 1 = 2 。
//对于 i = 4，前缀中有 3 个不同的元素，而在后缀中有 0 个不同的元素。因此，diff[4] = 3 - 0 = 3 。 
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n == nums.length <= 50 
// 1 <= nums[i] <= 50 
// 
//
// Related Topics 数组 哈希表 👍 30 👎 0

package main

import "fmt"

func main() {

}
//leetcode submit region begin(Prohibit modification and deletion)
func distinctDifferenceArray(nums []int) []int {
	set := make(map[int]bool)
	set1 := make(map[int]bool)
	valueMap := make(map[int][]int)
	for i := range nums {
		set[nums[i]] = true
		set1[nums[len(nums)-1-i]] = true
		if valueMap[i] == nil {
			valueMap[i] = make([]int, 2)
		}
		valueMap[i][0] = len(set)
		if valueMap[len(nums)-2-i] == nil {
			valueMap[len(nums)-2-i] = make([]int, 2)
		}
		valueMap[len(nums)-2-i][1] = len(set1)

	}
	fmt.Println(valueMap)
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = valueMap[i][0] - valueMap[i][1]
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
