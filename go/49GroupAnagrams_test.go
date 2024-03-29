package main

import (
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		a        []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
		{[]string{"", "b"}, [][]string{{""}, {"b"}}},
		{[]string{"", ""}, [][]string{{"", ""}}},
	}

	for _, tt := range tests {
		res := groupAnagrams(tt.a)
		if len(res) != len(tt.expected) {
			errorMessage(t, tt.a, tt.expected, res)
			continue
		}
		for _, val := range res {
			found := false
			for _, val2 := range tt.expected {
				if slices.Equal(val2, val) {
					found = true
				}
			}
			if !found {
				errorMessage(t, tt.a, tt.expected, res)
			}
		}
	}
}

func errorMessage(t *testing.T, a []string, expected [][]string, res [][]string) {
	t.Errorf("groupAnagram(%v) = %v; but got %v", a, expected, res)
}
