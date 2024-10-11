package main

import "fmt"

func main() {
	sightseeingPair := maxScoreSightseeingPair([]int{1, 2})
	fmt.Println(sightseeingPair)
}
func maxScoreSightseeingPair(values []int) int {
	values1 := make([]int, len(values))
	copy(values1, values)
	values[len(values)-1] = values[len(values)-1] - len(values) + 1
	for i := len(values) - 2; i >= 0; i-- {
		values[i] = max1(values[i]-i, values[i+1])
	}
	res := 0
	for i := 0; i <= len(values)-2; i++ {
		res = max1(res, values1[i]+i+values[i+1])
	}
	return res
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
