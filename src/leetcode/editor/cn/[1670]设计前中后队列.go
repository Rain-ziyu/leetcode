// 请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。
//
// 请你完成 FrontMiddleBack 类：
//
// FrontMiddleBack() 初始化队列。
// void pushFront(int val) 将 val 添加到队列的 最前面 。
// void pushMiddle(int val) 将 val 添加到队列的 正中间 。
// void pushBack(int val) 将 val 添加到队里的 最后面 。
// int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
// int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
// int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
//
// 请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：
//
// 将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
// 从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。
//
// 示例 1：
//
// 输入：
// ["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle",
// "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
// [[], [1], [2], [3], [4], [], [], [], [], []]
// 输出：
// [null, null, null, null, null, 1, 3, 4, 2, -1]
//
// 解释：
// FrontMiddleBackQueue q = new FrontMiddleBackQueue();
// q.pushFront(1);   // [1]
// q.pushBack(2);    // [1, 2]
// q.pushMiddle(3);  // [1, 3, 2]
// q.pushMiddle(4);  // [1, 4, 3, 2]
// q.popFront();     // 返回 1 -> [4, 3, 2]
// q.popMiddle();    // 返回 3 -> [4, 2]
// q.popMiddle();    // 返回 4 -> [2]
// q.popBack();      // 返回 2 -> []
// q.popFront();     // 返回 -1 -> [] （队列为空）
//
// 提示：
//
// 1 <= val <= 10⁹
// 最多调用 1000 次 pushFront， pushMiddle， pushBack， popFront， popMiddle 和 popBack 。
//
// Related Topics 设计 队列 数组 链表 数据流 👍 83 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
type FrontMiddleBackQueue struct {
	head []int
	tail []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{head: []int{}, tail: []int{}}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.head = append([]int{val},this.head... )
	if len(this.head)>len(this.tail) {
		i := this.head[len(this.head)-1]
		this.head = this.head[:len(this.head)-1]
		this.tail = append(this.tail, i)
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if len(this.head) < len(this.tail) {
		this.head = append(this.head, val)
	} else {
		this.tail = append(this.tail, val)
	}
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.tail = append([]int{val}, this.tail...)
	if len(this.tail)-len(this.head) >= 2 {
		i := this.tail[len(this.tail)-1]
		this.tail = this.tail[:len(this.tail)-1]
		this.head = append(this.head, i)
	}
}

func (this *FrontMiddleBackQueue) PopFront() int {
	res := 0
	if len(this.head) > 0 {
		res = this.head[0]
		this.head = this.head[1:]
		if len(this.tail)-len(this.head) >= 2 {
			i := this.tail[len(this.tail)-1]
			this.tail = this.tail[:len(this.tail)-1]
			this.head = append(this.head, i)
		}
	} else if len(this.tail) > 0 {
		res = this.tail[len(this.tail)-1]
		this.tail = this.tail[:len(this.tail)-1]
	} else {
		res = -1
	}
	return res
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.head) == 0 && len(this.tail) == 0 {
		return -1
	}
	res := 0
	if len(this.head)<len(this.tail){
		res = this.tail[len(this.tail)-1]
		this.tail = this.tail[:len(this.tail)-1]
		return res
	}else {
		res = this.head[len(this.head)-1]
		this.head = this.head[:len(this.head)-1]
		return res
	}
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if   len(this.tail) == 0 {
		return -1
	}
	res := this.tail[0]
	this.tail = this.tail[1:]
	if len(this.head)> len(this.tail){
		i := this.head[len(this.head)-1]
		this.head = this.head[:len(this.head)-1]
		this.tail = append(this.tail, i)
	}
	return res
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
//leetcode submit region end(Prohibit modification and deletion)
