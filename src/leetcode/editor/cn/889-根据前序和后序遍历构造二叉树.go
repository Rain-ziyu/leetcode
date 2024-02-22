//给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵
//树的后序遍历，重构并返回二叉树。
//
// 如果存在多个答案，您可以返回其中 任何 一个。
//
//
//
// 示例 1：
//
//
//
//
//输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
//输出：[1,2,3,4,5,6,7]
//
//
// 示例 2:
//
//
//输入: preorder = [1], postorder = [1]
//输出: [1]
//
//
//
//
// 提示：
//
//
// 1 <= preorder.length <= 30
// 1 <= preorder[i] <= preorder.length
// preorder 中所有值都 不同
// postorder.length == preorder.length
// 1 <= postorder[i] <= postorder.length
// postorder 中所有值都 不同
// 保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历
//
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 365 👎 0

package main

import "fmt"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	if len(preorder) > 2 {

		preorder = preorder[1:]
		postorder = postorder[:len(postorder)-1]
		preorderMap := make(map[int]int)
		for i := range preorder {
			preorderMap[preorder[i]] = i
		}
		postorderMap := make(map[int]int)
		for i := range postorder {
			postorderMap[postorder[i]] = i
		}
		fmt.Println(preorder[:preorderMap[postorder[len(postorder)-1]]], postorder[:postorderMap[preorder[0]]+1])
		fmt.Println(preorder[preorderMap[postorder[len(postorder)-1]]:], postorder[postorderMap[preorder[0]]+1:])
		preTree := &TreeNode{}
		postTree := &TreeNode{}
		if postorder[len(postorder)-1] == preorder[0] {
			preTree = constructFromPrePost(preorder[:preorderMap[postorder[len(postorder)-1]]],postorder[postorderMap[preorder[0]]+1:] )

			postTree = constructFromPrePost(preorder[preorderMap[postorder[len(postorder)-1]]:], postorder[:postorderMap[preorder[0]]+1])

		}else {
			preTree = constructFromPrePost(preorder[:preorderMap[postorder[len(postorder)-1]]], postorder[:postorderMap[preorder[0]]+1])

			postTree = constructFromPrePost(preorder[preorderMap[postorder[len(postorder)-1]]:], postorder[postorderMap[preorder[0]]+1:])
		}
		root.Left = preTree
		root.Right = postTree
	} else {
		if len(preorder) == 2 {
			root.Left = &TreeNode{preorder[1], nil, nil}
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
