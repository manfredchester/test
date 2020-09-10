package main

import (
	"testing"
)

func Test_zheng(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "数字",
			args: args{
				str: " ",
			},
		},
		{
			name: "小写",
			args: args{
				str: "q@erewe",
			},
		},
		{
			name: "大写",
			args: args{
				str: "QWE:QW",
			},
		},
		{
			name: "特殊",
			args: args{
				str: "w、",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zheng(tt.args.str)
		})
	}
}
