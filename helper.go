package main

import (
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func prepareRank(currentPlace int, totalPlace int) int {
	return int((1.0 - float64(currentPlace)/float64(totalPlace)) * 100)
}

func isPlayerTeam(match string, playerTeamName string) bool {
	parts := strings.Split(match, ":")
	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])

	re := regexp.MustCompile(`\d+`)
	leftTeam := strings.TrimSpace(re.ReplaceAllString(left, ""))
	rightTeam := strings.TrimSpace(re.ReplaceAllString(right, ""))

	return leftTeam == playerTeamName || rightTeam == playerTeamName
}
