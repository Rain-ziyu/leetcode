//给你一个链表的头节点 head 。
//
// 移除每个右侧有一个更大数值的节点。
//
// 返回修改后链表的头节点 head 。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [5,2,13,3,8]
//输出：[13,8]
//解释：需要移除的节点是 5 ，2 和 3 。
//- 节点 13 在节点 5 右侧。
//- 节点 13 在节点 2 右侧。
//- 节点 8 在节点 3 右侧。
//
//
// 示例 2：
//
//
//输入：head = [1,1,1,1]
//输出：[1,1,1,1]
//解释：每个节点的值都是 1 ，所以没有需要移除的节点。
//
//
//
//
// 提示：
//
//
// 给定列表中的节点数目在范围 [1, 10⁵] 内
// 1 <= Node.val <= 10⁵
//
//
// Related Topics 栈 递归 链表 单调栈 👍 82 👎 0

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNodes(head *ListNode) *ListNode {
	//先新建一个空节点指向head 并且保证这个节点是最大的
	dump := ListNode{Next: head, Val: 1<<31 - 1}
	dump1 := dump.Next
	nodes := make([]*ListNode, 0)
	nodes = append(nodes, &dump)
	for dump1 != nil {
		for len(nodes) != 0 && nodes[len(nodes)-1].Val < dump1.Val {
			//出栈
			nodes = nodes[:len(nodes)-1]
			//删除该节点
			nodes[len(nodes)-1].Next = dump1
		}
		//入栈
		nodes = append(nodes, dump1)
		//dump后移
		dump1 = dump1.Next
	}
	return nodes[0].Next
}

//leetcode submit region end(Prohibit modification and deletion)
