/*
Copyright 2020 BaiLian.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lc102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) (res [][]int) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var s []int
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[0] != nil {
				s = append(s, queue[0].Val)
				queue = append(queue, queue[0].Left, queue[0].Right)
			}
			queue = queue[1:]
		}
		if len(s) > 0 {
			res = append(res, s)
		}
	}
	return
}

// 尝试理解递归
// 主要依靠额外参数表示层级，本质还是 dfs
func levelOrder2(root *TreeNode) (res [][]int) {
	var traverse func(root *TreeNode, level int)
	traverse = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		traverse(root.Left, level+1)
		traverse(root.Right, level+1)
	}
	traverse(root, 0)
	return
}
