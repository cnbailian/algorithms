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

package lc230

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

// 230. 二叉搜索树中第K小的元素
// 二叉搜索树（BST）的定义：每一个节点 node，左子树节点的值都比 node 要小，右子树节点的值都比 node 要大，bst 的每一个节点，它的左右树都是 bst。
// BST 的中序遍历的结果（左、根、右）是有序的。
func kthSmallest(root *TreeNode, k int) int {
	var ordered []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		ordered = append(ordered, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return ordered[k-1]
}

// 既然是有序的，可以不用等所有值出来就可以找到第 k 个节点
func kthSmallest2(root *TreeNode, k int) int {
	var res, index int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		index++
		if k == index {
			res = root.Val
		}
		inorder(root.Right)
	}
	inorder(root)
	return res
}

// 进阶
// 如果节点知道自己当前排名，那么就能根据二叉树的特性将时间复杂度降为 O(logN)
// 因为是有序的，能知道向左树找还是向右树找
