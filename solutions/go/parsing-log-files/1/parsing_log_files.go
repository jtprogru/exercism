package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	var prefixes = []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}

	for _, p := range prefixes {
		if strings.HasPrefix(text, p) {
			return true
		}
	}

	return false
}

func SplitLogLine(text string) []string {
	str := `<[*~=-]*>`
	re := regexp.MustCompile(str)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	str := `(?i)".*password.*"`
	regex := regexp.MustCompile(str)
	for _, line := range lines {
		if regex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex := regexp.MustCompile(`end-of-line\d+`)
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var result []string
	regex := regexp.MustCompile(`\sUser\s+(\S+)`)
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		if len(match) > 1 {
			line = "[USR] " + match[1] + " " + line
		}
		result = append(result, line)
	}
	return result
}
