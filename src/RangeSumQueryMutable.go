package main

func main() {
	values := []int{1, 3, 5}
	Constructor(values)
}

type NumArray struct {
	nums []int
	tree []int
}

func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	//对于这种循环来说 i的值每次都是递增的 不受i++影响
	for i, x := range nums {
		i = i + 1
		tree[i] += x
		if nxt := i + i&-i; nxt < len(tree) {
			tree[nxt] += tree[i]
		}
	}
	return NumArray{nums, tree}
}

func (a *NumArray) Update(index, val int) {
	delta := val - a.nums[index]
	a.nums[index] = val
	//更新该位置修改所影响的所有的关键区间
	for i := index + 1; i < len(a.tree); i += i & -i {
		a.tree[i] += delta
	}
}

func (a *NumArray) prefixSum(i int) (s int) {
	for ; i > 0; i &= i - 1 { // i -= i & -i 的另一种写法
		s += a.tree[i]
	}
	return
}

func (a *NumArray) SumRange(left, right int) int {
	return a.prefixSum(right+1) - a.prefixSum(left)
}
