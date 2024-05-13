//在给定的 m x n 网格
// grid 中，每个单元格可以有以下三个值之一： 
//
// 
// 值 0 代表空单元格； 
// 值 1 代表新鲜橘子； 
// 值 2 代表腐烂的橘子。 
// 
//
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。 
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
//输出：4
// 
//
// 示例 2： 
//
// 
//输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
//输出：-1
//解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
// 
//
// 示例 3： 
//
// 
//输入：grid = [[0,2]]
//输出：0
//解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 10 
// grid[i][j] 仅为 0、1 或 2 
// 
//
// Related Topics 广度优先搜索 数组 矩阵 👍 869 👎 0

package main

import "math"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func orangesRotting(grid [][]int) int {
	//	用于记录队列里剩余值
	queue := make([][]int, 0)
	time := make([][]int, len(grid))
	for i := range grid {
		time[i] = make([]int, len(grid[i]))
		for i2 := range grid[i] {
			time[i][i2] = math.MaxInt
			if grid[i][i2] == 2 {
				time[i][i2] = 1
				queue = append(queue, []int{i, i2})
			}
		}
	}
	fangxiang := [][]int {{0,1},{0,-1},{1,0},{-1,0},}
	res := 1
	for len(queue)!=0 {
		res++
		tmpQueue := make([][]int, 0)
		for i := range queue {
			for i1 := range fangxiang {
				x := queue[i][0] + fangxiang[i1][0]
				y := queue[i][1] + fangxiang[i1][1]
				if x>=0&&x<len(grid)&&y>=0&&y<len(grid[x])&&grid[x][y] == 1 {
					if time[x][y] == math.MaxInt {
						tmpQueue = append(tmpQueue, []int{x,y})
						time[x][y] = res
					}
				}
			}
		}
		queue = tmpQueue
	}
	max:=1
	for i := range grid {
		for i2 := range grid[i] {
			if (grid[i][i2] == 1||grid[i][i2] == 2)&& time[i][i2] == math.MaxInt{
				return  -1
			}
			if (grid[i][i2] == 1||grid[i][i2] == 2)&& time[i][i2] != math.MaxInt{
				if time[i][i2]>max{
					max = time[i][i2]
				}
			}
		}
	}
return  max-1
}

//leetcode submit region end(Prohibit modification and deletion)
