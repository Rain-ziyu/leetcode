package main

import "fmt"

func main() {
	fmt.Printf("result:%d \n", minCount([]int{4, 2, 1}))
}

/*
*
桌上有 n 堆力扣币，每堆的数量保存在数组 coins 中。我们每次可以选择任意一堆，拿走其中的一枚或者两枚，求拿完所有力扣币的最少次数。
*/
func minCount(coins []int) int {
	var result int = 0
	for _, coin := range coins {
		result += coin / 2
		fmt.Println(coin)
		if coin%2 != 0 {
			result += 1
		}
	}
	return result
}
