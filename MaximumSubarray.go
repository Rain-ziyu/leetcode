package main

func main() {

}
func maxSubArray1(nums []int) int {
	pre := 0
	res := nums[0]
	for _, num := range nums {
		if pre < 0 {
			pre = num
		} else {
			pre = pre + num

		}
		if pre > res {
			res = pre
		}
	}
	return res
}

// 线段树
func maxSubArray(nums []int) int {
	sum := get(0, len(nums)-1, nums).mSum
	return sum
}

//lSum 表示 [l,r]内以 l 为左端点的最大子段和
//rSum {rSum}rSum 表示 [l,r] 内以 r 为右端点的最大子段和
//mSum   表示 [l,r] 内的最大子段和
//iSum  iSum 表示 [l,r] 的区间和
//首先最好维护的是 iSum   iSum，区间    [l,r] 的 iSum   iSum 就等于「左子区间」的 iSum   iSum 加上「右子区间」的 iSum   iSum。
//对于    [l,r] 的 lSum  {lSum}lSum，存在两种可能，它要么等于「左子区间」的 lSum  {lSum}lSum，要么等于「左子区间」的 iSum   iSum 加上「右子区间」的 lSum  {lSum}lSum，二者取大。
//对于    [l,r] 的 rSum  {rSum}rSum，同理，它要么等于「右子区间」的 rSum  {rSum}rSum，要么等于「右子区间」的 iSum   iSum 加上「左子区间」的 rSum  {rSum}rSum，二者取大。
//当计算好上面的三个量之后，就很好计算    [l,r] 的 mSum    了。
//我们可以考虑    [l,r] 的 mSum    对应的区间是否跨越 m——它可能不跨越 m，
//也就是说    [l,r] 的 mSum    可能是「左子区间」的 mSum    和 「右子区间」的 mSum    中的一个；
//它也可能跨越 m，可能是「左子区间」的 rSum  {rSum}rSum 和 「右子区间」的 lSum  {lSum}lSum 求和。
//三者取大。
//

type Status struct {
	iSum int
	lSum int
	rSum int
	mSum int
}

func get(start int, end int, nums []int) Status {
	if start == end {
		return Status{iSum: nums[start], rSum: nums[start], lSum: nums[start], mSum: nums[start]}
	}
	zhongjian := (start + end) / 2
	left := get(start, zhongjian, nums)
	right := get(zhongjian+1, end, nums)

	res := Status{iSum: left.iSum + right.iSum,
		lSum: max(left.iSum+right.lSum, left.lSum),
		rSum: max(left.rSum+right.iSum, right.rSum),
		mSum: max(max(left.rSum+right.lSum, left.mSum), right.mSum)}
	return res
}
