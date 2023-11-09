package main

func main() {

}

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 返回三个数，分别表示到达安全屋/安全屋左边/安全屋上边的最短时间
	bfs := func(q []pair) (int, int, int) {
		time := make([][]int, m)
		for i := range time {
			time[i] = make([]int, n)
			for j := range time[i] {
				time[i][j] = -1 // -1 表示未访问
			}
		}
		for _, p := range q {
			time[p.x][p.y] = 0
		}
		for t := 1; len(q) > 0; t++ { // 每次循环向外扩展一圈
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dirs { // 枚举上下左右四个方向
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 0 && time[x][y] < 0 {
						time[x][y] = t
						q = append(q, pair{x, y})
					}
				}
			}
		}
		return time[m-1][n-1], time[m-1][n-2], time[m-2][n-1]
	}
	//先bfs遍历仍能否到达安全屋
	manToHouseTime, m1, m2 := bfs([]pair{{}})
	if manToHouseTime < 0 { // 人无法到安全屋
		return -1
	}

	firePos := []pair{}
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				firePos = append(firePos, pair{i, j})
			}
		}
	}
	fireToHouseTime, f1, f2 := bfs(firePos) // 多个着火点同时跑 BFS
	if fireToHouseTime < 0 {                // 火无法到安全屋
		return 1_000_000_000
	}

	d := fireToHouseTime - manToHouseTime
	if d < 0 { // 火比人先到安全屋
		return -1
	}

	if m1 != -1 && m1+d < f1 || // 安全屋左边相邻格子，人比火先到
		m2 != -1 && m2+d < f2 { // 安全屋上边相邻格子，人比火先到
		return d // 图中第一种情况
	}
	return d - 1 // 图中第二种情况
}
