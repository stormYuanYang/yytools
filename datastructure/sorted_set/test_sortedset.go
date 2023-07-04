// Package test.

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
// 创建日期:2023/6/29
package sorted_set

import (
	"fmt"
	"yytools/algorithm/math"
	"yytools/common/assert"
)

var uniq = int64(0)

type Val struct {
	ID   int64 //唯一id
	Name string
	Meta string
}

func NewVal() *Val {
	uniq++
	return &Val{
		ID:   uniq,
		Meta: "",
	}
}

func (this *Val) LessThan(data Value) bool {
	x := data.(*Val)
	return this.ID < x.ID
}

func (this *Val) EqualTo(data Value) bool {
	x := data.(*Val)
	return this.ID == x.ID
}

const (
	TEST_SORTED_SET_SCORE_MIN = 1
	TEST_SORTED_SET_SCORE_MAX = 750
)

func SortedSetIsLegal(ss *SortedSet) {
	ss.lengthMustEqual()

	current := ss.Sl.Head.Levels[0].Forward
	rank := 1
	for current != nil {
		//print(int(current.Data.Score), "->")
		if current.Levels[0].Forward != nil {
			assert.Assert(current.Data.LessThan(current.Levels[0].Forward.Data),
				"跳跃表表必须是有序的:", fmt.Sprintf("current:%+v, forward:%+v", current, current.Levels[0].Forward))

			data := ss.GetByRank(rank)
			assert.Assert(data != nil, "rank实现有问题:", rank)
			assert.Assert(data.EqualTo(current.Data), "rank实现有问题", rank)
			rank++
		}
		current = current.Levels[0].Forward
	}
}

// 插入
func SortedSetOp_Insert(ss *SortedSet, num int) {
	for i := 0; i < num; i++ {
		n := math.RandInt(TEST_SORTED_SET_SCORE_MIN, TEST_SORTED_SET_SCORE_MAX)
		val := NewVal()
		data := NewNodeData(val.ID, float64(n), val)
		assert.Assert(ss.Insert(data), "插入不会失败:", data)
	}
}

// 更新分数
func SortedSetOp_UpdateScore(ss *SortedSet, num int) {
	for i := 0; i < num; i++ {
		if ss.Length() > 0 {
			randomRank := math.RandInt(1, ss.Length())
			data := ss.GetByRank(randomRank)
			assert.Assert(data != nil, "data 不能为nil, rank:", randomRank)

			n := math.RandInt(TEST_SORTED_SET_SCORE_MIN, TEST_SORTED_SET_SCORE_MAX)
			_, ok := ss.UpdateScore(data.Key, float64(n))
			assert.Assert(ok, "更新分数不能失败, data:", data, " newScore:", n)
		}
	}
}

func SortedSetOp_Delete(ss *SortedSet, num int) {
	for i := 0; i < num; i++ {
		if ss.Length() > 0 {
			randomRank := math.RandInt(1, ss.Length())
			data := ss.GetByRank(randomRank)
			assert.Assert(data != nil, "data 不能为nil, rank:", randomRank)
			ss.Delete(data.Key)
		}
	}
}

func SortedSetOp_GetRangeByScore(ss *SortedSet, num int) {
	if ss.Length() > 0 {
		for i := 0; i < num; i++ {
			min := float64(math.RandInt(TEST_SORTED_SET_SCORE_MIN, TEST_SORTED_SET_SCORE_MAX))
			max := float64(math.RandInt(TEST_SORTED_SET_SCORE_MIN, TEST_SORTED_SET_SCORE_MAX))
			if min > max {
				min, max = max, min
			}

			minEx := false
			if math.RandInt(0, 1) == 1 {
				minEx = true
			}

			maxEx := false
			if math.RandInt(0, 1) == 1 {
				maxEx = true
			}

			datas := ss.GetRangeByScore(min, minEx, max, maxEx)

			// 判断是否有序
			for j := 0; j < len(datas); j++ {
				if j+1 < len(datas) {
					assert.Assert(datas[j].LessThan(datas[j+1]), "返回的元素必须是有序的")
				}
			}

			// 判断返回的元素是否在指定范围内
			for _, one := range datas {
				if minEx {
					assert.Assert(one.Score > min, "分数要在范围内,", "score:", one.Score, " max:", max)
				} else {
					assert.Assert(one.Score >= min, "分数要在范围内,", "score:", one.Score, " max:", max)
				}
				if maxEx {
					assert.Assert(one.Score < max, "分数要在范围内,", "score:", one.Score, " max:", max)
				} else {
					assert.Assert(one.Score <= max, "分数要在范围内,", "score:", one.Score, " max:", max)
				}
			}
		}
	}
}

func SortedSetOp_GetRank(ss *SortedSet, num int) {
	for i := 0; i < num; i++ {
		randomRank := math.RandInt(1, ss.Length())
		data := ss.GetByRank(randomRank)
		assert.Assert(data != nil, "data 不能为nil, rank:", randomRank)

		rank := ss.GetRank(data.Key)
		assert.Assert(randomRank == rank, "排名不一致, randomRank:", randomRank, " rank:", rank, " key:", data.Key)
	}
}

func SortedSetTest() {
	num := 10000
	for j := 0; j < 10; j++ {
		ss := NewSortedSet()
		SortedSetOp_Insert(ss, num)

		SortedSetOp_UpdateScore(ss, num/2)
		SortedSetIsLegal(ss)

		SortedSetOp_GetRangeByScore(ss, 100)
		SortedSetIsLegal(ss)

		SortedSetOp_GetRank(ss, num)
		SortedSetIsLegal(ss)

		delNum := math.RandInt(1, num/2)
		SortedSetOp_Delete(ss, delNum)
		SortedSetIsLegal(ss)

		println("completed:", j)
	}
}