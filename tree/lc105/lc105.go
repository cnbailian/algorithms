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

package lc105

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

// 105. 从前序与中序遍历序列构造二叉树
// 前序遍历时第一个一定是根节点，通过根节点在中序遍历中确定左树和右树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 根节点总是前序排序的第一个
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	// 根据根节点取得在中序排序的索引
	index := getRootIndex(inorder, preorder[0])
	// 根据索引获得左侧节点和右侧节点
	inorderLeft := inorder[:index]
	inorderRight := inorder[index+1:]
	// 根据节点数量拆分前序排序
	root.Left = buildTree(preorder[1:len(inorderLeft)+1], inorderLeft)
	root.Right = buildTree(preorder[len(inorderLeft)+1:], inorderRight)
	return root
}

func getRootIndex(inorder []int, root int) int {
	for i, v := range inorder {
		if v == root {
			return i
		}
	}
	return 0
}
