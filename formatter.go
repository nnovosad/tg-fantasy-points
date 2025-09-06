package main

import (
	"fmt"
	"strings"
)

func formatTelegramMessage(league, tourInfo, matchesInfo, tourResults, seasonInfo string) string {
	var message strings.Builder

	message.WriteString(fmt.Sprintf("🏆 **%s**\n\n", league))

	message.WriteString(fmt.Sprintf("📅 **Тур:** %s\n", tourInfo))
	message.WriteString(fmt.Sprintf("📊 **Статус:** %s\n\n", tourResults))

	message.WriteString("⚽ **МАТЧИ:**\n")
	message.WriteString("```\n")
	message.WriteString(matchesInfo)
	message.WriteString("```\n\n")

	message.WriteString("📈 **СТАТИСТИКА СЕЗОНА:**\n")
	message.WriteString(seasonInfo)

	return message.String()
}

func formatMatchesInfo(matchesInfo string) string {
	lines := strings.Split(matchesInfo, "\n")
	var formatted strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Contains(line, ".") && (strings.Contains(line, "завершен") || strings.Contains(line, "первый тайм") || strings.Contains(line, "второй тайм") || strings.Contains(line, "перерыв") || strings.Contains(line, "– : –")) {
			if strings.Contains(line, "завершен") {
				line = strings.Replace(line, "завершен", "✅ завершен", 1)
			} else if strings.Contains(line, "первый тайм") {
				line = strings.Replace(line, "первый тайм", "🟡 первый тайм", 1)
			} else if strings.Contains(line, "второй тайм") {
				line = strings.Replace(line, "второй тайм", "🟡 второй тайм", 1)
			} else if strings.Contains(line, "перерыв") {
				line = strings.Replace(line, "перерыв", "⏸ перерыв", 1)
			} else if strings.Contains(line, "– : –") {
				line = strings.Replace(line, "– : –", "⏰ – : –", 1)
			}
			formatted.WriteString(line + "\n")
		} else if strings.HasPrefix(line, "---") {
			// Информация об игроках
			formatted.WriteString("  " + line + "\n")
		}
	}

	return formatted.String()
}

func formatPlayerInfo(playerInfo string) string {
	lines := strings.Split(playerInfo, "\n")
	var formatted strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "---") {
			if strings.Contains(line, "Goalkeeper") {
				line = strings.Replace(line, "Goalkeeper", "🥅 Goalkeeper", 1)
			} else if strings.Contains(line, "Defender") {
				line = strings.Replace(line, "Defender", "🛡 Defender", 1)
			} else if strings.Contains(line, "Midfielder") {
				line = strings.Replace(line, "Midfielder", "⚽ Midfielder", 1)
			} else if strings.Contains(line, "Forward") {
				line = strings.Replace(line, "Forward", "🎯 Forward", 1)
			}

			if strings.Contains(line, "Main cast") {
				line = strings.Replace(line, "Main cast", "⭐ Main cast", 1)
			} else if strings.Contains(line, "On the bench") {
				line = strings.Replace(line, "On the bench", "🪑 On the bench", 1)
			}

			if strings.Contains(line, "scored") {
				parts := strings.Split(line, "scored")
				if len(parts) == 2 {
					scorePart := strings.Split(parts[1], "points")
					if len(scorePart) >= 2 {
						score := strings.TrimSpace(scorePart[0])
						line = strings.Replace(line, "scored "+score+" points", fmt.Sprintf("scored **%s** points", score), 1)
					}
				}
			}

			formatted.WriteString("  " + line + "\n")
		}
	}

	return formatted.String()
}

func formatSeasonInfo(seasonInfo string) string {
	seasonInfo = strings.Replace(seasonInfo, "You scored", "🎯 Вы набрали", 1)
	seasonInfo = strings.Replace(seasonInfo, "points in the season", "очков в сезоне", 1)
	seasonInfo = strings.Replace(seasonInfo, "and are in", "и занимаете", 1)
	seasonInfo = strings.Replace(seasonInfo, "rd place out of", " место из", 1)

	parts := strings.Split(seasonInfo, "Rank:")
	if len(parts) == 2 {
		rankPart := strings.TrimSpace(parts[1])
		rankValue := extractRankValue(rankPart)

		rankType := getRankType(rankValue)

		seasonInfo = parts[0] + "\n" + rankType + ": " + rankValue
	}

	return seasonInfo
}

func extractRankValue(rankStr string) string {
	// Убираем лишние пробелы и возвращаем чистую строку
	return strings.TrimSpace(rankStr)
}

func getRankType(rankValue string) string {
	var rank int
	if _, err := fmt.Sscanf(rankValue, "%d", &rank); err != nil {
		return "🏆 Ранг"
	}

	if rank >= 99 {
		return "🥇 Топ-1"
	} else if rank >= 97 {
		return "🥉 Топ-3"
	} else if rank >= 95 {
		return "🏅 Топ-5"
	} else if rank >= 90 {
		return "🏆 Топ-10"
	} else {
		return "🏆 Ранг"
	}
}

func formatTourInfo(tourInfo string) string {
	if strings.Contains(tourInfo, "Opened") {
		tourInfo = strings.Replace(tourInfo, "Opened", "🟢 Открыт", 1)
	} else if strings.Contains(tourInfo, "Closed") {
		tourInfo = strings.Replace(tourInfo, "Closed", "🔴 Закрыт", 1)
	}

	return tourInfo
}

func formatTourResults(tourResults string) string {
	if strings.Contains(tourResults, "open") {
		tourResults = strings.Replace(tourResults, "The tour is open", "🟢 Тур открыт", 1)
	} else {
		tourResults = strings.Replace(tourResults, "You scored", "🎯 Вы набрали", 1)
		tourResults = strings.Replace(tourResults, "points", "очков", 1)
		tourResults = strings.Replace(tourResults, "Average score", "Средний балл", 1)
	}

	return tourResults
}
