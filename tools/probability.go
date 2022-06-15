// Package tools.

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

// 作者:  yangyuan
// 创建日期:2022/6/15
package tools

import (
	"strconv"
)

/**
*	提供和概率相关的工具方法
 */

// 根据权重计算概率
func CalculateIndexByWeight(weightList []int32) int32 {
	Assert(len(weightList) > 0, "权重数组长度要大于0")
	var totalWeight int32 = 0
	for _, weight := range weightList {
		Assert(weight >= 0, "元素的权重不能小于0", strconv.Itoa(int(weight)))
		// 计算总权重
		// TODO 这里累加实际上可能会溢出int32的，暂不考虑溢出时的情况
		totalWeight += weight
	}
	
	// 考虑一种特殊情况：即数组中所有元素的权重都为0
	// 此时可以认为就是等概率计算各个元素的概率
	if totalWeight == 0 {
		// 等概率随机数组下标
		return RandInt32(0, int32(len(weightList))-1)
	}
	// 接下来，总权重至少为1
	// 先根据总权重计算一个随机值，范围在[1,totalWeight]
	Assert(totalWeight > 0, "总权重需要大于0：", strconv.Itoa(int(totalWeight)))
	randNum := RandInt32(1, totalWeight)
	for i, weight := range weightList {
		newTotalWeight := totalWeight - weight
		if randNum > newTotalWeight {
			// 命中区间
			return int32(i)
		} else {
			// 未命中，则通过减去当前区间权重 得到新的总权重
			totalWeight = newTotalWeight
		}
	}
	return -1
}