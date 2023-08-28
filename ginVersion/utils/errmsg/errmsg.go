package errmsg

const (
	SUCCESS = 0
	ERROR   = -1

	//token相关
	TOKEN_NOT_EXISTS    = 10010
	TOKEN_TYPE_ERROR    = 10011
	PARSE_TOKRN_ERROR   = 10012
	TOKEN_RUNTIME_ERROR = 10013
)

var (
	codeMsg = map[int]string{
		SUCCESS:             "OK",
		ERROR:               "FAIL",
		TOKEN_NOT_EXISTS:    "token不存在",
		TOKEN_TYPE_ERROR:    "token类型错误",
		PARSE_TOKRN_ERROR:   "token解析错误",
		TOKEN_RUNTIME_ERROR: "token已过期",
	}
)

func GetErrMsg(code int) (msg string) {
	return codeMsg[code]
}
