//go:build !windows

package examiner

import (
	"fmt"
	"log"
	"pre-receive/internal/failed"
	"regexp"
	"strings"
)

func (e *Examiner) SetMessage(msg string) {
	e.Message = msg
}

func (e *Examiner) MsgLen(code int) {
	if len(e.Message) < 8 {
		failed.Exit(code, fmt.Sprintf(failed.CheckMsgLen, e.Hash))
	}
}

const (
	messagePattern = `^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-].*)\))?\: (.*)|^Merge\ branch(.*)`
	FEAT           = "feat"
	FIX            = "fix"
	DOCS           = "docs"
	STYLE          = "style"
	REFACTOR       = "refactor"
	TEST           = "test"
	CHORE          = "chore"
	PERF           = "perf"
	HOTFIX         = "hotfix"
)

var msgReg = regexp.MustCompile(messagePattern)

func (e *Examiner) MsgStyle(code int) {
	commitTypes := msgReg.FindAllStringSubmatch(e.Message, -1)
	if len(commitTypes) != 1 {
		log.Println("commit message样式不符合规范")
		failed.Exit(code, fmt.Sprintf(failed.CheckMessageStyle, e.Hash))
	} else {
		switch commitTypes[0][1] {
		case FEAT:
		case FIX:
		case DOCS:
		case STYLE:
		case REFACTOR:
		case TEST:
		case CHORE:
		case PERF:
		case HOTFIX:
		default:
			if !strings.HasPrefix(e.Message, "Merge branch") {
				failed.Exit(code, fmt.Sprintf(failed.CheckMessageStyle, e.Hash))
			}
		}

		// check message length
		if len(commitTypes[0][4]) < 8 {
			failed.Exit(code, fmt.Sprintf(failed.CheckMsgLen, e.Hash))
		}
	}
}
