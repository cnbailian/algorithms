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

package lc589

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

// 589. N叉树的前序遍历
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		for _, child := range root.Children {
			traverse(child)
		}
	}
	traverse(root)
	return res
}

// 利用栈迭代
type stack struct {
	Data []*Node
}

func (s *stack) Pop() *Node {
	res := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return res
}

func (s *stack) Add(data *Node) {
	s.Data = append(s.Data, data)
}

func (s *stack) Len() int {
	return len(s.Data)
}

func preorder2(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := stack{Data: []*Node{root}}
	for stack.Len() > 0 {
		node := stack.Pop()
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.Add(node.Children[i])
		}
	}
	return res
}
