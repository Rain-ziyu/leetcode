package main

func main() {
	println(findTheLongestBalancedSubstring("01011"))
}

// @Method: findTheLongestBalancedSubstring
// @Description: 最长平衡子字符串
// @param s
// @return int
func findTheLongestBalancedSubstring(s string) int {
	dp := make([][]int, len(s))
	res := 0
	//	初始化dp数组
	for i := 0; i < len(s); i++ {
		// 其实对于dp数组来说只需要对角线下面的一半即可
		dp[i] = make([]int, i)
		//	初始化dp数组，其中满足字符串中[i][i-1]字符相等时，dp[i][i-1] = 1
		if i > 0 && s[i] == '1' && s[i-1] == '0' {
			dp[i][i-1] = 2
			res = 2
		}
	}

	//	横向遍历dp数组 状态转移方程为 f(m,n) = f(m-1,n+1) + 1 且 m-n > 1 且 m-n+1一定是偶数
	for i := 2; i < len(s); i++ {
		j := 0
		if i%2 == 0 {
			j = 1
		}
		for ; j < i-1; j = j + 2 {
			if s[i] == '1' && s[j] == '0' && (j-i+1)%2 == 0 && dp[i-1][j+1] > 0 {
				dp[i][j] = dp[i-1][j+1] + 2
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

// @Method: findTheLongestBalancedSubstring
// @Description: 最长平衡子字符串 优化不使用
// @param s
// @return int
func findTheLongestBalancedSubstringFast(s string) int {
	res := 0
	length := len(s)

	for index := 0; index < length; {
		lingbiaoshi := 0
		yibiaoshi := 0
		for s[index] == '0' && index < length {
			lingbiaoshi++
			index++
		}
		for s[index] == '1' && index < length {
			yibiaoshi++
			index++
		}
		res = 2 * min(lingbiaoshi, yibiaoshi)
	}
	return res
}
