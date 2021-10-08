//go:build !windows

package git

import (
	"os"
	"os/exec"
	"strings"
)

func (g *gitCommit) GetMsg() (string, error) {
	cmd := exec.Command("git", "log", "-n", "1", g.hash, "--pretty=format:%s")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	bs, err := cmd.Output()
	if err != nil {
		return "", err
	}
	res := strings.ReplaceAll(string(bs), " ", "")
	res = strings.ReplaceAll(string(bs), "	", "")
	res = strings.ReplaceAll(string(bs), "\n", "")
	return res, nil
}
