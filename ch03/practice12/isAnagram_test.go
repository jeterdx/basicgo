package main

import (
	"strings"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s1: "abcd", s2: "abcd"}, want: true},
		{args: args{s1: "abcd", s2: "abdc"}, want: true},
		{args: args{s1: "abcd", s2: "acbd"}, want: true},
		{args: args{s1: "abcd", s2: "acdb"}, want: true},
		{args: args{s1: "abcd", s2: "adbc"}, want: true},
		{args: args{s1: "abcd", s2: "adcb"}, want: true},
		{args: args{s1: "abcd", s2: "bacd"}, want: true},
		{args: args{s1: "abcd", s2: "badc"}, want: true},
		{args: args{s1: "abcd", s2: "bcad"}, want: true},
		{args: args{s1: "abcd", s2: "bcda"}, want: true},
		{args: args{s1: "abcd", s2: "bdac"}, want: true},
		{args: args{s1: "abcd", s2: "bdca"}, want: true},
		{args: args{s1: "abcd", s2: "cabd"}, want: true},
		{args: args{s1: "abcd", s2: "cadb"}, want: true},
		{args: args{s1: "abcd", s2: "cbad"}, want: true},
		{args: args{s1: "abcd", s2: "cbda"}, want: true},
		{args: args{s1: "abcd", s2: "cdab"}, want: true},
		{args: args{s1: "abcd", s2: "cdba"}, want: true},
		{args: args{s1: "abcd", s2: "dabc"}, want: true},
		{args: args{s1: "abcd", s2: "dacb"}, want: true},
		{args: args{s1: "abcd", s2: "dbac"}, want: true},
		{args: args{s1: "abcd", s2: "dbca"}, want: true},
		{args: args{s1: "abcd", s2: "dcab"}, want: true},
		{args: args{s1: "abcd", s2: "dcba"}, want: true},
		{args: args{s1: "あえ", s2: "あえ	"}, want: true},
		{args: args{s1: "あいうえ", s2: "えういzz"}, want: false},
		{args: args{s1: "abcd", s2: "dcbx"}, want: false},
		{args: args{s1: "abcd", s2: "dcb"}, want: false},
		{args: args{s1: "aaab", s2: "bbba"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

//アナグラムは同じ文字が同じ数存在するかを判定すること
//文字列を先頭から比較して、同じ文字が存在した場合、その文字同士をピリオドに置き換える、という処理を繰り返して最後に2つのstringがイコールになっていたらtrue、って風に書いた。
//置き換える文字を""で実行したかったが、for loopの最中に文字列のサイズを小さくしてしまうと、indexの数が合わなくなってエラーに遭遇したので。。
func isAnagram(s1 string, s2 string) bool {
	for i1, v1 := range s1 {
		for i2, v2 := range s2 {
			if string(v1) == "" || string(v2) == "" {
				return false
			}
			if string(v1) == string(v2) {
				s1 = strings.Replace(s1, string(s1[i1]), ".", 1)
				s2 = strings.Replace(s2, string(s2[i2]), ".", 1)
			}
		}
	}
	return s1 == s2
}
