//go:build !windows

package git

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

// 使用git ls-tree --full-name -r -l HEAD 获取文件信息
// 输出格式 git ls-tree --full-name -r -l HEAD

func (g *gitCommit) GetFileInfos() (res []*FileInfo, err error) {
	var stdout io.ReadCloser
	cmd := exec.Command("git", "ls-tree", "--full-name", "-r", "-l", g.hash)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	render := bufio.NewReader(stdout)
	for {
		bs, _, err := render.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
		fileInfo := g.cutFileInfo(string(bs))
		if fileInfo == nil {
			continue
		}
		res = append(res, fileInfo)
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	exitCode := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	if exitCode != 0 {
		return nil, fmt.Errorf("exit code %d", exitCode)
	}
	return
}

func (g *gitCommit) cutFileInfo(s string) (res *FileInfo) {
	ss := strings.Fields(s)
	if len(ss) < 4 {
		return
	}

	switch {
	case strings.Contains(ss[4], ".idea"):
		return
	case strings.Contains(ss[4], ".vscode"):
		return
	}
	size, err := strconv.Atoi(ss[3])
	if err != nil {
		return
	}
	return &FileInfo{
		Mode:   ss[0],
		Type:   ss[1],
		Object: ss[2],
		Size:   size,
		Name:   filepath.Base(ss[4]),
		Path:   filepath.Dir(ss[4]),
		Suffix: strings.ReplaceAll(filepath.Ext(ss[4]), `"`, ""),
	}
}
