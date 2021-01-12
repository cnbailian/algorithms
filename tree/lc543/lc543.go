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

package lc543

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

// 543. 二叉树的直径
// 将问题转化为求每个节点最大路径的最大值
// 二叉树总和框架还不熟悉，需要多练习
func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	if root == nil {
		return res
	}
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := traverse(root.Left)
		r := traverse(root.Right)
		res = max(l+r, res)
		return max(l, r) + 1
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
