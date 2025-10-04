package utils

import (
	"regexp"
	"strings"
)

func CleanLLMOutput (raw string) string {
	re := regexp.MustCompile("(?s)```json(.*?)```");
	cleaned := re.ReplaceAllString(raw, "$1");

	cleaned = strings.Trim(cleaned, "` \n\t")

	return cleaned;
}