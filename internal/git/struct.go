//go:build !windows

package git

type FileInfo struct {
	Mode   string `json:"mode,omitempty"`
	Type   string `json:"type,omitempty"`
	Object string `json:"object,omitempty"`
	Size   int    `json:"size,omitempty"`
	Name   string `json:"name,omitempty"`
	Path   string `json:"path,omitempty"`
	Suffix string `json:"suffix,omitempty"`
}

type Commits struct {
	Commit               string `json:"commit,omitempty"`
	AbbreviatedCommit    string `json:"abbreviated_commit,omitempty"`
	Subject              string `json:"subject,omitempty"`
	SanitizedSubjectLine string `json:"sanitized_subject_line,omitempty"`
	Author               User   `json:"author,omitempty"`
	Committer            User   `json:"committer,omitempty"`
}

type User struct {
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Date         string `json:"date,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
	RelativeDate string `json:"relative_date,omitempty"`
}
