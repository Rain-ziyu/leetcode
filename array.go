package main

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
)

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			return false
		}
	}
	return true
}
func minimumValueSum(nums []int, andValues []int) int {
	const INF = (1 << 20) - 1
	n := len(nums)
	m := len(andValues)
	memo := make([]map[int]int, n*m)
	for i := range memo {
		memo[i] = make(map[int]int)
	}

	var dfs func(i, j, cur int) int
	dfs = func(i, j, cur int) int {
		key := i*m + j
		if i == n && j == m {
			return 0
		}
		if i == n || j == m {
			return INF
		}
		if val, ok := memo[key][cur]; ok {
			return val
		}

		cur &= nums[i]
		if cur&andValues[j] < andValues[j] {
			return INF
		}

		res := dfs(i+1, j, cur)
		if cur == andValues[j] {
			res = min(res, dfs(i+1, j+1, INF)+nums[i])
		}
		memo[key][cur] = res
		return res
	}

	res := dfs(0, 0, INF)
	if res < INF {
		return res
	}
	return -1
}

func waysToReachStair(k int) int {
	n, npow, ans := 0, 1, 0
	for {
		if npow-n-1 <= k && k <= npow {
			ans += comb(n+1, npow-k)
		} else if npow-n-1 > k {
			break
		}
		//枚举操作二的次数
		n++
		npow *= 2
	}
	return ans
}

func comb(n, k int) int {
	ans := 1
	for i := n; i >= n-k+1; i-- {
		ans *= i
		ans /= n - i + 1
	}
	return ans
}

func findMaximumNumber(k int64, x int) int64 {
	l, r := int64(1), (k+1)<<x
	for l < r {
		m := (l + r + 1) / 2
		if accumulatedPrice(x, m) > k {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func accumulatedBitPrice(x int, num int64) int64 {
	period := int64(1) << x
	res := period / 2 * (num / period)
	if num%period >= period/2 {
		res += num%period - (period/2 - 1)
	}
	return res
}

func accumulatedPrice(x int, num int64) int64 {
	res := int64(0)
	length := 64 - bits.LeadingZeros64(uint64(num))
	for i := x; i <= length; i += x {
		res += accumulatedBitPrice(i, num)
	}
	return res
}

//func findMaximumNumber(k int64, x int) int64 {
//	ans := sort.Search(int(k+1)<<x, func(num int) bool {
//		num++
//		n := bits.Len(uint(num))
//		memo := make([][]int, n)
//		for i := range memo {
//			memo[i] = make([]int, n+1)
//			for j := range memo[i] {
//				memo[i][j] = -1
//			}
//		}
//		var dfs func(int, int, bool) int
//		dfs = func(i, cnt1 int, limitHigh bool) (res int) {
//			if i < 0 {
//				return cnt1
//			}
//			if !limitHigh {
//				p := &memo[i][cnt1]
//				if *p >= 0 {
//					return *p
//				}
//				defer func() { *p = res }()
//			}
//			up := 1
//			if limitHigh {
//				up = num >> i & 1
//			}
//			for d := 0; d <= up; d++ {
//				c := cnt1
//				if d == 1 && (i+1)%x == 0 {
//					c++
//				}
//				res += dfs(i-1, c, limitHigh && d == up)
//			}
//			return
//		}
//		return dfs(n-1, 0, true) > int(k)
//	})
//	return int64(ans)
//}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	var dfs func(int) int
	m := make(map[int]*Employee)
	for _, e := range employees {
		m[e.Id] = e
	}
	dfs = func(id int) int {
		e := m[id]
		res := e.Importance
		for _, subId := range e.Subordinates {
			res += dfs(subId)
		}
		return res
	}

	return dfs(id)
}

func minimumSubstringsInPartition(s string) int {

	dp := make([]int, len(s)+1)
	dp[0] = 1
	cache := make(map[string]int)
	isPalindrome := func(s string) bool {
		if cache[s] != 0 {
			if cache[s] == 1 {
				return true
			} else {
				return false
			}
		}
		dp1 := make(map[uint8]int, 26)
		for i := range s {
			dp1[s[i]-'a']++
		}
		for i := range dp1 {
			if dp1[uint8(i)] != 0 && dp1[uint8(i)] != dp1[s[0]-'a'] {
				cache[s] = 2
				return false
			}
		}

		cache[s] = 1
		return true
	}

	for i := 1; i < len(s)+1; i++ {
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if isPalindrome(s[j:i]) {
				if j != 0 {
					dp[i] = min(dp[i], dp[j]+1)
				} else {
					dp[i] = min(dp[i], 1)
				}
			}
		}
	}
	return dp[len(s)]
}

func canMakeSquare(grid [][]byte) bool {
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			res := 0
			if grid[i][j] == 'B' {
				res++
			}
			if grid[i][j+1] == 'B' {
				res++
			}
			if grid[i+1][j] == 'B' {
				res++
			}
			if grid[i+1][j+1] == 'B' {
				res++
			}
			if res != 2 {
				return true
			}
		}
	}
	return false
}

func maxConsecutiveChar(answerKey string, k int, ch byte) (ans int) {
	left, sum := 0, 0
	for right := range answerKey {
		if answerKey[right] != ch {
			sum++
		}
		for sum > k {
			if answerKey[left] != ch {
				sum--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveChar(answerKey, k, 'T'),
		maxConsecutiveChar(answerKey, k, 'F'))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {

	palindrome := countWays([]int{6, 0, 3, 3, 6, 7, 2, 7})
	fmt.Println(palindrome)
}
func maxStrength(nums []int) int64 {
	negativeCount, zeroCount, positiveCount := 0, 0, 0
	prod := 1
	maxNegative := math.MinInt32

	for _, num := range nums {
		if num < 0 {
			negativeCount++
			prod *= num
			if num > maxNegative {
				maxNegative = num
			}
		} else if num == 0 {
			zeroCount++
		} else {
			prod *= num
			positiveCount++
		}
	}

	if negativeCount == 1 && zeroCount == 0 && positiveCount == 0 {
		return int64(nums[0])
	}
	if negativeCount <= 1 && positiveCount == 0 {
		return int64(0)
	}
	if prod < 0 {
		return int64(prod / maxNegative)
	} else {
		return int64(prod)
	}
}
func countWays(nums []int) int {
	sort.Ints(nums)
	res := 0
	// i表示当前位置之前都被选中
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < i+1 && nums[i+1] > i+1 {
			res++
		}
	}
	//全都 不被选中
	if nums[0] > 0 {
		res++
	}
	if nums[len(nums)-1] < len(nums) {
		res++
	}
	return res
}
