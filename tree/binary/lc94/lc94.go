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

package lc94

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

// 94. 二叉树的中序遍历
// 递归很简单，不符合 medium 难度
func inorderTraversal(root *TreeNode) (res []int) {
	inorder := func(root *TreeNode) {}
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return
}

// 利用栈迭代
type stack struct {
	Data []*TreeNode
}

func (s *stack) Pop() *TreeNode {
	res := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return res
}

func (s *stack) Add(data *TreeNode) {
	s.Data = append(s.Data, data)
}

func (s *stack) Len() int {
	return len(s.Data)
}

// 中序遍历：左、根、右
func inorderTraversalStack(root *TreeNode) (res []int) {
	stack := stack{}
	for stack.Len() > 0 || root != nil {
		// 入栈，相当于递归中的 inorder(root.Left)
		for root != nil {
			stack.Add(root)
			root = root.Left
		}
		// 出栈，递归中的中序遍历执行逻辑的位置
		node := stack.Pop()
		res = append(res, node.Val)
		// 将 root 为右侧，继续入栈，inorder(root.Right)
		root = node.Right
	}
	return
}
