package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
)

func uppercaseFirstCharacter(str string) string {
	languageSpecificTitleCasing := cases.Title(language.English)
	return languageSpecificTitleCasing.String(str)
}

func formatNumberWithSpaces(n int) string {
	s := strconv.Itoa(n)
	var result strings.Builder

	start := len(s) % 3
	if start == 0 && len(s) > 3 {
		start = 3
	}

	result.WriteString(s[:start])

	for i := start; i < len(s); i += 3 {
		if i != 0 {
			result.WriteByte(' ')
		}
		result.WriteString(s[i : i+3])
	}

	return result.String()
}
