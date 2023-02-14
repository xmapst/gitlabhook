package utils

import (
	"fmt"
	"github.com/go-cmd/cmd"
)

func Exec(name string, args ...string) (output []string, err error) {
	exec := cmd.NewCmd(name, args...)
	status := <-exec.Start()
	if status.Error != nil {
		return nil, err
	}
	if status.Exit != 0 {
		return nil, fmt.Errorf("exit code is %d", status.Exit)
	}
	if len(status.Stderr) != 0 {
		return nil, fmt.Errorf("%s", status.Stderr)
	}
	return status.Stdout, nil
}
