package utils

import (
	"strings"
)

var subscriptMap = map[rune]rune{
	'0': '₀',
	'1': '₁',
	'2': '₂',
	'3': '₃',
	'4': '₄',
	'5': '₅',
	'6': '₆',
	'7': '₇',
	'8': '₈',
	'9': '₉',
}

// ToSubscript converte una stringa numerica nei corrispondenti caratteri pedice Unicode.
// Esempio: "16" → "₁₆"
func ToSubscript(input string) string {
	var builder strings.Builder
	for _, ch := range input {
		if sub, ok := subscriptMap[ch]; ok {
			builder.WriteRune(sub)
		} else {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}
