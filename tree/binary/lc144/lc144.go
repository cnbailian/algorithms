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

package lc144

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

// 144. 二叉树的前序遍历
func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		vals = append(vals, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return vals
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

func preorderTraversalStack(root *TreeNode) (res []int) {
	stack := stack{}
	// 前序遍历从根开始，所以直接放入
	stack.Add(root)
	for stack.Len() > 0 {
		// 出栈，执行前序遍历逻辑
		root = stack.Pop()
		if root == nil {
			continue
		}
		res = append(res, root.Val)
		// 因为栈先入后出的特性，这里要先放入 right，才能正确的执行前序遍历 根、左、右的顺序
		stack.Add(root.Right)
		stack.Add(root.Left)
	}
	return
}
