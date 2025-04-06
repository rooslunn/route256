package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		rates rates
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1#1", args{
			rates{
				{
					{100, 1},
					{100, 1},
					{1, 100},
					{3, 2},
					{1, 100},
					{2, 3},
				}, {
					{100, 1},
					{100, 1},
					{1, 100},
					{3, 2},
					{1, 100},
					{2, 3},
				}, {
					{100, 1},
					{100, 1},
					{1, 100},
					{3, 2},
					{1, 100},
					{2, 3},
				},
			},
		}, 0.015},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answer(tt.args.rates); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
