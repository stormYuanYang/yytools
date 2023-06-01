// Package skiplist.

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
// 创建日期:2023/6/1
package skiplist

import (
	"math/rand"
	"yytools/tools/assert"
)

// 随机计算跳跃表中某个结点的高度(等级)
// 高度范围在闭区间[1, maxLevel]内
func randomLevel(maxLevel int, levelUpProbability float64) int {
	assert.Assert(maxLevel <= MAX_NODE_LEVEL,
		"超过设定的最大节点高度", "指定高度:", maxLevel, "限定高度:", MAX_NODE_LEVEL)
	assert.Assert(levelUpProbability >= 0 && levelUpProbability < 1,
		"提升节点高度概率不正确:", levelUpProbability, "正常范围:[0.0,1)")
	level := 1
	for {
		// 满足两个条件就可以提升等级:
		// 1.等级小于等于指定最大等级 且
		// 2.满足指定概率
		// 否则退出循环
		canLevelUp := (level <= maxLevel && rand.Float64() < levelUpProbability)
		if canLevelUp {
			level++
		} else {
			break
		}
	}
	return level
}