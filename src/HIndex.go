package main

import "fmt"

func main() {
	citations := []int{0}
	index := hIndex(citations)
	fmt.Println(index)
}
func hIndex(citations []int) int {
	//	为了达到log n 使用二分查找找到目标值
	//一名科研人员的 h 指数是指他（她）的 （n 篇论文中）总共有 h 篇论文分别被引用了至少 h 次。
	n := len(citations)
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if n-mid > citations[mid] {
			left = mid + 1
		} else if n-mid < citations[mid] {
			right = mid - 1
		} else {
			//n-mid == citations[mid]  时 即 h = citations[mid]
			return citations[mid]
		}
	}
	return n - left
}
