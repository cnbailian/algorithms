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

package lc114

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

// 114. 二叉树展开为链表
// 思路：首先是规定为原地展开，所以要在同一颗树上操作
// 左侧节点平行，右侧节点移至左侧末尾
// 后序遍历
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 左侧节点移至右侧
	right := root.Right
	root.Right, root.Left = root.Left, nil
	// 原右侧节点移至左侧末尾
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}
