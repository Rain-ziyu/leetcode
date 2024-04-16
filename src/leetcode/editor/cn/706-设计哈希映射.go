//不使用任何内建的哈希表库设计一个哈希映射（HashMap）。
//
// 实现 MyHashMap 类：
//
//
// MyHashMap() 用空映射初始化对象
// void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，
//则更新其对应的值 value 。
// int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
// void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。
//
//
//
//
// 示例：
//
//
//输入：
//["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
//[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
//输出：
//[null, null, null, 1, -1, null, 1, null, -1]
//
//解释：
//MyHashMap myHashMap = new MyHashMap();
//myHashMap.put(1, 1); // myHashMap 现在为 [[1,1]]
//myHashMap.put(2, 2); // myHashMap 现在为 [[1,1], [2,2]]
//myHashMap.get(1);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
//myHashMap.get(3);    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
//myHashMap.put(2, 1); // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
//myHashMap.get(2);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
//myHashMap.remove(2); // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
//myHashMap.get(2);    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]
//
//
//
//
// 提示：
//
//
// 0 <= key, value <= 10⁶
// 最多调用 10⁴ 次 put、get 和 remove 方法
//
//
// Related Topics 设计 数组 哈希表 链表 哈希函数 👍 416 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
type MyHashMap struct {
	val [][]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		val: make([][]int, 769),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	res := false
	for i := 0; i < len(this.val[key%769]); i = i + 2 {
		if this.val[key%769][i] == key {
			this.val[key%769][i+1] = value
			res = true
			break
		}
	}
	if !res {
		this.val[key%769] = append(this.val[key%769], []int{key, value}...)
	}
}

func (this *MyHashMap) Get(key int) int {
	for i := 0; i < len(this.val[key%769]); i = i + 2 {
		if this.val[key%769][i] == key {
			return this.val[key%769][i+1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	for i := 0; i < len(this.val[key%769]); i = i + 2 {
		if this.val[key%769][i] == key {
			this.val[key%769] = append(this.val[key%769][0:i], this.val[key%769][i+2:len(this.val[key%769])]...)
			break
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
//leetcode submit region end(Prohibit modification and deletion)
