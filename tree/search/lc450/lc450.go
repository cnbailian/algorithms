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

package lc450

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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	// 已找到目标节点
	// case1 左树无值
	if root.Left == nil {
		return root.Right
	}
	// case2 右树无值
	if root.Right == nil {
		return root.Left
	}
	// case3 查找后续，后续就是自身大的最小节点，也就是右树的最左边
	minNode := root.Right
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	root.Val = minNode.Val
	// 删除已被替换的后续节点
	root.Right = deleteNode(root.Right, minNode.Val)
	return root
}

// 查找前驱，也就是比当前节点小的最大节点
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	// 已找到目标节点
	// case1 左树无值
	if root.Left == nil {
		return root.Right
	}
	// case2 右树无值
	if root.Right == nil {
		return root.Left
	}
	// case3 查找前驱，也就是左树的最右边
	minNode := root.Left
	for minNode.Right != nil {
		minNode = minNode.Right
	}
	root.Val = minNode.Val
	// 删除已被替换的前驱节点
	root.Left = deleteNode(root.Left, minNode.Val)
	return root
}
