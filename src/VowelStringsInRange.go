package main

func main() {
	println(vowelStrings([]string{"are", "amy", "u"}, 1, 2))
}

// @Method: vowelStrings
// @Description: 统计范围内的元音字符串数
// @param words
// @param left
// @param right
// @return int
func vowelStrings(words []string, left int, right int) int {
	yuanyin := make(map[uint8]bool, 5)
	yuanyin['a'] = true
	yuanyin['e'] = true
	yuanyin['i'] = true
	yuanyin['o'] = true
	yuanyin['u'] = true
	res := 0
	for i := left; i <= right; i++ {
		// 判断words的首字母是否包含于yuanyin数组中
		if yuanyin[words[i][0]] && yuanyin[words[i][len(words[i])-1]] {
			res++
		}
	}
	return res
}
