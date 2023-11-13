package main

func main() {

}

type NumArray struct {
	nums []int
	tree []int
}

func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	for i, x := range nums {
		i++
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
