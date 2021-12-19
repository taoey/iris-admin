package util

import (
	"fmt"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	password := "123456"
	sysPassword, _ := GeneratePassword(password)
	fmt.Println(sysPassword)
}

func TestComparePassword(t *testing.T) {
	type args struct {
		sysPassword  string
		userPassword string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"密码正确用例",
			args{`$2a$10$ZPoBPKa3sC/CJzjQRodJLeP0KZadumvlVXRUIyW7rFwwROhpjOKWa`, `123456`},
			true,
		},
		{
			"密码错误用例",
			args{`$2a$10$ZPoBPKa3sC/CJzjQRodJLeP0KZadumvlVXRUIyW7rFwwROhpjOKWa`, `1234567`},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePassword(tt.args.sysPassword, tt.args.userPassword); got != tt.want {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesEncrypt(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"case 1",
			args{"123456"},
			"Fh43pURb0IK/iyXSfOA2ug==",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesEncrypt(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AesEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesDecrypt(t *testing.T) {
	type args struct {
		crypted string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"case 1",
			args{"Fh43pURb0IK/iyXSfOA2ug=="},
			"123456",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesDecrypt(tt.args.crypted)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AesDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
