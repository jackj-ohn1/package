package errno

type Errno struct {
	message    string `json:"message"`
	statusCode uint32 `json:"status_code"`
}

func (e *Errno) GetStatusCode() uint32 {
	return e.statusCode
}

func (e *Errno) GetMessage() string {
	return e.message
}

func (e *Errno) Error() string {
	return e.message
}

func (e *Errno) Is(err *Errno) bool {
	return e == err
}

func IsError(e *Errno, err error) bool {
	if one, ok := err.(*Errno); ok {
		return e.Is(one)
	}
	return false
}

var (
	LoginWrongInfoError = &Errno{message: "账号或密码错误", statusCode: 400401}
	LoginServerError    = &Errno{message: "登录出现异常", statusCode: 500500}
	UserNotFound        = &Errno{message: "用户不存在", statusCode: 200204}
	JsonDataError       = &Errno{message: "JSON数据绑定失败", statusCode: 400422}
)
