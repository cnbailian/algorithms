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

package lc236

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	// case 1
	// 说明 root 的左右子树都不包含 p, q
	if left == nil && right == nil {
		return nil
	}
	// case 2
	// left 为空、right 不为空时，说明 p, q 都不在 root 的左子树
	if left == nil {
		return right
	}
	// case 3
	// right 为空、left 不为空时，说明 p, q 都不在 root 的右子树
	if right == nil {
		return left
	}
	// case 4
	// 说明 p, q 分列在 root 的异侧（左/右树），root 为最近公共祖先
	return root
}
