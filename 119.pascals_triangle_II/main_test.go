package main

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{{
		name:     "example1",
		rowIndex: 3,
		want:     []int{1, 3, 3, 1},
	}, {
		name:     "example2",
		rowIndex: 4,
		want:     []int{1, 4, 6, 4, 1},
	}, {
		name:     "exampl3",
		rowIndex: 0,
		want:     []int{1},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRow(tt.rowIndex)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRow2(tt.rowIndex)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRow3(tt.rowIndex)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
