package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var s strings.Builder
	for _, str := range strs {
		length := len(str)
		s.WriteString(strconv.Itoa(length))
		s.WriteByte('#')
		s.WriteString(str)
	}
	return s.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var result []string
	var i int
	for i < len(strs) {
		j := i
		for strs[j] != '#' {
			j++
		}

		length, err := strconv.Atoi(strs[i:j])
		if err != nil {
			return []string{}
		}

		j++
		result = append(result, strs[j:j+length])

		i = j + length
	}
	return result
}
