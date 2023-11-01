package main

func main() {

}

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	children := make([][]int, n)
	//将数据重建成 children第i个元素内存储的是i的所有子节点 go语言中数据下标可以是-1
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	var dfs func(int) (map[int]bool, int)
	/**
	 * @Method: max
	 * @Description: 求当前节点的基因数 即最小的不包含在子树及其本身中的最小数
	 */
	dfs = func(node int) (map[int]bool, int) {
		geneSet := map[int]bool{nums[node]: true}
		for _, child := range children[node] {
			// 递归求子节点
			childGeneSet, y := dfs(child)
			// 更新当前节点的基因数
			res[node] = max(res[node], y)
			// 比较两个子节点元素集合大小
			if len(childGeneSet) > len(geneSet) {
				geneSet, childGeneSet = childGeneSet, geneSet
			}
			//，保证小的合并到打的里边
			for gene, _ := range childGeneSet {
				// 合并小的集合到大的集合中
				geneSet[gene] = true
			}
		}
		// 从res[node]开始遍历，无需从node=0开始，因为比node小的节点已经被调用，找到第一个不在geneSet中不为true的数，那么这个数就是当前节点最小的基因数
		for geneSet[res[node]] {
			res[node]++
		}
		return geneSet, res[node]
	}

	dfs(0)
	return res
}
