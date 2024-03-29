package main

import (
	"slices"
	"testing"
)

func TestTwoSumReal(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 3, 4}, 7, []int{1, 2}},
		{[]int{-1, 0}, -1, []int{0, 1}},
	}

	for _, tt := range tests {
		res := twoSumReal(tt.a, tt.b)
		if !slices.Equal(res, tt.expected) {
			t.Errorf("twoSumOne(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}

	for _, tt := range tests {
		res := twoSumOne(tt.a, tt.b)
		if !slices.Equal(res, tt.expected) {
			t.Errorf("twoSumOne(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
