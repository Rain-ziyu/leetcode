package main

import "fmt"

func main() {
	var array = []int{2, 7, 9, 3, 1}
	fmt.Println(minCapability(array, 2))
}

/*
沿街有一排连续的房屋。每间房屋内都藏有一定的现金。现在有一位小偷计划从这些房屋中窃取现金。

由于相邻的房屋装有相互连通的防盗系统，所以小偷 不会窃取相邻的房屋 。

小偷的 窃取能力 定义为他在窃取过程中能从单间房屋中窃取的 最大金额 。

给你一个整数数组 nums 表示每间房屋存放的现金金额。形式上，从左起第 i 间房屋中放有 nums[i] 美元。

另给你一个整数 k ，表示窃贼将会窃取的 最少 房屋数。小偷总能窃取至少 k 间房屋。

返回小偷的 最小 窃取能力。
*/
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCapability(nums []int, k int) int {
	lower, upper := nums[0], nums[0]
	for _, x := range nums {
		lower = min(lower, x)
		upper = max(upper, x)
	}

	for lower <= upper {
		middle := (lower + upper) / 2
		count, visited := 0, false
		for _, x := range nums {
			if x <= middle && !visited {
				count, visited = count+1, true
			} else {
				visited = false
			}
		}
		if count >= k {
			upper = middle - 1
		} else {
			lower = middle + 1
		}
	}
	return lower
}
