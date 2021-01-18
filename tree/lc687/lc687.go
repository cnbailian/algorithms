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

package lc687

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

// 687. 最长同值路径
// 前置简单题：543，本题多了"同值"的要求，也就是需要父节点的判断（f == root.val）和 path 的剪枝（else{ return 0 }）
func longestUnivaluePath(root *TreeNode) int {
	var res int
	if root == nil {
		return res
	}
	var traverse func(root *TreeNode, f int, path int) int
	traverse = func(root *TreeNode, f int, path int) int {
		if root == nil {
			return 0
		}
		l := traverse(root.Left, root.Val, path)
		r := traverse(root.Right, root.Val, path)
		res = max(l+r, res)
		if f == root.Val {
			return max(l, r) + 1
		} else {
			return 0
		}
	}
	traverse(root, -1, 0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
