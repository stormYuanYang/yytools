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
package math

import (
	"yytools/algorithm/math/random"
	"yytools/common/assert"
	"yytools/datastructure/stack"
)

/**
*	提供和概率相关的工具方法
 */

// 根据权重计算概率
// 时间复杂度:O(n)
// 此方法适用于用权重数组计算一次元素的概率
// 如果要反复使用同一个权重数组计算元素概率，则该函数的效率就会稍低
// 此时，使用CalculateIndexListByWeight()更佳
func CalculateIndexByWeight(weightList []int) int {
	assert.Assert(len(weightList) > 0, "权重数组长度要大于0")
	var totalWeight int = 0
	for _, weight := range weightList {
		assert.Assert(weight >= 0, "元素的权重不能小于0:", weight)
		// 计算总权重
		totalWeight += weight
	}
	
	// 考虑一种特殊情况：即数组中所有元素的权重都为0
	// 此时可以认为就是等概率计算各个元素的概率
	if totalWeight == 0 {
		// 等概率随机数组下标
		return random.RandInt(0, len(weightList)-1)
	}
	// 接下来，总权重至少为1
	return calcIndexWithWeight(weightList, totalWeight)
}

// 根据权重数组，得到指定数量的元素下标
func CalculateIndexListByWeight(weightList []int, num int) []int {
	assert.Assert(len(weightList) > 0, "权重数组长度要大于0")
	assert.Assert(num > 0, "需要返回的下标数量大于0.num:", num)
	var totalWeight int = 0
	tmpList := make([]int, len(weightList)+1)
	for i, weight := range weightList {
		assert.Assert(weight >= 0, "元素的权重不能小于0:", weight)
		// 计算总权重
		// TODO 这里累加实际上可能会溢出int32的，暂不考虑溢出时的情况
		totalWeight += weight
		tmpList[i+1] = totalWeight
	}
	
	// 考虑一种特殊情况：即数组中所有元素的权重都为0
	// 此时可以认为就是等概率计算各个元素的概率
	if totalWeight == 0 {
		// 等概率随机数组下标
		indexList := make([]int, num)
		for i := 0; i < int(num); i++ {
			indexList[i] = random.RandInt(0, len(weightList)-1)
		}
		return indexList
	}
	// 接下来，总权重至少为1
	indexList := make([]int, num)
	for i := 0; i < int(num); i++ {
		indexList[i] = calcIndexWithWeightByBinarySearch(tmpList)
	}
	return indexList
}

func calcIndexWithWeight(weightList []int, totalWeight int) int {
	assert.Assert(totalWeight > 0, "总权重需要大于0：", totalWeight)
	// 先根据总权重计算一个随机值，范围在[1,totalWeight]
	randNum := random.RandInt(1, totalWeight)
	for i, weight := range weightList {
		// 最后一次循环后，newTotalWeight会等于0,此时必然有randNum > newTotalWeight
		newTotalWeight := totalWeight - weight
		if randNum > newTotalWeight {
			// 命中区间
			return int(i)
		} else {
			// 未命中，则通过减去当前区间权重 得到新的总权重
			totalWeight = newTotalWeight
		}
	}
	// 直接断言 逻辑不应该执行到这里
	assert.Assert(false, "未命中任何区间,randNum:", randNum, "totalWeight:", totalWeight)
	// 为满足golang语法这里返回一个数字 但逻辑不能走这里返回
	return -1
}

func calcIndexWithWeightByBinarySearch(tmpList []int) int {
	assert.Assert(len(tmpList) > 0, "数组长度要大于0")
	totalWeight := tmpList[len(tmpList)-1]
	assert.Assert(totalWeight > 0, "总权重需要大于0：", totalWeight)
	// 先根据总权重计算一个随机值，范围在[1,totalWeight]
	randNum := random.RandInt(1, totalWeight)
	index := binarySearchInRange(tmpList, randNum)
	assert.Assert(index != -1, "未命中任何区间,randNum:", randNum, "totalWeight:", totalWeight)
	return index
}

// 在区间中进行二分查找（区别于普通的二分查找）
// (l[0], l[1]]
// (l[1], l[2]]
// (l[2], l[3]]
// ...
// (l[n-1], l[n]]
func binarySearchInRange(tmpList []int, n int) int {
	length := len(tmpList)
	if length == 0 {
		return -1
	}
	
	for low, high := 1, length-1; low <= high; {
		mid := low + (high-low)/2
		
		if n > tmpList[mid-1] && n <= tmpList[mid] {
			// 命中
			return mid
		} else if n > tmpList[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

type InterfaceProbabilityDistribution interface {
	Generate() int
}

// Vose's Alias Method(Vose的别名方法)
// 十分高效而优雅的实现方式
// 初始化阶段:
// 时间复杂度O(n),空间复杂度O(n)
// 生成阶段:
// 时间复杂度O(1) [初始化阶段决定了:生成器本身就具有O(n)的空间复杂度]
// 参考: https://www.keithschwarz.com/darts-dice-coins/
type VoseAliasMethod struct {
	Prob  []float // 概率数组
	Alias []int   // 别名数组(记录的是下标)
}

type float float64

func NewVoseAliasMethod(weights []int) *VoseAliasMethod {
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}
	assert.Assert(totalWeight > 0, "总权重肯定要大于0:", totalWeight)
	n := len(weights)
	prob := make([]float, n)
	alias := make([]int, n)
	
	small := stack.NewStackWithSize(n / 2)
	large := stack.NewStackWithSize(n / 2)
	
	// 初始化概率数组(每个原概率值都乘以n,等比例放大;总体的概率分布是不变的)
	// 概率值等比例放大后，概率平均值为1，概率总和为n
	for i := 0; i < n; i++ {
		prob[i] = float(weights[i]) / float(totalWeight) * float(n)
		// 设置小概率和大概率
		if prob[i] < 1 {
			small.Push(i)
		} else {
			large.Push(i)
		}
	}
	
	// 当小概率集合和大概率集合都不为空时，循环处理
	// 直到有一个集合为空，则结束循环
	for !small.Empty() && !large.Empty() {
		l := small.Pop().(int) // 小概率的下标
		g := large.Pop().(int) // 大概率的下标
		alias[l] = g           // 别名数组记录另一部分概率的下标
		// This is a more numerically stable option
		// 比起 prob[g] - (1 - prob[l])数值更稳定
		prob[g] = prob[g] + prob[l] - 1
		if prob[g] < 1 {
			small.Push(g) // 下标压入小概率下标集合
		} else {
			large.Push(g) // 下标压入大概率下标集合
		}
	}
	
	// 判断大概率下标集合
	for !large.Empty() {
		g := large.Pop().(int) // 得到下标
		prob[g] = 1            // 经过前面的处理，这里的概率肯定是1
		alias[g] = -1          // 对应别名下标(该情况下，没有对应别名所以其下标设置为一个无效的-1)
	}
	// 判断小概率下标集合
	for !small.Empty() {
		// 能进入这里是因为数值精度的不稳定(This is only possible due to numerical instability)
		l := small.Pop().(int) // 得到下标
		prob[l] = 1            // 经过前面的处理，这里的概率肯定是1
		alias[l] = -1          // 对应别名下标(该情况下，没有对应别名所以其下标设置为一个无效的-1)
	}
	// 至此，概率数组和别名数组都已构建完成
	
	ret := &VoseAliasMethod{
		Prob:  prob,
		Alias: alias,
	}
	return ret
}

// 返回概率(权重)数组的某一概率(权重)的下标
// 效率非常高的：时间复杂度O(1)
func (this *VoseAliasMethod) Generate() int {
	n := len(this.Prob)
	i := random.RandInt(0, n-1) // 随机得到一个概率数组的下标
	p := random.RandFloat64()   // 范围：[0.0,1.0)
	// 在[0.0, 1.0)范围内随机得到一个值,用这个值去判断得到该概率还是其别名
	if float(p) < this.Prob[i] {
		return i
	} else {
		index := this.Alias[i] // 别名数组保存的是下标
		assert.Assert(index >= 0 && index < n, "out of range:", index)
		return index
	}
}