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

package lc701

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

// 701. 二叉搜索树中的插入操作
// 二叉搜索树的插入总是可以作为叶子节点插入
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 迭代
// 不断二分向下查找，到头就说明到地方了，直接将其作为叶子节点插入
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	cur := root
	for root.Val != val {
		if cur.Val > val {
			if cur.Left == nil {
				cur.Left = node
				break
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				break
			}
			cur = cur.Right
		}
	}
	return root
}
