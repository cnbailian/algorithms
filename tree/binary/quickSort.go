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

// 快速排序的逻辑是，若要对 nums[lo..hi] 进行排序
// 我们先找一个分界点 p，通过交换元素使得 nums[lo..p-1] 都小于等于 nums[p]，且 nums[p+1..hi] 都大于 nums[p]
// 然后递归地去 nums[lo..p-1] 和 nums[p+1..hi] 中寻找新的分界点，最后整个数组就被排序了。

// 快速排序与二叉树
// 快排通过构造分界点 p 形成两个数组，左侧数组小于分界点 p，右侧数组大于分界点 p
// 随后再对左右数组构造分界点，直到数组整个排序完成
// 从实现上讲就是 二叉树前序遍历
// 前序遍历
func quickSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(nums, lo, hi)
	quickSort(nums, lo, p-1)
	quickSort(nums, p+1, hi)
}

// 通过交换元素，构建分界点
// 分界点左侧数组元素都小于分界点，右侧数组元素都大于分界点
func partition(a []int, lo, hi int) int {
	// 先去最高位为分界值
	pivot := a[hi]
	// 分界点索引
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			// 交换位置，更新分界索引
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	// 最终将分界值交换至分界点前面，以便于后续方便忽略
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
