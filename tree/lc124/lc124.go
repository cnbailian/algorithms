/*
Copyright 2021 BaiLian.

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

package lc124

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

// 124. 二叉树中的最大路径和
// 相比 687，少了同值的要求，变为计算路径上所有节点值的总和
// 但因为有负数的存在，所以还要某种剪枝（max(traverse(root.Left), 0)）
// 配不上 hard！
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := root.Val
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(traverse(root.Left), 0)
		r := max(traverse(root.Right), 0)
		res = max(l+r+root.Val, res)
		return max(l, r) + root.Val
	}
	traverse(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
