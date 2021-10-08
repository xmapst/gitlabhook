//go:build !windows

package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pre-receive/internal/failed"
)

var (
	AdminToken string
	AdminUrl   string
)

func GetUserInfo(userName string) *UserInfo {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v4/users?username=%s", AdminUrl, userName), nil)
	if err != nil {
		log.Println(err)
		failed.Exit(9, "GL-HOOK-ERR: exit code 9-1, 未知错误, 请联系管理员")
	}
	req.Header.Set("Private-Token", AdminToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		failed.Exit(9, "GL-HOOK-ERR: exit code 9-2, 未知错误, 请联系管理员")
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		failed.Exit(9, "GL-HOOK-ERR: exit code 9-3, 未知错误, 请联系管理员")
	}
	var userInfos []UserInfo
	if err := json.Unmarshal(bodyByte, &userInfos); err != nil {
		log.Println(err)
		failed.Exit(9, "GL-HOOK-ERR: exit code 9-1, 未知错误, 请联系管理员")
	}
	if len(userInfos) == 0 {
		failed.Exit(9, fmt.Sprintf("GL-HOOK-ERR: exit code 9-5, 未找到%s用户", userName))
	}
	return &userInfos[0]
}
