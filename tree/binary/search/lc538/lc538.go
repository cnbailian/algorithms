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

package lc538

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

// 538. 把二叉搜索树转换为累加树
// BST 的中序遍历结果是升序，只要调转顺序，就可以变为降序
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Right)
		sum += root.Val
		root.Val = sum
		inorder(root.Left)
	}
	inorder(root)
	return root
}
