package main

import "testing"

func Test_submit(t *testing.T) {
	type args struct {
		n    int
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n:    3,
				line: "1 2 3",
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				n:    4,
				line: "1 3 1 1",
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				n:    8,
				line: "27 23 76 2 3 5 62 52",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := submit(tt.args.n, tt.args.line); got != tt.want {
				t.Errorf("submit() = %v, want %v", got, tt.want)
			}
		})
	}
}
