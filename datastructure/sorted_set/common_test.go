// Package sorted_set.

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
package sorted_set

import "testing"

func Test_randomLevel(t *testing.T) {
	type args struct {
		maxLevel           int
		levelUpProbability float32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "测试1",
			args: args{
				maxLevel:           SKIPLIST_MAXLEVEL,
				levelUpProbability: DEFAULT_LEVELUP_PROBABILITY,
			},
			want: 1,
		},
		//{
		//	name: "测试2",
		//	args: args{
		//		maxLevel:           SKIPLIST_MAXLEVEL + 1,
		//		levelUpProbability: DEFAULT_LEVELUP_PROBABILITY,
		//	},
		//	want: 1,
		//},
		{
			name: "测试3",
			args: args{
				maxLevel:           SKIPLIST_MAXLEVEL,
				levelUpProbability: 1.1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomLevel(tt.args.levelUpProbability); got != tt.want {
				t.Errorf("randomLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}