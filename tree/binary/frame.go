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

package binary

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

func Traverse(root *TreeNode) {
	// 前序遍历
	Traverse(root.Left)
	// 中序遍历
	Traverse(root.Right)
	// 后序遍历
}
