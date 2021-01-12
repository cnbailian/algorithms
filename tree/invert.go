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

package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Leetcode 226 翻转二叉树
// Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so fuck off.
// 前序遍历
func invertTree(root *TreeNode) *TreeNode {
	// 终止条件
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 后序遍历
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree1(root.Right), invertTree1(root.Left)
	return root
}

// 中序遍历
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree2(root.Left)
	root.Left, root.Right = root.Right, root.Left
	// 已经交换过了，所以还是 left
	invertTree2(root.Left)
	return root
}
