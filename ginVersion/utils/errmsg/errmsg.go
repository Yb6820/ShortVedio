package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//token相关
	TOKEN_NOT_EXISTS = 10010
	TOKEN_TYPE_ERROR = 10011
)

var (
	codeMsg = map[int]string{
		SUCCESS:          "OK",
		ERROR:            "FAIL",
		TOKEN_NOT_EXISTS: "token不存在",
		TOKEN_TYPE_ERROR: "token类型错误",
	}
)

func GetErrMsg(code int) (msg string) {
	return codeMsg[code]
}
