package main

import "fmt"

func main() {
	digits := countDigits(1248)
	fmt.Println(digits)
}

//
//  @Method: countDigits
//  @Description: 统计能整除数字的位数
//  @param num
//  @return int
//

func countDigits(num int) int {
	res := 0
	current := num
	next := 0
	// 从右向左遍历一个整数的所有位
	for current > 0 {
		// 下一次计算能否整除的数
		next = current % 10
		// 下一次求余操作的数
		current = current / 10
		// 判断当前位能否整除num
		if next != 0 && num%next == 0 {
			res++
		}

	}
	return res
}
