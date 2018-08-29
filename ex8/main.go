package main

import (
	"regexp"
	"strings"
)

func normalizeItr(phone string) string {
	var buf strings.Builder
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func normalizeRegex(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}
