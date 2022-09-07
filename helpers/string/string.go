package helperString

import "strings"

func IsEmpty(val string) bool {
	s := strings.TrimSpace(val)
	return len(s) == 0
}
