package main

import "testing"

func Test_submit(t *testing.T) {
	type args struct {
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
				line: "1 3 4",
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				line: "3 2 3",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := submit(tt.args.line); got != tt.want {
				t.Errorf("submit() = %v, want %v", got, tt.want)
			}
		})
	}
}
