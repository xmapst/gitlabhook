//go:build !windows

package engine

import (
	"log"
	"os"
	"pre-receive/internal/examiner"
	"pre-receive/internal/failed"
	"pre-receive/internal/git"
	"pre-receive/internal/gitlab"
	"strconv"
)

type Engine struct {
	UserInfo   *gitlab.UserInfo
	PushUser   string
	RevList    []string
	SkipTime   int
	StrictMode bool // strictMode 是否开启严格模式，严格模式下将校验所有的提交信息格式(多 commit 下)
	MaxBytes   int  // 文件最大字节
}

func (e *Engine) Run() {
	for _, commit := range e.RevList {
		gitCommit := git.New(commit)
		info, err := gitCommit.GetInfo()
		if err != nil {
			log.Println(err)
			failed.Exit(403, "can't get commit user info"+"\n"+commit+"\n"+err.Error())
		}

		// continue old commit
		_t, err := strconv.Atoi(info.Committer.Timestamp)
		if err != nil {
			log.Println(err)
			failed.Exit(500, "未知错误，请联系仓管")
		}
		if _t < e.SkipTime {
			continue
		}

		check := examiner.New(examiner.Examiner{
			UserInfo:   e.UserInfo,
			CommitInfo: info,
			PushUser:   e.PushUser,
			MaxBytes:   e.MaxBytes,
		})
		if check.SkipUser() {
			continue
		}
		// 只检查提交者
		//check.AuthorUser(601)
		//check.AuthorEmail(602)
		check.CommitterUser(603)
		check.CommitterEmail(604)

		msg, err := gitCommit.GetMsg()
		if err != nil {
			log.Println(err)
			failed.Exit(700, "commit描述提取失败")
		}
		check.SetMessage(msg)
		check.MsgLen(701)
		// 暂时不需要检查message的规范
		//check.MsgStyle(702)

		fileInfos, err := gitCommit.GetFileInfos()
		if err != nil {
			log.Println(err)
			failed.Exit(800, "文件列表获取失败"+"\n"+err.Error())
		}
		check.SetFileInfo(fileInfos)
		check.FileMaxSize(801)

		if !e.StrictMode {
			os.Exit(0)
		}
	}
}
