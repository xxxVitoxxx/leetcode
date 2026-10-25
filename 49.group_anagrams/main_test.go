package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{{
		name:     "example1",
		strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	}, {
		name:     "example2",
		strs:     []string{""},
		expected: [][]string{{""}},
	}, {
		name:     "example3",
		strs:     []string{"a"},
		expected: [][]string{{"a"}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := groupAnagrams(tt.strs)

			normalize(tt.expected)
			normalize(actual)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := groupAnagrams2(tt.strs)

			normalize(tt.expected)
			normalize(actual)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}

func normalize(groups [][]string) {
	for i := range groups {
		sort.Strings(groups[i])
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
}
