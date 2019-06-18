package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				i: 0,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				i: 3,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				i: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.i); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
