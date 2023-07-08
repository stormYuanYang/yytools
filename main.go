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
	"yytools/common/assert"
	"yytools/datastructure/sorted_set"
	"yytools/datastructure/stack"
)

var commandsMap = map[string]int{}

var commands = []string{
	"sorted_set",
	"stack",
}

var handlers = []func(int){
	sorted_set.SortedSetTest,
	stack.StackTest,
}

func init() {
	assert.Assert(len(commands) == len(handlers), "len of commands must equal to handlers")
	for i, str := range commands {
		commandsMap[str] = i
	}
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
		println("使用参考：yytools sorted_set 5\n表示执行5轮sorted_set相关测试代码")
		println("已支持的测试:")
		for _, str := range commands {
			println(str)
		}
		return
	}
	index, ok := commandsMap[command]
	if !ok {
		println("不支持的命令")
		return
	}
	
	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Command: %s Error: %+v\n", command, err)
		return
	}
	
	handler := handlers[index]
	handler(num)
}