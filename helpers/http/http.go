package helper_http

import (
	"fmt"
)

// GenRange generate http range header
func GenRange(start, size int64) string {
	return fmt.Sprintf("bytes=%d-%d", start, start+size-1)
}
