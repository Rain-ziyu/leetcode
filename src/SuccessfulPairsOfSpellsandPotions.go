package main

import (
	"fmt"
	"sort"
)

func main() {
	pairs := successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7)
	//打印排序后的spells
	for i := 0; i < len(pairs); i++ {
		fmt.Printf(" %d", pairs[i])
	}

}

// @Method: successfulPairs
// @Description: 成功的法术和药水对数
// 给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[j] 表示第 j 瓶药水的能量强度。
//
// 同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合。
//
// 请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目。
// @param spells
// @param potions
// @param success
// @return []int
func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))
	sort.Ints(potions)
	//打印排序后的spells
	for i := 0; i < len(potions); i++ {
		fmt.Printf(" %d", potions[i])
	}
	idx := make([]int, len(spells))
	for i, _ := range idx {
		idx[i] = i
	}
	//按照spells的值对idx进行排序 idx内的值就是spells的下标
	sort.Slice(idx, func(i, j int) bool {
		return spells[idx[i]] < spells[idx[j]]
	})

	fmt.Println()
	//打印排序后的idx
	for i := 0; i < len(idx); i++ {
		fmt.Printf(" %d", idx[i])
	}
	fmt.Println()
	r := 0
	minspell := len(potions) - 1
	for i := 0; i < len(idx); i++ {
		for j := minspell; j >= 0 && int64(spells[idx[i]])*int64(potions[j]) >= success; j-- {
			r++
			minspell = j - 1
		}
		res[idx[i]] = r
	}
	return res
}
