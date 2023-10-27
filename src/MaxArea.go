package main

import "sort"

func main() {

}
func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func area(a, b int) int {
	return a * b
}
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)

	sort.Ints(verticalCuts)
	i := maxbian(horizontalCuts, h)
	i2 := maxbian(verticalCuts, w)
	return i * i2 % 1000000007
}
func maxbian(bian []int, h int) int {
	left := 0
	res := 0
	for i := 0; i < len(bian); i++ {
		res = maxInt(bian[i]-left, res)
		left = bian[i]
	}
	res = max(res, h-left)
	return res
}
