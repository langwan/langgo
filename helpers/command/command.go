package helper_command

import (
	"context"
	"os/exec"
)

func Execute(ctx context.Context, name string, args ...string) (output []byte, err error) {
	cmd := exec.CommandContext(ctx, name, args...)
	output, err = cmd.Output()
	return output, err
}
