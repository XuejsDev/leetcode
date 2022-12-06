/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  func levelOrder(root *TreeNode) [][]int {
// 	res := make([][]int, 0)
// 	var dfsLevel func(node *TreeNode, level int)
// 	dfsLevel = func(node *TreeNode, level int) {
// 		if node == nil {
// 			return
// 		}
// 		if len(res) == level {
// 			res = append(res, []int{node.Val})
// 		} else {
// 			res[level] = append(res[level], node.Val)
// 		}
// 		dfsLevel(node.Left, level+1)
// 		dfsLevel(node.Right, level+1)
// 	}
// 	dfsLevel(root, 0)
// 	return res
// }

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		res = append(res, tmp)
	}
	return res
}

// @lc code=end
