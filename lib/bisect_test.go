package lib

import (
	"testing"
)

func Test_bisect(t *testing.T) {
	type args struct {
		arr []int
		v   int
		l   int
		r   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				v:   5,
				l:   0,
				r:   9,
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				v:   4,
				l:   0,
				r:   9,
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				v:   3,
				l:   0,
				r:   9,
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				v:   0,
				l:   0,
				r:   9,
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				v:   9,
				l:   0,
				r:   9,
			},
			want: 9,
		},
		{
			name: "6",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				v:   5,
				l:   3,
				r:   8,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bisect(tt.args.arr, tt.args.v, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("bisect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bisectLeft(t *testing.T) {
	type args struct {
		arr []int
		v   int
		l   int
		r   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   1,
				l:   0,
				r:   9,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   0,
				l:   0,
				r:   9,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   100,
				l:   0,
				r:   9,
			},
			want: 9,
		},
		{
			name: "4",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   15,
				l:   0,
				r:   9,
			},
			want: 5,
		},
		{
			name: "5",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   30,
				l:   0,
				r:   9,
			},
			want: 6,
		},
		{
			name: "6",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   -10,
				l:   0,
				r:   9,
			},
			want: -1,
		},
		{
			name: "7",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   15,
				l:   3,
				r:   8,
			},
			want: 5,
		},
		{
			name: "8",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   15,
				l:   6,
				r:   8,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bisectLeft(tt.args.arr, tt.args.v, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("bisectLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bisectRight(t *testing.T) {
	type args struct {
		arr []int
		v   int
		l   int
		r   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   -10,
				l:   0,
				r:   9,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   0,
				l:   0,
				r:   9,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   88,
				l:   0,
				r:   9,
			},
			want: 9,
		},
		{
			name: "4",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   11,
				l:   0,
				r:   9,
			},
			want: 5,
		},
		{
			name: "5",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   15,
				l:   0,
				r:   9,
			},
			want: 6,
		},
		{
			name: "6",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   100,
				l:   0,
				r:   9,
			},
			want: -1,
		},
		{
			name: "7",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   15,
				l:   3,
				r:   8,
			},
			want: 6,
		},
		{
			name: "8",
			args: args{
				arr: []int{0, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				v:   60,
				l:   6,
				r:   8,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bisectRight(tt.args.arr, tt.args.v, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("bisectRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
