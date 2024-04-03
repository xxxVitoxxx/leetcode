package main

import (
	"reflect"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	tests := []struct {
		name       string
		products   []string
		searchWord string
		output     [][]string
	}{
		{
			"example1",
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
			[][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}},
		},
		{
			"example2",
			[]string{"havana"},
			"havana",
			[][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products := make([]string, len(tt.products))
			copy(products, tt.products)

			res := suggestedProducts(products, tt.searchWord)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products := make([]string, len(tt.products))
			copy(products, tt.products)

			res := suggestedProducts2(products, tt.searchWord)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products := make([]string, len(tt.products))
			copy(products, tt.products)

			res := suggestedProducts3(products, tt.searchWord)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
