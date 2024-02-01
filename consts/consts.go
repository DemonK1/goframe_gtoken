package consts

var (
	CreateFieldEx = []any{"update_by", "update_name", "update_time", "delete_time"}
	UpdateFieldEx = []any{"id", "create_name", "create_time", "delete_time"}
)

const (
	Token                     = "token"
	UserKey                   = "userKey"
	UsernameOrPasswordNotNull = "用户名或密码不能为空"
	UserNotExist              = "用户不存在"
	UserExist                 = "用户已存在"
	UsernameOrPasswordErr     = "用户名或密码错误"
	UsernameAndPasswordLen    = "用户名的长度在3-15之间，密码的长度在4-16之间"
	ServerErr                 = "内部错误"
)
