package main

import "math"

func main() {

}
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	ans := []int{math.MaxInt32 / 2, -1}
	mp := make([][]int, n)
	dis := make([][]int, n)
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
		dis[i] = make([]int, n)
		vis[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			mp[i][j] = math.MaxInt32 / 2
			dis[i][j] = math.MaxInt32 / 2
			vis[i][j] = false
		}
	}

	for _, eg := range edges {
		from, to, weight := eg[0], eg[1], eg[2]
		mp[from][to], mp[to][from] = weight, weight
	}
	//从0开始枚举每一个节点道达其他节点的最短路径
	for i := 0; i < n; i++ {

		dis[i][i] = 0
		//dijkstra算法，用于计算一个固定的源节点i 到达其他所有结点的最短路径
		//根据已知的距离列表 选择  不在已知距离列表中 且 距离源节点最近的节点。
		//原理是 每次从 已被确认了从源节点到目标节点最短路径的值的那些节点 集合中找到与这些节点相邻的未被确认节点 找出其中距离源节点最短的那个 加入已知的距离列表
		for j := 0; j < n; j++ {
			//初始化最短距离为-1表示目前没有任何的已知距离列表
			t := -1
			//遍历所有不在距离列表中的节点 找到哪个目前离源节点最近
			for k := 0; k < n; k++ {
				if !vis[i][k] && (t == -1 || dis[i][k] < dis[i][t]) {
					t = k
				}
			}
			//根据上边找到的新的距离源节点最近的那个节点 使用他的距离尝试去更新其他的节点距离源节点的距离
			//其实这里会被更新的都是未被放到最短距离中的节点
			for k := 0; k < n; k++ {
				dis[i][k] = min(dis[i][k], dis[i][t]+mp[t][k])
			}
			vis[i][t] = true
		}
	}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if dis[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= ans[0] {
			ans[0] = cnt
			ans[1] = i
		}
	}
	return ans[1]
}
