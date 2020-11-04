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

package lc145

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

// 144. 二叉树的后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var postorder func(*TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val)
	}
	postorder(root)
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

// 后序遍历：左、右、根
func postorderTraversalStack(root *TreeNode) (res []int) {
	stack := stack{}
	var prev *TreeNode
	for stack.Len() > 0 || root != nil {
		// 先将左侧入栈，相当于 postorder(root.Left)
		for root != nil {
			stack.Add(root)
			root = root.Left
		}
		// 出栈
		node := stack.Pop()
		// 因为顺序是左右根，如果右侧有值，则要重新入栈，并将右侧也入栈，相当于 postorder(root.Right)
		// 右侧无值或右侧值已出栈，则执行出栈逻辑
		if node.Right == nil || node.Right == prev {
			res = append(res, node.Val)
			prev = node
			root = nil
		} else {
			stack.Add(node)
			root = node.Right
		}
	}
	return
}
