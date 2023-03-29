package errno

type Errno struct {
	message    string `json:"message"`
	statusCode int    `json:"status_code"`
}

func GetCode(message string) int {
	return storage[message]
}

func (e *Errno) Error() string {
	return e.message
}

var storage = map[string]int{
	JsonDataError.message:       JsonDataError.statusCode,
	LoginServerError.message:    LoginServerError.statusCode,
	LoginWrongInfoError.message: LoginWrongInfoError.statusCode,
	UserNotExistError.message:   UserNotExistError.statusCode,
}

var (
	// user
	LoginWrongInfoError = &Errno{message: "账号或密码错误", statusCode: 400401}
	LoginServerError    = &Errno{message: "登录出现异常", statusCode: 500500}
	JsonDataError       = &Errno{message: "JSON数据绑定失败", statusCode: 400422}
	UserNotExistError   = &Errno{message: "用户不存在", statusCode: 200204}
)
