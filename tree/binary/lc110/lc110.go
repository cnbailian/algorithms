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

package lc110

import "math"

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

// 110. 平衡二叉树
// 自顶向下比对所有节点的左右子树是否平衡
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 当前树为平衡树
	cur := math.Abs(height(root.Left)-height(root.Right)) < 2
	// 左侧所有节点都是平衡树
	left := isBalanced(root.Left)
	// 右侧所有节点也都是平衡树
	right := isBalanced(root.Right)
	return cur && left && right
}

func height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(height(node.Left), height(node.Right)) + 1
}

// 自底向上计算最大深度，遇到非平衡树直接返回 -1 作为依据向上响应
func isBalanced2(root *TreeNode) bool {
	postorder := func(root *TreeNode) float64 { return 0 }
	postorder = func(root *TreeNode) float64 {
		if root == nil {
			return 0
		}
		left, right := postorder(root.Left), postorder(root.Right)
		// 将错误结果向上响应
		if left == -1 || right == -1 {
			return -1
		}
		// 遇到非平衡树直接响应错误结果
		if math.Abs(left-right) > 1 {
			return -1
		}
		// 后序遍历，自底向上获取高度
		return math.Max(left, right) + 1
	}
	return postorder(root) != -1
}
