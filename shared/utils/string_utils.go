package utils

import (
	"regexp"
	"strings"
)

func NormalizeString(input string) string {
	re := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	normalized := re.ReplaceAllString(input, "-")

	return strings.Trim(normalized, "-")
}
