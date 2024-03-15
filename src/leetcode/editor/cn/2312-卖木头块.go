//给你两个整数 m 和 n ，分别表示一块矩形木块的高和宽。同时给你一个二维整数数组 prices ，其中 prices[i] = [hi, wi,
//pricei] 表示你可以以 pricei 元的价格卖一块高为 hi 宽为 wi 的矩形木块。
//
// 每一次操作中，你必须按下述方式之一执行切割操作，以得到两块更小的矩形木块：
//
//
// 沿垂直方向按高度 完全 切割木块，或
// 沿水平方向按宽度 完全 切割木块
//
//
// 在将一块木块切成若干小木块后，你可以根据 prices 卖木块。你可以卖多块同样尺寸的木块。你不需要将所有小木块都卖出去。你 不能 旋转切好后木块的高和宽
//。
//
// 请你返回切割一块大小为 m x n 的木块后，能得到的 最多 钱数。
//
// 注意你可以切割木块任意次。
//
//
//
// 示例 1：
//
//
//
//
//输入：m = 3, n = 5, prices = [[1,4,2],[2,2,7],[2,1,3]]
//输出：19
//解释：上图展示了一个可行的方案。包括：
//- 2 块 2 x 2 的小木块，售出 2 * 7 = 14 元。
//- 1 块 2 x 1 的小木块，售出 1 * 3 = 3 元。
//- 1 块 1 x 4 的小木块，售出 1 * 2 = 2 元。
//总共售出 14 + 3 + 2 = 19 元。
//19 元是最多能得到的钱数。
//
//
// 示例 2：
//
//
//
//
//输入：m = 4, n = 6, prices = [[3,2,10],[1,4,2],[4,1,3]]
//输出：32
//解释：上图展示了一个可行的方案。包括：
//- 3 块 3 x 2 的小木块，售出 3 * 10 = 30 元。
//- 1 块 1 x 4 的小木块，售出 1 * 2 = 2 元。
//总共售出 30 + 2 = 32 元。
//32 元是最多能得到的钱数。
//注意我们不能旋转 1 x 4 的木块来得到 4 x 1 的木块。
//
//
//
// 提示：
//
//
// 1 <= m, n <= 200
// 1 <= prices.length <= 2 * 10⁴
// prices[i].length == 3
// 1 <= hi <= m
// 1 <= wi <= n
// 1 <= pricei <= 10⁶
// 所有 (hi, wi) 互不相同 。
//
//
// Related Topics 记忆化搜索 数组 动态规划 👍 106 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sellingWood(m int, n int, prices [][]int) int64 {
	//将prices按照宽高建立hash
	value := make(map[[2]int]int)
	for _, price := range prices {
		value[[2]int{price[0], price[1]}] = price[2]
	}
	//dp[i][j]表示i*j的木块的最大价格
	dp := make([][]int64, m+1)
	//初始化dp
	for i := range dp {
		dp[i] = make([]int64, n+1)
		for i2 := range dp[i] {
			dp[i][i2] = -1
		}
	}
	var dfs func(int, int, [][]int64) int64

	dfs = func (m int, n int, dp [][]int64) int64 {
		if dp[m][n]!=-1{
			return dp[m][n]
		}
		res := int64(0)
		val, ok := value[[2]int{m, n}]
		if ok {
			res = int64(val)
		}
		if m>1{
			for i := 1; i < m; i++ {
				res = max(res, dfs(i, n, dp)+dfs(m-i, n, dp))
			}
		}
		if n>1{
			for i := 1; i < n; i++ {
				res = max(res,dfs(m,i,dp)+dfs(m,n-i,dp))
			}
		}
		dp[m][n] = res
		return  res;
	}
	res := dfs(m, n, dp)
	return res
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
