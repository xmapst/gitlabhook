//go:build !windows

package failed

const CheckMessageStyle = `##############################################################################
##                                                                          ##
## Commit message style check failed!                                       ##
##                                                                          ##
## Commit message style must satisfy this regular:                          ##
##   ^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-].*)\))?\: (. *)|^Merge\ branch(.*) ##
##                                                                          ##
## Example:                                                                 ##
##   feat(test): test commit style check.                                   ##
##                                                                          ##
##############################################################################`

const CheckUser = `##############################################

  提交用户名与gitlab登录用户名不匹配, 请修改!
  git config --global user.name "%s"

  修改后, 先操作撤销commit, 以下示例只是撤销一个
  git reset --soft HEAD~1

##############################################`

const CheckEmail = `##############################################

  提交邮箱与gitlab登录邮箱不匹配, 请修改!
  git config --global user.email "%s"

  修改后, 先操作撤销commit, 以下示例只是撤销一个
  git reset --soft HEAD~1

##############################################`

const CheckMsgLen = `################################################################
  先撤销commit， 然后再进行commit提交。以下示例只是撤销一个
  git reset --soft HEAD~1

################################################################`

const CheckFile = `####################################################################

  您的推送被拒绝，因为它包含的文件大于 %d MB或包含非源码文本类型文件
  先撤销commit， 参考以下链接开启lfs  
  请使用 https://git-lfs.github.com/ 存储更大的文件.
  然后再重新进行commit提交。以下示例只是撤销一个
  git reset --soft HEAD~1
  
####################################################################`

const CheckFileSizePrintln = `文件 %s 大于 %d MB。它的大小是 %d MB。`
const CheckFileSuffixPrintln = `文件 %s 属于非源码文本类型文件。请启用lfs。`
