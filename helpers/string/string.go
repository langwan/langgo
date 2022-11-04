package helper_string

import (
	"github.com/aquilax/truncate"
	"strings"
	"unicode/utf8"
)

func IsEmpty(val string) bool {
	s := strings.TrimSpace(val)
	return len(s) == 0
}

func Utf8StringLength(str string) int {
	return utf8.RuneCountInString(str)
}

func Utf8TruncateText(text string, max int, omission string) string {
	return truncate.Truncate(text, max, omission, truncate.PositionEnd)
}
