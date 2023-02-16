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
	".png", ".jpg", ".jpeg", ".bmp", ".icon", ".gif", ".tif", ".swf", // 图片类型
	".dll", ".dat", ".jar", ".exe", ".meta", ".out", ".so", ".apk", ".msi", ".pkg", // 二进制文件类型
	".mp3", ".mp4", ".wma", ".flv", ".wmv", ".ogg", ".avi", // 音视类型
	".zip", ".tar", ".gz", ".bz", ".bz2", ".xz", ".arj", ".7z", ".rar", ".mpk", ".ipa", ".rpm", ".deb", //压缩类型
	".iso", ".mdf", ".vhd", ".vhdx", ".dmg", // 镜像
	".db",  // sqlite
	".pdb", // 调试文件
}

func (e *Examiner) SetFileInfo(fileInfo []git.FileInfo) {
	e.FileInfos = fileInfo
}

func (e *Examiner) CheckFile(code int) {
	e.code = code
	e.checkSuffix()
	e.checkSize()
}

func (e *Examiner) checkSuffix() {
	overList := make([]string, 0)
	temp := map[string]struct{}{}
	for _, f := range e.FileInfos {
		// 处理常见的非文本类型文件
		exclude := e.inSliceString(f.Suffix)
		if exclude && f.Size >= 140 {
			overList = append(overList, fmt.Sprintf(failed.CheckFileSuffixPrintln, f.Name))
			if _, ok := temp[f.Suffix]; !ok {
				temp[f.Suffix] = struct{}{}
			}
			continue
		}
	}
	if len(overList) == 0 {
		return
	}
	title := "包含不属于文本类型文件"
	e.fileInfoFailed(title, overList, temp)
}

func (e *Examiner) checkSize() {
	overList := make([]string, 0)
	temp := map[string]struct{}{}
	for _, f := range e.FileInfos {
		// 文件超过最大限制
		if f.Size >= e.MaxBytes {
			overList = append(overList, fmt.Sprintf(failed.CheckFileSizePrintln, f.Name, f.Size/1048576, e.MaxBytes/1048576))
			if _, ok := temp[f.Suffix]; !ok {
				temp[f.Suffix] = struct{}{}
			}
			continue
		}
	}
	if len(overList) == 0 {
		return
	}
	title := fmt.Sprintf("包含大于 %d MB的文件", e.MaxBytes/1048576)
	e.fileInfoFailed(title, overList, temp)
}

func (e *Examiner) fileInfoFailed(title string, overList []string, temp map[string]struct{}) {
	log.Println("文件类型检验不通过")
	failed.Exit(e.code, strings.Join(overList, "\n"), "\n",
		fmt.Sprintf(failed.CheckFile, e.Hash, title, func(temp map[string]struct{}) string {
			var res []string
			for k := range temp {
				log.Printf("本次修改的文件后缀: %s\n", k)
				res = append(res, fmt.Sprintf(`git lfs track "*%s"`, k))
			}
			return strings.Join(res, "\n  ")
		}(temp)))
}

func (e *Examiner) inSliceString(s string) bool {
	for _, v := range staticFileSuffixList {
		if strings.HasSuffix(strings.ToLower(s), v) {
			return true
		}
	}
	return false
}
