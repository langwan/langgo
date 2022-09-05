package langgo

import (
	"github.com/langwan/langgo/components/jwt"
	"github.com/langwan/langgo/core"
	"testing"
)

func TestLoadComponents(t *testing.T) {
	Init()
	core.AddComponents(&jwt.Instance{})
	core.LoadComponents()
}
