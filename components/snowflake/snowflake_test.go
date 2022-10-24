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
	id := Get().Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())
	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())
	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())
	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())
	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", Get().Generate().Int64())
}
