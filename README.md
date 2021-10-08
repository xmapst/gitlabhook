# GitlabHook

gitlab服务端的控制狂钩钩子, 主要实现类似Bit-Booster的Control Freak Hook

## 实现原理
在进行`push`操作时，[GitLab](https://docs.gitlab.com/ee/administration/server_hooks.html) 会调用这个钩子文件，并且从`stdin`输入三个参数，
分别为:
- 上次`push`的`hash`
- 现在`push`的`hash`
- 现在`push`的`branch` 

根据`commit hash`我们就可以很轻松的获取到提交信息，从而实现进一步检测动作；
根据[GitLab](https://docs.gitlab.com/ee/administration/server_hooks.html) 的文档说明，
当这个`hook`执行后以非`0`状态退出则认为执行失败，从而拒绝`pus`；同时会将`stderr`信息返回给`client`

## 编译  
+ 测试的token: bDx85ho5cNF8DC1fp4bP  
+  生产的token: (保密) 
```bash
CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "-s -w -X 'main.GitlabAdminUrl=http://127.0.0.1' -X 'main.GitlabAdminToken=bDx85ho5cNF8DC1fp4bP'"
```
- -X importpath.name=value
  Set the value of the string variable in importpath named name to value.
  Note that before Go 1.5 this option took two separate arguments.
  Now it takes one argument split on the first = sign.

## 部署
### 单独项目配置
钩子文件必须放在`/var/opt/gitlab/git-data/repositories/<group>/<project>.git/custom_hooks`目录中，
当然具体路径也可能是`/home/git/repositories/<group>/<project>.git/custom_hooks`；
`custom_hooks`目录需要自己创建

### 全局项目配置
- 对于从源安装的默认目录通常是`/home/git/gitlab-shell/hooks`。在此位置创建一个新目录。
  对于`Omnibus GitLab`(`docker`部署的`gitlab`也是在此路径)，通常是安装`/opt/gitlab/embedded/service/gitlab-shell/hooks`

取决于钩的类型，它可以是一个`pre-receive.d`,`post-receive.d`或`update.d`目录,`hooks`有可能需要自己创建. 
具体路径示例`/opt/gitlab/embedded/service/gitlab-shell/hooks/pre-receive.d/`


## 参考文档
[https://docs.gitlab.com/ee/administration/server_hooks.html](https://docs.gitlab.com/ee/administration/server_hooks.html)  
[https://docs.gitlab.com/ee/administration/server_hooks.html#create-a-global-server-hook-for-all-repositories](https://docs.gitlab.com/ee/administration/server_hooks.html#create-a-global-server-hook-for-all-repositories)