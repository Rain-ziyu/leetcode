// 给你区间的 空 集，请你设计并实现满足要求的数据结构：
//
// 新增：添加一个区间到这个区间集合中。
// 统计：计算出现在 至少一个 区间中的整数个数。
//
// 实现 CountIntervals 类：
//
// CountIntervals() 使用区间的空集初始化对象
// void add(int left, int right) 添加区间 [left, right] 到区间集合之中。
// int count() 返回出现在 至少一个 区间中的整数个数。
//
// 注意：区间 [left, right] 表示满足 left <= x <= right 的所有整数 x 。
//
// 示例 1：
//
// 输入
// ["CountIntervals", "add", "add", "count", "add", "count"]
// [[], [2, 3], [7, 10], [], [5, 8], []]
// 输出
// [null, null, null, 6, null, 8]
//
// 解释
// CountIntervals countIntervals = new CountIntervals(); // 用一个区间空集初始化对象
// countIntervals.add(2, 3);  // 将 [2, 3] 添加到区间集合中
// countIntervals.add(7, 10); // 将 [7, 10] 添加到区间集合中
// countIntervals.count();    // 返回 6
//
//	// 整数 2 和 3 出现在区间 [2, 3] 中
//	// 整数 7、8、9、10 出现在区间 [7, 10] 中
//
// countIntervals.add(5, 8);  // 将 [5, 8] 添加到区间集合中
// countIntervals.count();    // 返回 8
//
//	// 整数 2 和 3 出现在区间 [2, 3] 中
//	// 整数 5 和 6 出现在区间 [5, 8] 中
//	// 整数 7 和 8 出现在区间 [5, 8] 和区间 [7, 10] 中
//	// 整数 9 和 10 出现在区间 [7, 10] 中
//
// 提示：
//
// 1 <= left <= right <= 10⁹
// 最多调用 add 和 count 方法 总计 10⁵ 次
// 调用 count 方法至少一次
//
// Related Topics 设计 线段树 有序集合 👍 59 👎 0
package main

import (
	"github.com/emirpasic/gods/maps/treemap"
)

// leetcode submit region begin(Prohibit modification and deletion)
type CountIntervals struct {
	*treemap.Map
	cnt int
}

func Constructor() CountIntervals {
	return CountIntervals{
		treemap.NewWithIntComparator(), 0,
	}
}

func (this *CountIntervals) Add(left int, right int) {
	//循环查找已有的区间中left小于（传入）right（且（传入）left小于已有区间中的right）
	//那么此时说明这些区间存在重复的区间  那么合并两个区间 并更新 left与right
	//继续循环合并 中间计数 cnt 减去重复区间循环完之后再加上最终合并完成的区间大小
	for k, v := this.Floor(right); k != nil && v.(int) >= left; k, v = this.Floor(right) {
		l, r := k.(int), v.(int)
		left, right = min(left, l), max(right, r)
		this.cnt -= r - l + 1
		this.Remove(k)
	}
	this.cnt += right - left + 1
	//将map中key存为left value为right
	this.Put(left, right)
}

func (this *CountIntervals) Count() int {
	return this.cnt
}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */
//leetcode submit region end(Prohibit modification and deletion)
