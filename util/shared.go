package util

import (
	"strings"
)

// return s with items lowercased and without leading/trailing commas
func Clean(s []string) []string {
	var cleaned []string
	for _, v := range s {
		cleaned = append(cleaned, strings.ToLower(strings.Trim(v, ",")))
	}
	return cleaned
}
