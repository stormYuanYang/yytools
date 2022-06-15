// Package tools.
// 版权所有[成都创人所爱科技股份有限公司]
// 根据《保密信息使用许可证》获得许可;
// 除非符合许可，否则您不得使用此文件。
// 您可以在以下位置获取许可证副本，链接地址：
// https://wiki.tap4fun.com/display/MO/Confidentiality
// 除非适用法律要求或书面同意，否则保密信息按照使用许可证要求使用，不附带任何明示或暗示的保证或条件。
// 有关管理权限的特定语言，请参阅许可证副本。

// 作者:  yangyuan
// 创建日期:2022/6/15
package tools

import (
	"math"
	"testing"
)

func TestRandInt32(t *testing.T) {
	type args struct {
		low  int32
		high int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "边界情况测试1",
			args: args{
				low:  0,
				high: 0,
			},
			want: 0,
		},
		{
			name: "边界情况测试2",
			args: args{
				low:  math.MaxInt32,
				high: math.MaxInt32,
			},
			want: math.MaxInt32,
		},
		//{
		//	name: "测试3",
		//	args: args{
		//		low:  -1,
		//		high: math.MaxInt32,
		//	},
		//	want: math.MaxInt32,
		//},
		{
			name: "测试4",
			args: args{
				low:  1,
				high: math.MaxInt32,
			},
			want: 0,
		},
		{
			name: "测试5",
			args: args{
				low:  0,
				high: math.MaxInt32,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandInt32(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("RandInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandInt64(t *testing.T) {
	type args struct {
		low  int64
		high int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandInt64(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("RandInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}