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

package lc297

import (
	"strconv"
	"strings"
)

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

// 297. 二叉树的序列化与反序列化
// bfs，层次遍历
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// pop
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			res = append(res, "#")
			continue
		}
		res = append(res, strconv.Itoa(cur.Val))
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
// 层次遍历中每个节点与下层节点的位置是固定的，所以只需要偏移数组索引就可以确定位置
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	if len(nodes) == 0 {
		return nil
	}
	if nodes[0] == "#" {
		return nil
	}
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(nodes[0])
	queue := []*TreeNode{root}
	for i := 0; i < len(nodes); {
		if len(queue) == 0 {
			return root
		}
		node := queue[0]
		queue = queue[1:]
		// 取 left 节点
		i++
		left := nodes[i]
		// 只向队列中添加有值的节点，保证索引位置
		if left != "#" {
			node.Left = &TreeNode{}
			node.Left.Val, _ = strconv.Atoi(left)
			queue = append(queue, node.Left)
		}
		// 取 right 节点
		i++
		right := nodes[i]
		if right != "#" {
			node.Right = &TreeNode{}
			node.Right.Val, _ = strconv.Atoi(right)
			queue = append(queue, node.Right)
		}
	}
	return root
}

// 中序遍历不可以，因为没办法确定根节点位置
// 后序遍历
type PostorderCodec struct {
}

func PostorderConstructor() PostorderCodec {
	return PostorderCodec{}
}

// Serializes a tree to a single string.
func (this *PostorderCodec) serialize(root *TreeNode) string {
	var res []string
	var postorder func(*TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			res = append(res, "#")
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, strconv.Itoa(root.Val))
	}
	postorder(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *PostorderCodec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	var postorder func() *TreeNode
	postorder = func() *TreeNode {
		if len(nodes) == 0 {
			return nil
		}
		last := nodes[len(nodes)-1]
		nodes = nodes[0 : len(nodes)-1]
		if last == "#" {
			return nil
		}
		val, _ := strconv.Atoi(last)
		root := &TreeNode{Val: val}
		// 后序遍历要保证先执行 Right 的递归
		root.Right = postorder()
		root.Left = postorder()
		return root
	}
	return postorder()
}

// 前序遍历
type PreorderCodec struct {
}

func PreorderConstructor() PreorderCodec {
	return PreorderCodec{}
}

// Serializes a tree to a single string.
func (this *PreorderCodec) serialize(root *TreeNode) string {
	var res []string
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			res = append(res, "#")
			return
		}
		res = append(res, strconv.Itoa(root.Val))
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *PreorderCodec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	var preorder func() *TreeNode
	preorder = func() *TreeNode {
		if len(nodes) == 0 {
			return nil
		}
		first := nodes[0]
		nodes = nodes[1:]
		if first == "#" {
			return nil
		}
		val, _ := strconv.Atoi(first)
		root := &TreeNode{Val: val}
		root.Left = preorder()
		root.Right = preorder()
		return root
	}
	return preorder()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
