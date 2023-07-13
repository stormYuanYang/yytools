// Package main.

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
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"yytools/algorithm/math"
	"yytools/algorithm/math/probability_distribution"
	"yytools/common/assert"
	"yytools/datastructure/heap"
	"yytools/datastructure/queue"
	"yytools/datastructure/sorted_set"
	"yytools/datastructure/stack"
)

var commandsMap = map[string]int{}

var commands = []string{
	"heap",
	"mathcommon",
	"maxheap",
	"prob",
	"pq",
	"queue",
	"sortedset",
	"stack",
}

var notes = []string{
	"最小堆",
	"公共数学方法（比如gcd）",
	"最大堆",
	"概率分布",
	"优先级队列",
	"队列",
	"有序集合",
	"栈",
}

var handlers = []func(int){
	heap.HeapTest,
	math.MathCommonTest,
	heap.MaxHeapTest,
	probability_distribution.ProbabilityDistributionTest,
	heap.PriorityQueueTest,
	queue.QueueTest,
	sorted_set.SortedSetTest,
	stack.StackTest,
}

func init() {
	assert.Assert(len(commands) == len(handlers), "len of commands must equal to handlers")
	assert.Assert(len(commands) == len(notes), "len of commands must equal to notes")
	for i, str := range commands {
		commandsMap[str] = i
	}
}

func testAll(num int) {
	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]
		handler(num)
	}
	println("\n所有测试完毕...")
}

func main() {
	// 第一个参数是可执行文件本身的路径
	// 后续的参数是通过控制台传递的参数
	args := os.Args[1:]
	if len(args) == 0 {
		println("需要传入参数,如需要帮助, 请使用: yytools help")
		return
	}
	command := strings.ToLower(args[0])
	if command == "help" {
		println("使用参考：yytools sorted_set 5\n表示执行5轮sorted_set相关测试代码\n yytools all 5\n表示对所有测试进行5轮测试\n")
		println("已支持的命令:")
		fmt.Printf("%-20s\t说明:执行所有命令\n", "all")
		for i := 0; i < len(commands); i++ {
			fmt.Printf("%-20s\t说明:%s\n", commands[i], notes[i])
		}
		return
	}
	
	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Command: %s Error: %+v\n", command, err)
		return
	}
	
	assert.SetAssert(true)
	
	if command == "all" {
		testAll(num)
	} else {
		index, ok := commandsMap[command]
		if !ok {
			println("不支持的命令")
			return
		}
		handler := handlers[index]
		handler(num)
	}
	
}