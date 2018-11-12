package core

const (
	HOST_NETWORK = "tcp"
	HOST_NAME    = "localhost:8000"
)

const (
	OPERATE_REGISTER = "register"
	OPERATE_LOGIN    = "login"
	OPERATE_LOGOFF   = "logoff"
	OPERATE_EXIT     = "exit"
	OPERATE_ADD      = "add"
	OPERATE_DELETE   = "delete"
	OPERATE_LIST     = "list"
	OPERATE_SENDTO   = "sendTo"
	OPERATE_SENDALL  = "sendAll"
)

// 文件常量
const FILE_PATH_USER = "./files/user.txt"


const USAGE =
`
please choose options:
	- register : 注册。            格式:"register"
	- login    : 登录。            格式:"login"
	- logoff   : 注销登录。         格式:"logoff"
	- exit     : 退出系统。         格式:"exit"
	- add      : 添加好友。         格式:"add"
	- delete   : 删除好友。         格式:"delete"
	- list     : 查看好友列表。      格式:"list"
	- sendTo   : 给某位好友发送消息。 格式:"sendTo"
	- sendAll  : 给全部好友发送消息。 格式:"sendAll"
`
