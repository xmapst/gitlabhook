//go:build !windows

package git

import (
	"log"
	"path/filepath"
	"pre-receive/internal/utils"
	"strconv"
	"strings"
)

// 使用git ls-tree --full-name -r -l HEAD 获取文件信息
// 输出格式 git ls-tree --full-name -r -l HEAD

func (g *gitCommit) GetFileInfos() (res []FileInfo, err error) {
	allFileInfo := g.allFileInfo()
	fileInfo := g.fileInfos()
	for _, v := range allFileInfo {
		if g.equal(v.FullPath, fileInfo) {
			log.Printf("本次有修改的文件: %s\n", v.FullPath)
			res = append(res, v)
		}
	}
	return
}

func (g *gitCommit) equal(name string, files []string) bool {
	for _, file := range files {
		if name == file {
			return true
		}
	}
	return false
}

func (g *gitCommit) fileInfos() (res []string) {
	output, err := utils.Exec("git", "diff-tree", "-r", "--diff-filter=ACMRT", g.hash)
	if err != nil {
		log.Println(err)
		return
	}
	for _, line := range output {
		name := g.cutFileInfo(line)
		if name == "" {
			continue
		}
		res = append(res, name)
	}
	return
}

func (g *gitCommit) cutFileInfo(s string) string {
	ss := strings.Fields(s)
	if len(ss) != 6 {
		return ""
	}
	return ss[5]
}

func (g *gitCommit) allFileInfo() (res []FileInfo) {
	output, err := utils.Exec("git", "ls-tree", "--full-name", "-r", "-l", g.hash)
	if err != nil {
		log.Println(err)
		return
	}
	for _, line := range output {
		fileInfo := g.cutAllFileInfo(line)
		if fileInfo == nil {
			continue
		}
		res = append(res, *fileInfo)
	}
	return
}

func (g *gitCommit) cutAllFileInfo(s string) (res *FileInfo) {
	ss := strings.Fields(s)
	if len(ss) != 5 {
		return
	}
	if strings.HasPrefix(ss[4], ".") {
		return
	}
	size, err := strconv.Atoi(ss[3])
	if err != nil {
		return
	}
	return &FileInfo{
		Mode:     ss[0],
		Type:     ss[1],
		Object:   ss[2],
		Size:     size,
		Name:     filepath.Base(ss[4]),
		Path:     filepath.Dir(ss[4]),
		FullPath: ss[4],
		Suffix:   strings.ReplaceAll(filepath.Ext(ss[4]), `"`, ""),
	}
}
