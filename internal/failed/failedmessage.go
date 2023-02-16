//go:build !windows

package failed

const CheckMessageStyle = `
  Hash: %s
  Commit message样式不符合规范
  Commit消息风格必须满足这个规则
    ^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-].*)\))?\: (. *)|^Merge\ branch(.*)

  示例:
    feat(test): test commit style check.
`

const CheckUser = `
  Hash: %s
  Commit用户名与gitlab登录用户名不匹配, 请修改!
  git config --global user.name "%s"

  修改后, 先操作撤销commit, 以下示例只是撤销一个
  git reset --soft HEAD~1
`

const CheckEmail = `
  Hash: %s
  Commit邮箱与gitlab登录邮箱不匹配, 请修改!
  git config --global user.email "%s"

  修改后, 先操作撤销commit, 以下示例只是撤销一个
  git reset --soft HEAD~1
`

const CheckMsgLen = `
  Hash: %s
  commit描述太短，没有任何意义! 先撤销本次commit，
  然后再进行commit提交. 以下示例只是撤销一个
  git reset --soft HEAD~1
`

const CheckFile = `
  Hash: %s
  您的推送被拒绝，%s, 请按要求修正后重新提交
  参考以下命令进行撤销commit并设置追踪需要使用 Git LFS 管理的文件
  git reset --soft HEAD~1
  git lfs install
  %s
`

const CheckFileSizePrintln = `  文件 %s 的大小是 %d MB, 大于单文件 %d MB的容忍`
const CheckFileSuffixPrintln = `  文件 %s 不属于文本类型文件, 请启用lfs`
