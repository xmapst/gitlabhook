//go:build !windows

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "pre-receive/engine"
    "pre-receive/internal/git"
    "pre-receive/internal/gitlab"
    "strings"
)

var (
	// GitlabAdminUrl gitlab 地址
	GitlabAdminUrl string
	// GitlabAdminToken 获取api认证token
	GitlabAdminToken string
	projectID        string
	pushProtocol     string
	pushUser         string
)

func init() {
	// 设置日志输出到/tmp/pre-receive.log中
	logPath := filepath.Join(os.TempDir(), "pre-receive.log")
	logWriter, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	if GitlabAdminUrl == "" {
		GitlabAdminUrl = "http://127.0.0.1"
	}
	// 获取项目id
	projectID = os.Getenv("GL_PROJECT_PATH")
	// 提交协议: web/http/ssh
	pushProtocol = os.Getenv("GL_PROTOCOL")
	// 获取用户信息
	pushUser = os.Getenv("GL_USERNAME")

	// 设置日志输出
	log.SetOutput(logWriter)
	log.SetPrefix(fmt.Sprintf("project=%s user=%s push_protocol=%s msg=", projectID, pushUser, pushProtocol))
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC | log.Lmsgprefix)
	log.Println("Gitlab URL", GitlabAdminUrl)
	gitlab.AdminUrl = GitlabAdminUrl
	gitlab.AdminToken = GitlabAdminToken
}

func main() {
	// 在进行 push 操作时，GitLab 会调用这个钩子文件，并且从 stdin 输入三个参数，分别为 之前的版本 commit ID、push 的版本 commit ID 和 push 的分支
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	param := strings.Fields(string(input))
	if len(param) < 3 {
		os.Exit(0)
	}
	// 允许删除分支/标签
	if param[1] == git.ZeroCommit {
		os.Exit(0)
	}

	// 检查新的分支或标签
	var span string
	if param[0] == git.ZeroCommit {
		span = param[1]
	} else {
		span = fmt.Sprintf("%s...%s", param[0], param[1])
	}
	gitCommit := git.New(span)
	revList, err := gitCommit.GetRevList()
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	// 空列表不检查
	if len(revList) == 0 {
		os.Exit(0)
	}

	userInfo := gitlab.GetUserInfo(pushUser)
	e := engine.Engine{
		UserInfo:   userInfo,
		PushUser:   pushUser,
		RevList:    revList,
		StrictMode: true,
		MaxBytes:   5 * 1048576, // 5MB
	}
	e.Run()
}
