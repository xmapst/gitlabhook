//go:build !windows

package examiner

import (
	"fmt"
	"log"
	"pre-receive/internal/failed"
	"pre-receive/internal/git"
	"strings"
)

var staticFileSuffixList = []string{
	".png", ".jpg", ".jpeg", ".bmg", ".icon", ".gif", // 图片类型
	".dll", ".dat", ".jar", ".exe", ".meta", ".out", ".so", ".apk", // 二进制文件类型
	".mp3", ".mp4", ".wma", ".flv", ".wmv", ".ogg", ".avi", // 音视类型
	".zip", ".tar", ".gz", ".bz", ".bz2", ".xz", //压缩类型
}

func (e *Examiner) SetFileInfo(fileInfo []git.FileInfo) {
	e.FileInfos = fileInfo
}

func (e *Examiner) FileMaxSize(code int) {
	overMaxList := make([]string, 0)
	for _, f := range e.FileInfos {
		sizeMb := f.Size / 1048576
		// 处理常见的非文本类型文件
		exclude := e.inSliceString(f.Suffix)
		if exclude && f.Size >= 140 {
			overMaxList = append(overMaxList, fmt.Sprintf(failed.CheckFileSuffixPrintln, f.Name))
			continue
		}
		// 文件超过最大限制
		if f.Size >= e.MaxBytes {
			overMaxList = append(overMaxList, fmt.Sprintf(failed.CheckFileSizePrintln, f.Name, e.MaxBytes/1048576, sizeMb))
			continue
		}
	}
	if len(overMaxList) == 0 {
		return
	}
	log.Println("文件大小/类型检验不通过")
	failed.Exit(code, strings.Join(overMaxList, "\n")+"\n"+fmt.Sprintf(failed.CheckFile, e.MaxBytes/1048576))
}

func (e *Examiner) inSliceString(s string) bool {
	for _, v := range staticFileSuffixList {
		if strings.HasSuffix(s, v) {
			return true
		}
	}
	return false
}
