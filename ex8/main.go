package main

import "strings"

func normalize(phone string) string {
	var buf strings.Builder
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
