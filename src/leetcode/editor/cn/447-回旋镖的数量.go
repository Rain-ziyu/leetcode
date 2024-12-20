//给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。回旋镖 是由点 (i, j, k) 表示的元组 ，其中
// i 和 j 之间的距离和 i 和 k 之间的欧式距离相等（需要考虑元组的顺序）。
//
// 返回平面上所有回旋镖的数量。
//
// 示例 1：
//
//
//输入：points = [[0,0],[1,0],[2,0]]
//输出：2
//解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
//
//
// 示例 2：
//
//
//输入：points = [[1,1],[2,2],[3,3]]
//输出：2
//
//
// 示例 3：
//
//
//输入：points = [[1,1]]
//输出：0
//
//
//
//
// 提示：
//
//
// n == points.length
// 1 <= n <= 500
// points[i].length == 2
// -10⁴ <= xi, yi <= 10⁴
// 所有点都 互不相同
//
//
// Related Topics 数组 哈希表 数学 👍 297 👎 0

package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfBoomerangs(points [][]int) int {
	if len(points) < 3 {
		return 0
	}

	res := 0
	//枚举中心点
	for i := 0; i < len(points); i++ {
		m := make(map[int]int)
		//	循环计算到达中心点的距离
		for j := 0; j < len(points); j++ {
			//	计算距离的平方
			m[(points[i][0]-points[j][0])*(points[i][0]-points[j][0])+(points[i][1]-points[j][1])*(points[i][1]-points[j][1])]++
		}
		fmt.Println(m)
		for _, value := range m {
			res += value * (value - 1)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
