package main

func main() {

}
func maximumInvitations(favorite []int) int {
	n := len(favorite)
	// 统计入度，便于进行拓扑排序
	indeg := make([]int, n)
	for _, x := range favorite {
		indeg[x]++
	}

	used := make([]bool, n)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = 1
	}

	q := []int{}
	//求没有人喜欢的人
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	//使用没有人喜欢的人作为出发点进行计算之后的每一个节点
	for len(q) > 0 {
		u := q[0]
		//每一个被遍历到的节点都说明该节点不属于环中的一部分
		used[u] = true
		q = q[1:]
		v := favorite[u]
		// 状态转移 f中相当于最终存储的是经过当前节点最长的那个线的长度（不处于环上的节点）
		f[v] = max(f[v], f[u]+1)
		indeg[v]--
		// 直到该节点出度为零之后就增加到append进行计算其喜欢的人
		if indeg[v] == 0 {
			q = append(q, v)
		}
	}
	// ring 表示最大的环的大小
	// total 表示所有环大小为 2 的「基环内向树」上的最长的「双向游走」路径之和
	ring, total := 0, 0
	for i := 0; i < n; i++ {
		if !used[i] {
			j := favorite[i]
			// favorite[favorite[i]] = i 说明环的大小为 2
			if favorite[j] == i {
				total += f[i] + f[j]
				used[i], used[j] = true, true
			} else {
				// 否则环的大小至少为 3，我们需要找出环
				u, cnt := i, 0
				for true {
					cnt++
					u = favorite[u]
					used[u] = true
					if u == i {
						break
					}
				}
				ring = max(ring, cnt)
			}
		}
	}
	return max(ring, total)
}
