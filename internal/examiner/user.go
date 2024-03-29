//go:build !windows

package examiner

import (
	"fmt"
	"pre-receive/internal/failed"
	"strings"
)

var skipUsers = []string{"github", "gitlab", "gitee", "codeup", "coding"}

func (e *Examiner) SkipUser() bool {
	for _, user := range skipUsers {
		if strings.ToLower(e.CommitInfo.Author.Name) == user {
			return true
		}
		if strings.ToLower(e.CommitInfo.Committer.Name) == user {
			return true
		}
	}
	return false
}

func (e *Examiner) AuthorUser(code int) {
	if e.CommitInfo.Author.Name == e.PushUser || e.CommitInfo.Author.Name == e.UserInfo.Name {
		return
	}
	msg := fmt.Sprintf("用户名不对：%s，应为：%s", e.CommitInfo.Author.Name, e.UserInfo.Name)
	failed.Exit(code, msg, "\n", fmt.Sprintf(failed.CheckUser, e.Hash, e.UserInfo.Name))
}

func (e *Examiner) CommitterUser(code int) {
	if e.CommitInfo.Committer.Name == e.PushUser || e.CommitInfo.Committer.Name == e.UserInfo.Name {
		return
	}
	msg := fmt.Sprintf("用户名不对：%s，应为：%s", e.CommitInfo.Committer.Name, e.UserInfo.Name)
	failed.Exit(code, msg, "\n", fmt.Sprintf(failed.CheckUser, e.Hash, e.UserInfo.Name))
}

func (e *Examiner) AuthorEmail(code int) {
	if e.CommitInfo.Author.Email == e.UserInfo.Email {
		return
	}
	msg := fmt.Sprintf("邮箱不对：%s，应为：%s", e.CommitInfo.Author.Email, e.UserInfo.Email)
	failed.Exit(code, msg, "\n", fmt.Sprintf(failed.CheckEmail, e.Hash, e.UserInfo.Email))
}

func (e *Examiner) CommitterEmail(code int) {
	if e.CommitInfo.Committer.Email == e.UserInfo.Email {
		return
	}
	msg := fmt.Sprintf("邮箱不对：%s，应为：%s", e.CommitInfo.Committer.Email, e.UserInfo.Email)
	failed.Exit(code, msg, "\n", fmt.Sprintf(failed.CheckEmail, e.Hash, e.UserInfo.Email))
}
