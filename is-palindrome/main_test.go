package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	t.Parallel()

	type args struct {
		word string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "English palindrome word",
			args: args{
				word: "repaper",
			},
			want: true,
		},
		{
			name: "English short palindrome word",
			args: args{
				word: "wow",
			},
			want: true,
		},
		{
			name: "Russian palindrome word",
			args: args{
				word: "шалаш",
			},
			want: true,
		},
		{
			name: "English non palindrome word",
			args: args{
				word: "check",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isPalindrome(tt.args.word); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
