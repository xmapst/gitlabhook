//go:build !windows

package git

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// LogFormat commit日志格式
const LogFormat = `{"commit":"%H","sanitized_subject_line":"%f","author":{"name":"%an","email":"%ae","timestamp":"%at"},"committer":{"name":"%cn","email":"%ce","timestamp":"%ct"}}`

func (g *gitCommit) GetInfo() (res *Commits, err error) {
	var bs []byte
	cmd := exec.Command("git", "log", "-n", "1", g.hash, fmt.Sprintf(`--pretty=format:%s`, LogFormat))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	bs, err = cmd.Output()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &res)
	if err != nil {
		return nil, err
	}
	return
}
