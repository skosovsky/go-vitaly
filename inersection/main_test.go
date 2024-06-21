package main

import (
	"reflect"
	"testing"
)

func Test_intersectionBrutForce(t *testing.T) {
	t.Parallel()

	type args struct {
		setA []int
		setB []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "With 2 intersection, but same lengths",
			args: args{setA: []int{1, 2, 3, 23}, setB: []int{2, 4, 6, 23}},
			want: []int{2, 23},
		},
		{
			name: "With 3 intersection, but different lengths and same values",
			args: args{setA: []int{1, 1, 1}, setB: []int{1, 1, 1, 1}},
			want: []int{1, 1, 1},
		},
		{
			name: "With 3 intersection, but different lengths and same values",
			args: args{setA: []int{1, 1, 1, 1}, setB: []int{1, 1, 1}},
			want: []int{1, 1, 1},
		},
		{
			name: "With set from benchmark",
			args: args{setA: []int{1, 2, 3, 23, 32, 34, 38, 40}, setB: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := intersectionBruteforce(tt.args.setA, tt.args.setB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectionBrutForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersectionTwoPointer(t *testing.T) {
	t.Parallel()

	type args struct {
		setA []int
		setB []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "With 2 intersection, but same lengths",
			args: args{setA: []int{1, 2, 3, 23}, setB: []int{2, 4, 6, 23}},
			want: []int{2, 23},
		},
		{
			name: "With 3 intersection, but different lengths and same values",
			args: args{setA: []int{1, 1, 1}, setB: []int{1, 1, 1, 1}},
			want: []int{1, 1, 1},
		},
		{
			name: "With 3 intersection, but different lengths and same values",
			args: args{setA: []int{1, 1, 1, 1}, setB: []int{1, 1, 1}},
			want: []int{1, 1, 1},
		},
		{
			name: "With set from benchmark",
			args: args{setA: []int{1, 2, 3, 23, 32, 34, 38, 40}, setB: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := intersectionTwoPointer(tt.args.setA, tt.args.setB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectionTwoPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_intersectionBruteforce(b *testing.B) {
	setA := []int{1, 2, 3, 23, 32, 34, 38, 40}
	setB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	for i := 0; i < b.N; i++ {
		intersectionBruteforce(setA, setB)
	}
}

func Benchmark_intersectionTwoPointer(b *testing.B) {
	setA := []int{1, 2, 3, 23, 32, 34, 38, 40}
	setB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	for i := 0; i < b.N; i++ {
		intersectionTwoPointer(setA, setB)
	}
}
