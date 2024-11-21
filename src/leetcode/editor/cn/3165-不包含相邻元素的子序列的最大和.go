//给你一个整数数组 nums 和一个二维数组 queries，其中 queries[i] = [posi, xi]。
//
// 对于每个查询 i，首先将 nums[posi] 设置为 xi，然后计算查询 i 的答案，该答案为 nums 中 不包含相邻元素 的 子序列 的 最大 和。
//
//
// 返回所有查询的答案之和。
//
// 由于最终答案可能非常大，返回其对 10⁹ + 7 取余 的结果。
//
// 子序列 是指从另一个数组中删除一些或不删除元素而不改变剩余元素顺序得到的数组。
//
//
//
// 示例 1：
//
//
// 输入：nums = [3,5,9], queries = [[1,-2],[0,-3]]
//
//
// 输出：21
//
// 解释： 执行第 1 个查询后，nums = [3,-2,9]，不包含相邻元素的子序列的最大和为 3 + 9 = 12。 执行第 2 个查询后，nums =
// [-3,-2,9]，不包含相邻元素的子序列的最大和为 9 。
//
// 示例 2：
//
//
// 输入：nums = [0,-1], queries = [[0,-5]]
//
//
// 输出：0
//
// 解释： 执行第 1 个查询后，nums = [-5,-1]，不包含相邻元素的子序列的最大和为 0（选择空子序列）。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -10⁵ <= nums[i] <= 10⁵
// 1 <= queries.length <= 5 * 10⁴
// queries[i] == [posi, xi]
// 0 <= posi <= nums.length - 1
// -10⁵ <= xi <= 10⁵
//
//
// Related Topics 线段树 数组 分治 动态规划 👍 37 👎 0

package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
const MOD = 1000000007

func maximumSumSubsequence(nums []int, queries [][]int) int {
	n := len(nums)
	tree := NewSegTree(n)
	tree.Init(nums)

	ans := int64(0)
	for _, q := range queries {
		tree.Update(q[0], q[1])
		ans = (ans + tree.Query()) % MOD
	}
	return int(ans)
}

type SegNode struct {
	v00, v01, v10, v11 int64
}

func NewSegNode() *SegNode {
	return &SegNode{0, 0, 0, 0}
}

func (sn *SegNode) Set(v int64) {
	sn.v00, sn.v01, sn.v10 = 0, 0, 0
	sn.v11 = int64(math.Max(float64(v), 0))
}

func (sn *SegNode) Best() int64 {
	return sn.v11
}

type SegTree struct {
	n    int
	tree []*SegNode
}

func NewSegTree(n int) *SegTree {
	tree := make([]*SegNode, n*4+1)
	for i := range tree {
		tree[i] = NewSegNode()
	}
	return &SegTree{n, tree}
}

func (st *SegTree) Init(nums []int) {
	st.internalInit(nums, 1, 1, st.n)
}

func (st *SegTree) Update(x, v int) {
	st.internalUpdate(1, 1, st.n, x+1, int64(v))
}

func (st *SegTree) Query() int64 {
	return st.tree[1].Best()
}

func (st *SegTree) internalInit(nums []int, x, l, r int) {
	if l == r {
		st.tree[x].Set(int64(nums[l-1]))
		return
	}
	mid := (l + r) / 2
	st.internalInit(nums, x*2, l, mid)
	st.internalInit(nums, x*2+1, mid+1, r)
	st.pushup(x)
}

func (st *SegTree) internalUpdate(x, l, r int, pos int, v int64) {
	if l > pos || r < pos {
		return
	}
	if l == r {
		st.tree[x].Set(v)
		return
	}
	mid := (l + r) / 2
	st.internalUpdate(x*2, l, mid, pos, v)
	st.internalUpdate(x*2+1, mid+1, r, pos, v)
	st.pushup(x)
}

func (st *SegTree) pushup(x int) {
	l, r := x*2, x*2+1
	st.tree[x].v00 = max(st.tree[l].v00+st.tree[r].v10, st.tree[l].v01+st.tree[r].v00)
	st.tree[x].v01 = max(st.tree[l].v00+st.tree[r].v11, st.tree[l].v01+st.tree[r].v01)
	st.tree[x].v10 = max(st.tree[l].v10+st.tree[r].v10, st.tree[l].v11+st.tree[r].v00)
	st.tree[x].v11 = max(st.tree[l].v10+st.tree[r].v11, st.tree[l].v11+st.tree[r].v01)
}

//leetcode submit region end(Prohibit modification and deletion)
