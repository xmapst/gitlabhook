//go:build !windows

package git

import (
	"os"
	"os/exec"
	"strings"
)

func (g *gitCommit) GetRevList() ([]string, error) {
	cmd := exec.Command("git", "rev-list", g.hash, "--not", "--all")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	bs, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	commitList := strings.Split(string(bs), "\n")
	var res []string
	for _, commits := range commitList {
		if strings.ReplaceAll(commits, " ", "") == "" {
			continue
		}
		res = append(res, commits)
	}
	return res, nil
}
