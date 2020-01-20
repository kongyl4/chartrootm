package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
type LoginMes struct {
	UserId string `json:"userId"`
	UserName string `json:"userName"`
	UserPwd string `json:"userPwd"`
}
type LoginResMes struct {
	Code int `json:"code"`
	Error string `json:"error"`
} 
