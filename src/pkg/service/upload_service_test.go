package service

import "testing"

func TestUploadAliBillPrint(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		{"01", args{"hello"}},
		{"02", args{"你好"}},
		{"03", args{"世界"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OnUploadAliBillPrint(tt.args.str)
		})
	}
}
