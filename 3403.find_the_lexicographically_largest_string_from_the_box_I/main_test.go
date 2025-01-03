package main

import (
	"testing"
)

func TestAnswerString(t *testing.T) {
	tests := []struct {
		name       string
		word       string
		numFriends int
		want       string
	}{{
		name:       "example1",
		word:       "annabelle",
		numFriends: 2,
		want:       "nnabelle",
	}, {
		name:       "example2",
		word:       "dbca",
		numFriends: 2,
		want:       "dbc",
	}, {
		name:       "example3",
		word:       "bif",
		numFriends: 2,
		want:       "if",
	}, {
		name:       "example4",
		word:       "aann",
		numFriends: 2,
		want:       "nn",
	}, {
		name:       "example5",
		word:       "gh",
		numFriends: 1,
		want:       "gh",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := answerString(tt.word, tt.numFriends)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
