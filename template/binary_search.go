// Package algorithm.

// 版权所有(Copyright)[yangyuan]
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// 二分搜索

// 作者:  yangyuan
// 创建日期:2023/7/7
package template

import (
	"yytools/common/assert"
)

// 经典实现一: 在有序数组里查找指定的元素
// 正确执行的条件: 数组是有序的(允许重复的元素出现)
// nums必须是非降序(即排除重复元素的话，就是升序)
// 可以用力扣题目验证方法的正确性: https://leetcode.cn/problems/binary-search/submissions/
func BinarySearch(nums []int, target int) int{
	// 左右均为闭区间
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 有另外一种计算方式: (left+right)/2
		// 但是这种方式不太好，当left和right都比较大时容易出现整型溢出
		// 如下实现可以避免整型溢出
		mid := left + (right - left) / 2
		// 二分搜索代码实现框架并不难，难在边界上的处理
		// 边界处理不好，程序的正确性就是玄学
		// 把每种情况的条件和处理都写清楚
		if target == nums[mid] {
			// 找到
			return mid
		} else if target < nums[mid] {
			// 目标值小于nums[mid]——目标值位于nums[mid]左侧
			// 左边界不变，收缩右边界，在范围[left, mid-1]中继续查找
			right = mid - 1
		} else if target > nums[mid] {
			// 目标值大于nums[mid]——目标值位于nums[mid]右侧
			// 右边界不变，收缩左边界，在范围[mid+1,right]中继续查找
			left = mid + 1
		} else {
			// 不可能执行到这里! target(小于|等于|大于)nums[mid] 三种情况前文都已列举
			assert.Assert(false, "logic error, mid:", mid, " target:", target)
		}
	}
	// 结束循环时, right 小于 left
	// 不停缩小的查找范围里没有元素了,仍然没有找到目标值,则目标值不存在
	return -1
}