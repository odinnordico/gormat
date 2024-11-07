package format

import (
	"strings"
	"unicode"
)

func CleanString(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) && unicode.IsPrint(r) {
			return r
		}
		return -1
	}, strings.TrimSpace(s))
}
