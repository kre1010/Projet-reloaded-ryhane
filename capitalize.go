package main

import (
	"strings"
)

func ModifierCasse(phrase, node string) string {
	switch node {
	case "low":
		return strings.ToLower(phrase)
	case "up":
		return strings.ToUpper(phrase)
	case "cap":
		return strings.ToTitle(phrase)
	default:
		return phrase
	}
}
