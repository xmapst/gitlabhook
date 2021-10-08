//go:build !windows

package failed

import (
	"fmt"
	"os"
)

func Exit(code int, msg string) {
    _, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf(`GL-HOOK-ERR: exit code %d`, code))
	_, _ = fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
