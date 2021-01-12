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

package lc449

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

type Codec struct {
}

// 449. 序列化和反序列化二叉搜索树
// 重点：编码的字符串应尽可能紧凑。 也就是最好不要在字符串中体现出 nil
// 那么如何在字符串中确认左右树的结束位置呢？
// 利用二叉搜索树的排序性质，左树一定比根小，右树一定比根大
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, strconv.Itoa(root.Val))
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
// 官方解法 O(n)
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nodes := strings.Split(data, ",")
	var traverse func(lower, upper int) *TreeNode
	traverse = func(lower, upper int) *TreeNode {
		if len(nodes) == 0 {
			return nil
		}
		node, _ := strconv.Atoi(nodes[0])
		if node < lower || node > upper {
			return nil
		}
		nodes = nodes[1:]
		root := &TreeNode{Val: node}
		// 利用二叉搜索树性质，放在合适的位置
		root.Left = traverse(lower, root.Val)
		root.Right = traverse(root.Val, upper)
		return root
	}
	// 0 <= Node.val <= 10000
	return traverse(-1, 10001)
}

// Deserializes your encoded data to tree.
// 时间复杂度 O(n ㏒n)
// 不是最好解法
func (this *Codec) deserialize2(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nums := make([]int, 0)
	for _, v := range strings.Split(data, ",") {
		ints, _ := strconv.Atoi(v)
		nums = append(nums, ints)
	}
	node := TreeNode{
		Val: nums[0],
	}
	for i := 0; i < len(nums); i++ {
		helperDeserialize(&node, &TreeNode{Val: nums[i]})
	}
	return &node
}

func helperDeserialize(a *TreeNode, b *TreeNode) {
	if b.Val < a.Val {
		if a.Left != nil {
			helperDeserialize(a.Left, b)
		} else {
			a.Left = b
		}
	}
	if b.Val > a.Val {
		if a.Right != nil {
			helperDeserialize(a.Right, b)
		} else {
			a.Right = b
		}
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
