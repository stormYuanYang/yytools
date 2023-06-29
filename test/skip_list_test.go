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
package test

import (
	"math/rand"
	"yytools/common/assert"
	"yytools/datastructure/skiplist"
)

var uniq = int64(0)

type Data struct {
	ID   int64 //唯一id
	Meta string
}

func NewData() *Data {
	uniq++
	return &Data{
		ID:   uniq,
		Meta: "",
	}
}

func (this *Data) LessThan(data skiplist.NodeData) bool {
	x := data.(*Data)
	return this.ID < x.ID
}

func (this *Data) EqualTo(data skiplist.NodeData) bool {
	x := data.(*Data)
	return this.ID == x.ID
}

func SkipListTest() {
	for j := 0; j < 10; j++ {
		num := int32(10000)
		sl := skiplist.NewSkipList()
		for i := 0; i < int(num); i++ {
			n := rand.Int31n(999) + 1
			sl.Insert(float64(n), NewData())
		}
		
		//upNum := rand.Int31n(num) + 1
		//for i := 0; i < int(upNum); i++ {
		//	n1 := rand.Int31n(999) + 1
		//	n2 := rand.Int31n(999) + 1
		//	sl.Update(float64(n1), float64(n2))
		//}
		//
		//delNum := rand.Int31n(num) + 1
		//for i := 0; i < int(delNum); i++ {
		//	n := rand.Int31n(999) + 1
		//	sl.Delete(float64(n))
		//}
		
		current := sl.Head.Forward[0]
		//rank := 1
		for current != sl.Tail {
			//print(int64(current.Score), "->")
			if current.Forward[0] != sl.Tail {
				assert.Assert(current.Score <= current.Forward[0].Score, "跳跃表表必须是有序的", current, current.Forward[0])
				
				//node, ok := skiplist.GetRank(rank)
				//assert.Assert(ok, "rank实现有问题:", skiplist)
				//assert.Assert(node == current, "rank实现有问题", skiplist)
			}
			current = current.Forward[0]
		}
		//println()
	}
}