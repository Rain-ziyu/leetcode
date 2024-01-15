//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//
//
// 示例 2：
//
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列
//
//
// Related Topics 链表 双指针 👍 1255 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func deleteDuplicates(head *ListNode) *ListNode {
	dump := &ListNode{Next: head}
	preVal := -101
	pre := dump
	for head != nil {
		if head.Val == preVal {
			preVal = head.Val
			pre.Next = head.Next
			head = head.Next
		} else if head.Next != nil && head.Val == head.Next.Val {
			preVal = head.Val
			pre.Next = head.Next.Next
			head = head.Next.Next
		} else {
			preVal = head.Val
			pre = head
			head = head.Next
		}
	}
	return dump.Next

}

//leetcode submit region end(Prohibit modification and deletion)
