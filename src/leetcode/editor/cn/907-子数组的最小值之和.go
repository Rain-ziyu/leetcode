//给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
//
// 由于答案可能很大，因此 返回答案模 10^9 + 7 。
//
//
//
// 示例 1：
//
//
//输入：arr = [3,1,2,4]
//输出：17
//解释：
//子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
//最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
//
// 示例 2：
//
//
//输入：arr = [11,81,94,43,3]
//输出：444
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 3 * 10⁴
// 1 <= arr[i] <= 3 * 10⁴
//
//设x（起始位置），y位置的数组
// f(x,y) = min(f(x,y-1),arr[y])
//
// Related Topics 栈 数组 动态规划 单调栈 👍 693 👎 0

package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sumSubarrayMins1(arr []int) int {
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = math.MaxInt
	}
	res := 0
	//枚举长度
	for index := 0; index < len(arr); index++ {
		//枚举起始位置
		for i := 0; i+index < len(arr); i++ {
			dp[i] = min(dp[i], arr[i+index])
			res += dp[i]
		}
	}
	return res % (1000000000 + 7)
}
func sumSubarrayMins(arr []int) int {
	dp := make([]int, len(arr))
	monoStack := []int{}
	res := 0
	//枚举以arr[i]结束的子串
	for i, value := range arr {
		//循环枚举以arr[i]结尾时 arr[i]可以作为最小值的场景
		for len(monoStack) > 0 && arr[monoStack[len(monoStack)-1]] > value {
			//出栈操作将数组最后一位删除  保证最终栈顶是大于arr[i]的
			monoStack = monoStack[:len(monoStack)-1]
		}
		k := i + 1
		dp[i] = k * value
		//如果没有全部出栈，表示arr[i]作为最小值的子字符串仅有i-monoStack[len(monoStack)-1]个
		if len(monoStack) > 0 {
			k = i - monoStack[len(monoStack)-1]
		}
		dp[i] = k * value
		if len(monoStack) > 0 {
			dp[i] += dp[i-k]
		}
		res += dp[i]
		//将当前值入栈等待比较
		monoStack = append(monoStack, i)
	}
	return res % (1000000000 + 7)
}

//leetcode submit region end(Prohibit modification and deletion)
