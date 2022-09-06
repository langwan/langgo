package log

import "testing"

func TestLogger(t *testing.T) {
	Logger("app", "test").Info().Msg("ok")
}
