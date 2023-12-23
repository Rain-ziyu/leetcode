//给你一个整数数组 piles ，数组 下标从 0 开始 ，其中 piles[i] 表示第 i 堆石子中的石子数量。另给你一个整数 k ，请你执行下述操作 恰
//好 k 次：
//
//
// 选出任一石子堆 piles[i] ，并从中 移除 floor(piles[i] / 2) 颗石子。
//
//
// 注意：你可以对 同一堆 石子多次执行此操作。
//
// 返回执行 k 次操作后，剩下石子的 最小 总数。
//
// floor(x) 为 小于 或 等于 x 的 最大 整数。（即，对 x 向下取整）。
//
//
//
// 示例 1：
//
//
//输入：piles = [5,4,9], k = 2
//输出：12
//解释：可能的执行情景如下：
//- 对第 2 堆石子执行移除操作，石子分布情况变成 [5,4,5] 。
//- 对第 0 堆石子执行移除操作，石子分布情况变成 [3,4,5] 。
//剩下石子的总数为 12 。
//
//
// 示例 2：
//
//
//输入：piles = [4,3,6,7], k = 3
//输出：12
//解释：可能的执行情景如下：
//- 对第 2 堆石子执行移除操作，石子分布情况变成 [4,3,3,7] 。
//- 对第 3 堆石子执行移除操作，石子分布情况变成 [4,3,3,4] 。
//- 对第 0 堆石子执行移除操作，石子分布情况变成 [2,3,3,4] 。
//剩下石子的总数为 12 。
//
//
//
//
// 提示：
//
//
// 1 <= piles.length <= 10⁵
// 1 <= piles[i] <= 10⁴
// 1 <= k <= 10⁵
//
//
// Related Topics 贪心 数组 堆（优先队列） 👍 35 👎 0

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {

}


// leetcode submit region begin(Prohibit modification and deletion)
func minStoneSum(piles []int, k int) int {
	// 1. 将数组转换为大根堆
	// 2. 取出堆顶元素，计算减去后的值，再放入堆中
	// 3. 重复2，直到k次
	// 4. 计算剩余元素的和
	// 1.
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range piles {
		heap.Push(h, v)
	}
	// 2.
	for i := 0; i < k; i++ {
		// 2.1
		top := heap.Pop(h).(int)
		fmt.Println(top)
		// 2.2
		top -= top / 2
		// 2.3
		heap.Push(h, top)

	}
	// 3.
	res := 0
	for _, v := range *h {
		res += v
	}
	return res
}
type IntHeap sort.IntSlice

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
//leetcode submit region end(Prohibit modification and deletion)
