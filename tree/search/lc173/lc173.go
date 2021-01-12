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

package lc173

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

type BSTIterator struct {
	stack stack
}

// 173. 二叉搜索树迭代器
func Constructor(root *TreeNode) BSTIterator {
	res := BSTIterator{}
	// 将最左树入栈，这样第一个 pop 的一定是最小值
	if root != nil {
		res.stack.Add(root)
		root = root.Left
	}
	return res
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack.Pop()
	// 如果左右树都为空，证明已经是最小值
	if node.Left == nil && node.Right == nil {
		return node.Val
	}
	// 右树的值比节点大，所以要先入栈
	if node.Right != nil {
		this.stack.Add(node.Right)
		// 清理，避免二次入栈
		node.Right = nil
	}
	// 因为节点并非此次所选的值，所以再次入栈
	this.stack.Add(node)
	// 最后入栈左树
	if node.Left != nil {
		this.stack.Add(node.Left)
		node.Left = nil
	}
	// 递归，直到找到最小值，这算不算 O(1) 时间复杂度呢
	return this.Next()
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stack.Len() > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
