package main

import "testing"

func Test_upperBound(t *testing.T) {
	type args struct {
		a []int
		l int
		r int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				l: 0,
				r: 8,
				v: 2,
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				a: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				l: 0,
				r: 8,
				v: 0,
			},
			want: 8,
		},
		{
			name: "3",
			args: args{
				a: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				l: 0,
				r: 8,
				v: 8,
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				a: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				l: 1,
				r: 7,
				v: 5,
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				a: []int{9, 9, 9, 6, 5, 4, 3, 2, 1},
				l: 0,
				r: 8,
				v: 6,
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				a: []int{9, 9, 9, 9, 5, 4, 3, 2, 1},
				l: 0,
				r: 8,
				v: 5,
			},
			want: 3,
		},
		{
			name: "7",
			args: args{
				a: []int{9, 9, 9, 9, 9, 4, 3, 2, 1},
				l: 0,
				r: 8,
				v: 4,
			},
			want: 4,
		},
		{
			name: "8",
			args: args{
				a: []int{9, 9, 9, 9, 9, 9, 3, 2, 1},
				l: 0,
				r: 8,
				v: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upperBound(tt.args.a, tt.args.l, tt.args.r, tt.args.v); got != tt.want {
				t.Errorf("upperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
