package jsonassert

import (
	"testing"
)

func TestIsRegEx(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "ok",
			str:  "<<<^ok$>>>",
			want: true,
		},
		{
			name: "ok",
			str:  "<<<hacker>>>",
			want: true,
		},
		{
			name: "ok",
			str:  "<<<this is reg ex words>>>",
			want: true,
		},
		{
			name: "ok",
			str:  "<<< t h i s is reg ex words >>>",
			want: true,
		},
		{
			name: "fail",
			str:  " <<<this is reg ex words>>>",
			want: false,
		},
		{
			name: "fail",
			str:  "<<this is reg ex words>>>",
			want: false,
		},
		{
			name: "fail",
			str:  "<<this is reg ex words>>",
			want: false,
		},
		{
			name: "fail",
			str:  "<<<this is reg ex words>>> ",
			want: false,
		},
		{
			name: "fail",
			str:  " <<<this is reg ex words>>> ",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := isRegEx(tt.str); err != nil || got != tt.want {
				t.Errorf("isRegEx(%v) = %v, %v, want %v", tt.str, got, err, tt.want)
			}
		})
	}
}

func TestGetReqExPattern(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "ok",
			str:  "<<<^ok$>>>",
			want: "^ok$",
		},
		{
			name: "ok",
			str:  "<<<pattern>>>",
			want: "pattern",
		},
		{
			name: "ok",
			str:  "<<<^pattern$>>>",
			want: "^pattern$",
		},
		{
			name: "fail",
			str:  " <<<^pattern$>>> ",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReqExPattern(tt.str); got != tt.want {
				t.Errorf("getReqExPattern(%v) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}
