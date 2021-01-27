/*
Copyright 2021 BaiLian.

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

package list

// 2. 两数相加
// 与 989 题相比，难点在于链表的个数不固定，取值有困难
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	current := res
	carry := 0
	for l1 != nil || l2 != nil {
		var l1val, l2val int
		if l1 != nil {
			l1val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2val = l2.Val
			l2 = l2.Next
		}
		sum := l1val + l2val + carry
		current.Val = sum % 10
		carry = sum / 10
		if l1 != nil || l2 != nil {
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return res
}
