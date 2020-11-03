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

package lc222

import (
	"math"
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

// 222. 完全二叉树的节点个数
// 三种二叉树定义
// 完全二叉树：它是一棵空树或者它的叶子节点只出在最后两层，若最后一层不满则叶子节点只在最左侧。
// 满二叉树：每个节点有 0 或 2 个子节点。
// 完美二叉树：除了叶子节点外，所有节点都有两个子节点，并且所有叶子节点拥有相同的高度。
// 完美二叉树的节点数计算：层数为 h，节点数为 2^h -1
// 完全二叉树的节点数计算：层数为 h，节点数为 2^(h-1)-1 + 最后一层节点数
// 所以这道题考虑的是如何更好的算出最后一层节点数。二分法
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 获得树最高层级
	maxLevel := countLevel(root)
	if maxLevel == 1 {
		return 1
	}
	// 获得当前最后一层节点总数，上层节点总数 = 最后一层节点总数 - 1
	lastLayerCount := int(math.Pow(2, float64(maxLevel-1)))
	// 当前二分位置
	index := lastLayerCount / 2
	// 二分数量
	count := 0
	for root != nil {
		// 每次都是找下一层，所以 --
		maxLevel--
		// 获取右树高度
		rightLevel := countLevel(root.Right)
		// 高度一致说明此处最后一层有节点，所以要继续二分到右树中
		if rightLevel == maxLevel {
			root = root.Right
			count += index
		} else {
			// 不一致则要二分回左树
			root = root.Left
		}
		index /= 2
		if index == 0 {
			index = 1
		}
	}
	// 最后一层节点总数 - 1 = 上层节点总数
	return lastLayerCount - 1 + count
}

func countLevel(root *TreeNode) (res int) {
	for root != nil {
		res++
		root = root.Left
	}
	return
}
