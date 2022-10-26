package snowflake

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"testing"
)

func TestRun(t *testing.T) {
	core.EnvName = core.Development
	core.LoadConfigurationFile("../../testdata/configuration_test.app.yml")
	langgo.Run(&Instance{})
	fmt.Printf("Int64  ID: %d\n", Get().machineID)
	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", Gen())
	fmt.Printf("Int64  ID: %d\n", Gen())
	fmt.Printf("Int64  ID: %d\n", Gen())
	fmt.Printf("Int64  ID: %d\n", Gen())
	fmt.Printf("Int64  ID: %d\n", Gen())
}
