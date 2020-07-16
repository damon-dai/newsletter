package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息数据
}

type LoginMes struct {
	UserId   int    `json:"user_id"`   // 用户id
	UserPwd  string `json:"user_pwd"`  // 用户密码
	UserName string `json:"user_name"` // 用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
