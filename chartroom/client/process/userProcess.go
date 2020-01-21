package process

import (
	"encoding/json"
	"fmt"
	"github.com/common/message"
	"net"
    "github.com/utils"
)

type UserProcess struct {

}

func (this *UserProcess) Login(username string,password string) (er error)  {
	conn,er:=net.Dial("tcp","localhost:8080")
	if er!=nil{
		fmt.Println("client connect server error")
		return
	}else {

		loginMes:=&message.LoginMes{
			UserId:   "1",
			UserName: username,
			UserPwd:  password,
		}
		loginMesJson,error:=json.Marshal(loginMes)
		if error!=nil{
			fmt.Println("marshal failed")
			return error
		}
		mes:=&message.Message{
			Type: "LoginMes",
			Data: string(loginMesJson),
		}
		data,_:=json.Marshal(mes)

		tf:=utils.Transfer{
			Conn: conn,
		}
		fmt.Println("data=",string(data))
		tf.WritePkg(data)
		tf.ReadPkg()
	}
	return er
}