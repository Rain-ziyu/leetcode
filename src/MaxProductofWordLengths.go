package main

import "fmt"

func main() {
	tmp := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	product := maxProduct(tmp)
	fmt.Println(product)
}

/*
*
最大长度的不重复字符串的乘积
*/
func maxProduct(words []string) int {
	resmap := make(map[int]int, len(words))
	for _, word := range words {
		res := 0
		for _, i2 := range word {
			res = res | (1 << (i2 - 'a'))
		}
		resmap[res] = max(resmap[res], len(word))
	}
	res := 0
	for key, value := range resmap {
		for key1, value1 := range resmap {
			if key&key1 == 0 {
				res = max(value*value1, res)
			}
		}
	}
	return res
}
