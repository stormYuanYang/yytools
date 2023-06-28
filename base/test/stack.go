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
// 创建日期:2023/6/7
package test

import (
	"strconv"
	"yytools/base/algorithm/math"
	"yytools/base/common/assert"
	"yytools/base/common/constant"
	"yytools/base/datastructure/stack"
)

func IntStackTest0() {
	s := stack.NewStack()
	for i := 1; i < 10; i++ {
		s.Push(i)
	}
	for !s.Empty() {
		t := s.Top().(int)
		print(t, ",")
		p := s.Pop().(int)
		print(p, ";")
	}
	println()
}

func IntStackTest1() {
	id := 0
	
	for x := 1; x <= 10; x++ {
		s := stack.NewStack()
		for i := 1; i < 10000; i++ {
			id++
			r := math.RandInt32(1, constant.TEN_THOUSAND)
			if r < constant.TEN_THOUSAND/2 {
				s.Push(id)
			} else {
				if !s.Empty() {
					s.Pop()
				}
			}
			for j := 1; j < s.Len()-1; j++ {
				// 不管怎么入栈和出栈,留在栈里的整数肯定是递增的
				assert.Assert(s.Items[j].(int) < s.Items[j+1].(int), "必须保证栈的先进后出性质!", s.Items[j], s.Items[j+1])
			}
		}
		println("verify stack ok:", x)
	}
}

type Student struct {
	Name string
	Age  int
}

func StructStackTest0() {
	s := stack.NewStack()
	for i := 1; i < 10; i++ {
		student := &Student{
			Name: strconv.Itoa(i),
			Age:  i,
		}
		s.Push(student)
	}
	for !s.Empty() {
		t := s.Top().(*Student)
		print(t.Name, ",")
		p := s.Pop().(*Student)
		print(p.Name, "; ")
	}
	println()
}