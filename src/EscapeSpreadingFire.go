package main

import "sort"

func main() {

}

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// @Method: maximumMinutes
// @Description: 逃离火灾的最少时间
// @param grid
// @return int
func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := sort.Search(m*n+1, func(t int) bool {
		onFire := make([][]bool, m)
		for i := range onFire {
			onFire[i] = make([]bool, n)
		}
		f := []pair{}
		for i, row := range grid {
			for j, x := range row {
				if x == 1 {
					onFire[i][j] = true // 标记着火的位置
					f = append(f, pair{i, j})
				}
			}
		}
		// 火的 BFS
		spreadFire := func() {
			tmp := f
			f = nil
			for _, p := range tmp {
				for _, d := range dirs { // 枚举上下左右四个方向
					x, y := p.x+d.x, p.y+d.y
					if 0 <= x && x < m && 0 <= y && y < n && !onFire[x][y] && grid[x][y] == 0 {
						onFire[x][y] = true // 标记着火的位置
						f = append(f, pair{x, y})
					}
				}
			}
		}
		for ; t > 0 && len(f) > 0; t-- { // 如果火无法扩散就提前退出
			spreadFire() // 火扩散
		}
		if onFire[0][0] {
			return true // 起点着火，寄
		}

		// 人的 BFS
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[0][0] = true
		q := []pair{{}}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, p := range tmp {
				if onFire[p.x][p.y] { // 人走到这个位置后，火也扩散到了这个位置
					continue
				}
				for _, d := range dirs { // 枚举上下左右四个方向
					x, y := p.x+d.x, p.y+d.y
					if 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && !onFire[x][y] && grid[x][y] == 0 {
						if x == m-1 && y == n-1 {
							return false // 我们安全了…暂时。
						}
						vis[x][y] = true // 避免反复访问同一个位置
						q = append(q, pair{x, y})
					}
				}
			}
			spreadFire() // 火扩散
		}
		return true // 人被火烧到，或者没有可以到达安全屋的路
	}) - 1
	if ans < m*n {
		return ans
	}
	return 1_000_000_000
}
