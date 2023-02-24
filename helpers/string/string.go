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

func SetSuffix(s string, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s
	} else {
		var builder strings.Builder
		builder.Grow(len(s) + len(suffix))
		builder.WriteString(s)
		builder.WriteString(suffix)
		return builder.String()
	}
}
func RemoveSuffix(s string, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		return s
	} else {
		return s[:len(s)-len(suffix)]
	}
}

func SetPrefix(s string, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s
	} else {
		var builder strings.Builder
		builder.Grow(len(s) + len(prefix))
		builder.WriteString(prefix)
		builder.WriteString(s)
		return builder.String()
	}
}

func RemovePrefix(s string, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return s
	} else {
		return s[len(prefix):]
	}
}
