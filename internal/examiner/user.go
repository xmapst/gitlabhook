//go:build !windows

package examiner

import (
	"fmt"
	"pre-receive/internal/failed"
)

func (e *Examiner) SkipUser() bool {
	if e.CommitInfo.Author.Name == "Github" {
		return true
	}
	if e.CommitInfo.Committer.Name == "Github" {
		return true
	}
	return false
}

func (e *Examiner) AuthorUser(code int) {
	if e.CommitInfo.Author.Name == e.PushUser || e.CommitInfo.Author.Name == e.UserInfo.Name {
		return
	}
	msg := fmt.Sprintf("用户名不对：%s，应为：%s", e.CommitInfo.Author.Name, e.UserInfo.Name)
	failed.Exit(code, msg+"\n"+fmt.Sprintf(failed.CheckUser, e.UserInfo.Name))
}

func (e *Examiner) CommitterUser(code int) {
	if e.CommitInfo.Committer.Name == e.PushUser || e.CommitInfo.Committer.Name == e.UserInfo.Name {
		return
	}
	msg := fmt.Sprintf("用户名不对：%s，应为：%s", e.CommitInfo.Committer.Name, e.UserInfo.Name)
	failed.Exit(code, msg+"\n"+fmt.Sprintf(failed.CheckUser, e.UserInfo.Name))
}

func (e *Examiner) AuthorEmail(code int) {
	if e.CommitInfo.Author.Email == e.UserInfo.Email {
		return
	}
	msg := fmt.Sprintf("邮箱不对：%s，应为：%s", e.CommitInfo.Author.Email, e.UserInfo.Email)
	failed.Exit(code, msg+"\n"+fmt.Sprintf(failed.CheckEmail, e.UserInfo.Email))
}

func (e *Examiner) CommitterEmail(code int) {
	if e.CommitInfo.Committer.Email == e.UserInfo.Email {
		return
	}
	msg := fmt.Sprintf("邮箱不对：%s，应为：%s", e.CommitInfo.Committer.Email, e.UserInfo.Email)
	failed.Exit(code, msg+"\n"+fmt.Sprintf(failed.CheckEmail, e.UserInfo.Email))
}
