//go:build !windows

package git

type Interface interface {
	GetInfo() (*Commits, error)
	GetMsg() (string, error)
	GetRevList() ([]string, error)
	GetFileInfos() ([]*FileInfo, error)
}

type gitCommit struct {
	hash string
}

func New(hash string) Interface {
	return &gitCommit{
		hash: hash,
	}
}

// ZeroCommit 空提交
const ZeroCommit = `0000000000000000000000000000000000000000`
