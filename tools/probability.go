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
// 时间复杂度:O(n)
// 此方法适用于用权重数组计算一次元素的概率
// 如果要反复使用同一个权重数组计算元素概率，则该函数的效率就会稍低
// 此时，使用CalculateIndexListByWeight()更佳
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
	return calcIndexWithWeight(weightList, totalWeight)
}

// 根据权重数组，得到指定数量的元素下标
func CalculateIndexListByWeight(weightList []int32, num int32) []int32 {
	Assert(len(weightList) > 0, "权重数组长度要大于0")
	Assert(num > 0, "需要返回的下标数量大于0.num:", strconv.Itoa(int(num)))
	var totalWeight int32 = 0
	tmpList := make([]int32, len(weightList)+1)
	for i, weight := range weightList {
		Assert(weight >= 0, "元素的权重不能小于0", strconv.Itoa(int(weight)))
		// 计算总权重
		// TODO 这里累加实际上可能会溢出int32的，暂不考虑溢出时的情况
		totalWeight += weight
		tmpList[i+1] = totalWeight
	}
	
	// 考虑一种特殊情况：即数组中所有元素的权重都为0
	// 此时可以认为就是等概率计算各个元素的概率
	if totalWeight == 0 {
		// 等概率随机数组下标
		indexList := make([]int32, num)
		for i := 0; i < int(num); i++ {
			indexList[i] = RandInt32(0, int32(len(weightList))-1)
		}
		return indexList
	}
	// 接下来，总权重至少为1
	indexList := make([]int32, num)
	for i := 0; i < int(num); i++ {
		indexList[i] = calcIndexWithWeightByBinarySearch(tmpList)
	}
	return indexList
}

func calcIndexWithWeight(weightList []int32, totalWeight int32) int32 {
	Assert(totalWeight > 0, "总权重需要大于0：", strconv.Itoa(int(totalWeight)))
	// 先根据总权重计算一个随机值，范围在[1,totalWeight]
	randNum := RandInt32(1, totalWeight)
	for i, weight := range weightList {
		// 最后一次循环后，newTotalWeight会等于0,此时必然有randNum > newTotalWeight
		newTotalWeight := totalWeight - weight
		if randNum > newTotalWeight {
			// 命中区间
			return int32(i)
		} else {
			// 未命中，则通过减去当前区间权重 得到新的总权重
			totalWeight = newTotalWeight
		}
	}
	// 直接断言 逻辑不应该执行到这里
	Assert(false, "未命中任何区间,randNum:", strconv.Itoa(int(randNum)),
		"totalWeight:", strconv.Itoa(int(totalWeight)))
	// 为满足golang语法这里返回一个数字 但逻辑不能走这里返回
	return -1
}

func calcIndexWithWeightByBinarySearch(tmpList []int32) int32 {
	Assert(len(tmpList) > 0, "数组长度要大于0")
	totalWeight := tmpList[len(tmpList)-1]
	Assert(totalWeight > 0, "总权重需要大于0：", strconv.Itoa(int(totalWeight)))
	// 先根据总权重计算一个随机值，范围在[1,totalWeight]
	randNum := RandInt32(1, totalWeight)
	index := binarySearchInRange(tmpList, randNum)
	Assert(index != -1, "未命中任何区间,randNum:", strconv.Itoa(int(randNum)),
		"totalWeight:", strconv.Itoa(int(totalWeight)))
	return index
}

// 在区间中进行二分查找（区别于普通的二分查找）
// (l[0], l[1]]
// (l[1], l[2]]
// (l[2], l[3]]
// ...
// (l[n-1], l[n]]
func binarySearchInRange(tmpList []int32, n int32) int32 {
	length := len(tmpList)
	if length == 0 {
		return -1
	}
	
	for low, high := 1, length-1; low <= high; {
		mid := low + (high-low)/2
		
		if n > tmpList[mid-1] && n <= tmpList[mid] {
			// 命中
			return int32(mid)
		} else if n > tmpList[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}