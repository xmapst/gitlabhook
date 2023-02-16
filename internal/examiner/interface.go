//go:build !windows

package examiner

import (
	"pre-receive/internal/git"
	"pre-receive/internal/gitlab"
)

type Interface interface {
	SkipUser() bool
	AuthorUser(int)
	CommitterUser(int)
	AuthorEmail(int)
	CommitterEmail(int)
	SetMessage(msg string)
	MsgLen(int)
	MsgStyle(code int)
	SetFileInfo(fileInfo []git.FileInfo)
	CheckFile(code int)
}

type Examiner struct {
	Hash       string
	UserInfo   *gitlab.UserInfo
	CommitInfo *git.Commits
	FileInfos  []git.FileInfo
	Message    string
	PushUser   string
	MaxBytes   int
	code       int
}

func New(e Examiner) Interface {
	return &e
}
