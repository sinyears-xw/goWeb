package mem

type Err struct {
	Code    int
	Message string
}

func (ep *Err) Error() string {
	return ep.Message
}

func (ep *Err)DecodeErr() (int, string) {
	if ep == nil {
		return OK.Code, OK.Message
	}
	return ep.Code, ep.Message
}

var (
	NotOK = &Err{Code: -1, Message: "服务器异常"}
	OK    = &Err{Code: 0, Message: "OK"}

	DBNotExist = &Err{Code: 10002, Message: "数据库不存在"}
	HeaderError = &Err{Code: 10003, Message: "HTTP头信息不全"}
	TypeAssertError = &Err{Code: 10004, Message: "类型断言错误"}
	TokenNotExit = &Err{Code: 10005, Message: "token不存在"}
	TokenNotValid = &Err{Code: 10006, Message: "token无效"}
	UserInfoNotValid = &Err{Code: 10007, Message: "用户信息无效"}
	HttpPortNotValid = &Err{Code: 10008, Message: "http服务端口无效"}
	GinNotValid = &Err{Code: 10009, Message: "Gin无效"}

	EcdsaError = &Err{Code: 20001, Message: "无法创建ECDSA"}
	EthAccountNotExit = &Err{Code: 20002, Message: "以太坊账户不存在"}
	EthAccountMoreThanOne = &Err{Code: 20003, Message: "以太坊账户不只一个"}
)
