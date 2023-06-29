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
// 创建日期:2023/6/2
package skiplist

// TODO span的实现还有bug

import (
	"yytools/common/assert"
)

type ISkipList interface {
	Get(score float64) (*Node, bool)
	Insert(score float64) (*Node, bool)
	Delete(score float64) (*Node, bool)
	Update(score float64) (*Node, bool)
	GetRank(rank int) (*Node, bool)
	GetRanks(minRank int, maxRank int) ([]*Node, bool)
	GetRange(minScore float64, maxScore float64) ([]*Node, bool)
}

type SkipList struct {
	Head            *Node   // 头结点(哨兵结点)
	Tail            *Node   // 尾结点(哨兵结点)
	Length          int     // 结点总数(不包含头结点和尾结点)
	MaxLevel        int     // 结点最大高度(限定值,不能超过这个高度)
	CurrentMaxLevel int     // 链表中当前结点的最大高度(有效结点中的最高的高度)
	LevelUpProb     float32 // 提升结点高度的概率
}

type NodeData interface {
	LessThan(n NodeData) bool
	EqualTo(y NodeData) bool
}

type Node struct {
	Forward  []*Node  // 向前的结点数组
	Backward *Node    // 上一个结点(这样最下层就是双向链表，方便向后的遍历)
	Score    float64  // 分数
	Span     int      // 结点之间的跨度(方便取结点的排名) 默认跨度是基于1的
	Data     NodeData // 卫星数据
}

/*
	method of Node
*/
func CreateNode(level int, score float64, data NodeData) *Node {
	node := &Node{
		Forward:  make([]*Node, level),
		Backward: nil,
		Score:    score,
		Span:     0,
		Data:     data,
	}
	return node
}

func (this *Node) High() int {
	return len(this.Forward)
}

/*
	method of SkipList
*/
func NewSkipList() *SkipList {
	return NewSkipListByParams(MAX_NODE_LEVEL, DEFAULT_LEVELUP_PROBABILITY)
}

func NewSkipListByParams(maxLevel int, nodeLevelUpProb float32) *SkipList {
	assert.Assert(maxLevel <= MAX_NODE_LEVEL,
		"超过设定的最大节点高度", "指定高度:", maxLevel, "限定高度:", MAX_NODE_LEVEL)
	assert.Assert(nodeLevelUpProb >= 0 && nodeLevelUpProb < 1,
		"提升节点高度概率不正确:", nodeLevelUpProb, "正常范围:[0.0,1)")
	
	skipList := &SkipList{
		Head:            CreateNode(maxLevel, 0, nil),
		Tail:            CreateNode(0, 0, nil), // 尾结点不需要向前的结点索引数组
		Length:          0,
		MaxLevel:        maxLevel,
		CurrentMaxLevel: 0,
		LevelUpProb:     nodeLevelUpProb,
	}
	for i := 0; i < maxLevel; i++ {
		skipList.Head.Forward[i] = skipList.Tail
	}
	return skipList
}

func (this *SkipList) Get(score float64, data NodeData) (*Node, bool) {
	prev := this.Head
	for i := this.CurrentMaxLevel - 1; i >= 0; i-- {
		// 满足两个条件，查找就继续在当前层级向前继续查找：
		// 1.指定分数大于当前结点的分数，说明要查找的节点一定在当前结点的前方;
		// 2.当前节点不能是(哨兵)尾结点
		current := prev.Forward[i]
		for (score > current.Score || score == current.Score && current.Data.LessThan(data)) && current != this.Tail {
			prev = current
			current = current.Forward[i]
		}
		if (score == current.Score && current.Data.EqualTo(data)) && current != this.Tail {
			// 找到了
			return current, true
		}
		// 否则,说明
		// 1.score分数比当前结点的分数更大,或者
		// 2.当前结点是尾结点
		// 所以再次进入循环，到更低层级去查找
	}
	return nil, false
}

func (this *SkipList) Insert(score float64, data NodeData) (*Node, bool) {
	level := randomLevel(this.MaxLevel, this.LevelUpProb)
	height := level
	if height < this.CurrentMaxLevel {
		height = this.CurrentMaxLevel
	}
	update := make([]*Node, height)
	spans := make([]int, height)
	/*
		跨度是基于1的
		跳跃表结点记录的是，当前结点最高层级的链表的前置结点的跨度，是一个相对值
		也正因为如此，就可以通过累计这些相对值，从而取得指定结点的排名
	
		新插入节点，要考虑到两部分：
		1.新结点自身的跨度；
	
		2.新结点影响其他结点的跨度
			2.1比新结点层级高的结点
			在插入新结点时，比新结点层级高的每一层的链表中，如果新结点刚好位于当前层级的左右结点中间
			那么该右结点的跨度也需要+1
			思考一下一个食堂排队的队列，你和小明、小红分别排在第7,5，3位，你到小明的跨度是2，小明到小红的跨度是2，
			这时中间有一个人插队插到小红的后面一个位置，那么你和小明的跨度依旧不变仍然是2,而小明到小红的跨度就变为3
			你们的位置变为：8，6，3
			2.3比新结点层级低的结点(包含同层级)
			1.如果结点算跨度时不会越过新结点，即
			算跨度的两个结点要么都在新结点前，要么都在新结点后
			那么就不可能被新结点影响跨度
			2.或者新结点插入到两个结点中间(顶替左边结点重新算跨度),则可能会影响右侧结点的跨度
	*/
	// 头结点的跨度是0，因为没有比头结点更靠后的结点了
	prev := this.Head
	for i := this.CurrentMaxLevel - 1; i >= 0; i-- {
		current := prev.Forward[i]
		for (score > current.Score || score == current.Score && current.Data.LessThan(data)) && current != this.Tail {
			// 指针向前行进说明，要插入的结点在当前结点的前方
			// 那么加上当前结点的跨度
			spans[i] += current.Span
			
			// 当前层链表，指针向前行进
			prev = current
			current = current.Forward[i]
		}
		
		if (score == current.Score && current.Data.EqualTo(data)) && current != this.Tail {
			// 能找到，则不能插入(如果逻辑走到这里)
			return current, false
		}
		
		// 需要记录每一层的前置结点
		update[i] = prev
		// 否则继续进入下一层级，直到最后进入最低的层级后退出循环
	}
	
	if level > this.CurrentMaxLevel {
		// 比跳跃表原有的结点高度还高,则需要将头结点作为高层级的前置结点
		for i := this.CurrentMaxLevel; i < level; i++ {
			update[i] = this.Head
		}
		// 更新最大高度
		this.CurrentMaxLevel = level
	}
	
	/*
		1.创建新的结点,设置相关数据
		2.并插入指定位置,并且
		3.更新和维护结点每一层的索引关系
	*/
	// <span 1>在新结点层级上面的链表，在其右侧第一个结点跨度+1
	for i := level; i < height; i++ {
		forward := update[i].Forward[i]
		if forward != this.Tail {
			forward.Span++
		}
	}
	
	span := 0
	// <span 2>设置新结点右侧结点的跨度 因为新结点的插入其跨度可能就减小了
	for i := 0; i < level; i++ {
		forward := update[i].Forward[i]
		if len(forward.Forward) == i+1 && forward != this.Tail {
			forward.Span -= span
		}
		span += spans[i]
	}
	newNode := CreateNode(level, score, data)
	// <span 3>最后设置新结点的跨度
	newNode.Span = span + 1
	
	// 最下面一层实际是个双向链表
	// 先将向后(方向)的链表链接起来
	newNode.Backward = update[0]
	update[0].Forward[0].Backward = newNode
	// 然后将每一层向前(方向)的链表都重新链接起来
	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}
	// 更新跳跃表中的结点总数
	this.Length++
	
	// 至此，结点正确插入到跳跃表中
	// 并且结点间的索引关系也得到正确维护
	return newNode, true
}

func (this *SkipList) Delete(score float64, data NodeData) (*Node, bool) {
	update := make([]*Node, this.CurrentMaxLevel)
	spans := make([]int, this.CurrentMaxLevel)
	prev := this.Head
	for i := this.CurrentMaxLevel - 1; i >= 0; i-- {
		current := prev.Forward[i]
		for (score > current.Score || score == current.Score && current.Data.LessThan(data)) && current != this.Tail {
			spans[i] += current.Span
			
			prev = current
			current = current.Forward[i]
		}
		// 更新每一层级的前置结点
		update[i] = prev
		// 继续到下一层级查找和处理
	}
	
	// 没有找到对应的结点的前置结点
	if update[0] == nil {
		return nil, false
	}
	// 没有找到对应的结点
	current := update[0].Forward[0]
	if current == this.Tail || current.Score != score || !current.Data.EqualTo(data) {
		return nil, false
	}
	/*
		1.移除结点
		2.处理结点每一层的索引关系
	*/
	// <span 1>比当前结点层级高的右侧结点跨度减一
	for i := len(current.Forward); i < this.CurrentMaxLevel; i++ {
		forward := update[i].Forward[i]
		if forward != this.Tail {
			forward.Span--
		}
	}
	span := 0
	// <span 2>小于等于当前结点层级的右结点跨度添加
	for i := 0; i < len(current.Forward); i++ {
		if len(current.Forward[i].Forward) == i+1 && current.Forward[i] != this.Tail {
			current.Forward[i].Span += span
		}
		span += spans[i]
	}
	// 先删除反方向链表的链接
	current.Forward[0].Backward = update[0]
	current.Backward = nil
	// 然后删除每一层级里向前的链表的链接
	for i := 0; i < len(current.Forward); i++ {
		update[i].Forward[i] = current.Forward[i]
		current.Forward[i] = nil
	}
	// 结点总数减一
	this.Length--
	
	// 至此，指定某个结点已被删除
	// 并正确维护了剩余结点间的关系
	return current, true
}

func (this *SkipList) Update(oldScore float64, newScore float64, data NodeData) (*Node, bool) {
	oldNode, ok := this.Get(oldScore, data)
	if !ok {
		// 找不到结点，直接插入新的
		return this.Insert(newScore, data)
	}
	
	// 先看能不能复用之前的结点对象
	// 如果新分数和旧的分数的位置一样不会变化的话就可以复用之前的旧结点
	// 那么就只需要更新结点的分数即可
	if ((oldNode.Backward.Score < newScore || oldNode.Backward.Data.LessThan(data)) && oldNode.Backward != this.Head || oldNode.Backward == this.Head) &&
		((newScore < oldNode.Forward[0].Score || data.LessThan(oldNode.Forward[0].Data)) && oldNode.Forward[0] != this.Tail || oldNode.Forward[0] == this.Tail) {
		oldNode.Score = newScore
		oldNode.Data = data
		return oldNode, true
	}
	
	// 不能复用的话，那就需要删除结点
	oldNode, ok = this.Delete(oldScore, data)
	if !ok {
		return oldNode, ok
	}
	// 然后创建新的结点，并插入到跳跃表中
	return this.Insert(newScore, data)
}

// 根据排名查找指定元素 排名从1开始
func (this *SkipList) GetRank(rank int) (*Node, bool) {
	assert.Assert(rank > 0, "排名从1开始:", rank)
	prev := this.Head
	span := 0
	for i := this.CurrentMaxLevel - 1; i >= 0; i-- {
		current := prev.Forward[i]
		for rank > span+current.Span && current != this.Tail {
			span += current.Span
			
			prev = current
			current = current.Forward[i]
		}
		
		if rank == span+current.Span && current != this.Tail {
			// 找到了
			return current, true
		}
		// 进入更低层级查找
	}
	return nil, false
}