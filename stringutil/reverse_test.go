package stringutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Reverse: len even",
			args: args{"abcd"},
			want: "dcba",
		},
		{
			name: "Test Reverse: len odd",
			args: args{"abcde"},
			want: "edcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseNew(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test ReverseNew: len even",
			args: args{"abcd"},
			want: "dcba",
		},
		{
			name: "Test ReverseNew: len odd",
			args: args{"abcde"},
			want: "edcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseNew(tt.args.s); got != tt.want {
				t.Errorf("ReverseNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
