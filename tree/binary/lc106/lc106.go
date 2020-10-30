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

package lc106

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

// 106. 从中序与后序遍历序列构造二叉树
// 思路与 105 一样，后序遍历最后一个是根节点，根据根节点在中序遍历中寻找左树和右树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	// 根节点总是后序排序的最后一个
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(postorder) == 1 {
		return root
	}

	// 根据根节点取得在中序排序的索引
	index := getRootIndex(inorder, postorder[len(postorder)-1])
	// 根据索引获得左侧节点和右侧节点
	inorderLeft := inorder[:index]
	inorderRight := inorder[index+1:]
	// 根据节点数量拆分后序排序
	root.Left = buildTree(inorderLeft, postorder[:len(inorderLeft)])
	root.Right = buildTree(inorderRight, postorder[len(inorderLeft):len(postorder)-1])
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
