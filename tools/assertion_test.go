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

import "testing"

func TestAssert(t *testing.T) {
	type args struct {
		condition bool
		strList   []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		//{
		//	name: "字符串数组为空时",
		//	args: args{
		//		condition: false,
		//		strList:   nil,
		//	},
		//},
		//{
		//	name: "字符串数组长度为1",
		//	args: args{
		//		condition: false,
		//		strList:   []string{"hello"},
		//	},
		//},
		//{
		//	name: "字符串数组长度为2",
		//	args: args{
		//		condition: false,
		//		strList:   []string{"hello", "yytools"},
		//	},
		//},
		{
			name: "条件为真时",
			args: args{
				condition: true,
				strList:   []string{"hello", "yytools"},
			},
		},
		{
			name: "条件为真时 strList为空",
			args: args{
				condition: true,
				strList:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Assert(tt.args.condition, tt.args.strList...)
		})
	}
}