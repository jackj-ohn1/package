package errno

type Errno struct {
	Message    string `json:"message"`
	StatusCode uint32 `json:"status_code"`
}

func (e *Errno) Error() string {
	return e.Message
}

func (e *Errno) Is(err *Errno) bool {
	return e == err
}

var (
	LoginWrongInfoError = &Errno{Message: "账号或密码错误", StatusCode: 400401}
	LoginServerError    = &Errno{Message: "登录出现异常", StatusCode: 500500}
	UserNotFound        = &Errno{Message: "用户不存在", StatusCode: 200204}
)
