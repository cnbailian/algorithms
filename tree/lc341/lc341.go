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

package lc341

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// 341. 扁平化嵌套列表迭代器
// 整体思路：data slice 作为栈结构保存数据，HasNext 递归拆解列表入栈，Next 出栈
type NestedIterator struct {
	Data []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		Data: nestedList,
	}
}

func (this *NestedIterator) Next() int {
	// pop 头部整数
	res := this.Data[0].GetInteger()
	this.Data = this.Data[1:]
	return res
}

func (this *NestedIterator) HasNext() bool {
	if len(this.Data) == 0 {
		return false
	}
	// 整数正常返回
	next := this.Data[0]
	if next.IsInteger() {
		return true
	}
	// 列表则拆分后入栈，递归检测列表型数据
	this.Data = append(next.GetList(), this.Data[1:]...)
	return this.HasNext()
}

// N 叉树解法
// N 叉树标准框架，叶节点放入数组，树节点递归
func Tree(nestedList []*NestedInteger) *NestedIterator {
	res := &NestedIterator{}
	var traverse = func(nestedList []*NestedInteger) {}
	traverse = func(nestedList []*NestedInteger) {
		for _, integer := range nestedList {
			if integer.IsInteger() {
				res.Data = append(res.Data, integer)
				continue
			}
			traverse(integer.GetList())
		}
	}
	traverse(nestedList)
	return res
}

func (this *NestedIterator) Next() int {
	res := this.Data[0].GetInteger()
	this.Data = this.Data[1:]
	return res
}

func (this *NestedIterator) HasNext() bool {
	return len(this.Data) > 0
}
