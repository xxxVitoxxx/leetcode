package main

import (
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		output    []float64
	}{
		{
			"example1",
			[][]string{{"a", "b"}, {"b", "c"}},
			[]float64{2, 3},
			[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			[]float64{6, 0.5, -1, 1, -1},
		},
		{
			"example2",
			[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			[]float64{1.5, 2.5, 5},
			[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			[]float64{3.75, 0.4, 5, 0.2},
		},
		{
			"example3",
			[][]string{{"a", "b"}},
			[]float64{0.5},
			[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			[]float64{0.5, 2, -1, -1},
		},
		{
			"example4",
			[][]string{{"a", "b"}, {"c", "d"}},
			[]float64{2, 3},
			[][]string{{"a", "c"}, {"a", "d"}},
			[]float64{-1, -1},
		},
		{
			"example5",
			[][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}, {"d", "e"}},
			[]float64{2, 3, 5, 2},
			[][]string{{"a", "e"}, {"a", "d"}, {"b", "e"}, {"e", "c"}},
			[]float64{60, 30, 30, 0.1},
		},
		{
			"example6",
			[][]string{{"a", "b"}, {"a", "c"}},
			[]float64{2, 3},
			[][]string{{"a", "c"}, {"b", "c"}, {"c", "a"}},
			[]float64{3, 1.5, 0.3333333333333333},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := calcEquation(tt.equations, tt.values, tt.queries)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := calcEquation2(tt.equations, tt.values, tt.queries)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
