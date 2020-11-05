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

package lc98

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

// 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	// 递归每个节点，同时使用参数保留验证信息，一旦任何一个节点出错都会因为 && 而反馈到最上层
	var bst func(root *TreeNode, min *TreeNode, max *TreeNode) bool
	bst = func(root *TreeNode, min *TreeNode, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		return bst(root.Left, min, root) && bst(root.Right, root, max)
	}
	return bst(root, nil, nil)
}

// 中序遍历解法，bst 中序遍历的输出是一个有序数列，所以只需要记录上一个的值就能比对是否正确
func isValidBST2(root *TreeNode) bool {
	var inorder func(root *TreeNode) bool
	var pre *TreeNode
	inorder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		left := inorder(root.Left)
		// 如果当前值小于等于上一个节点值，则证明不是 bst，直接返回 false
		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		right := inorder(root.Right)
		return left && right
	}
	return inorder(root)
}
