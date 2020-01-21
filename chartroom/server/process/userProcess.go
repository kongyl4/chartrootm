package process

import (
	"encoding/json"
	"fmt"
	"github.com/common/message"
	"github.com/gomodule/redigo/redis"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *UserProcess)ServerLogin(mes *message.Message)  {
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data),&loginMes)
	userid:=loginMes.UserId

	conn,err:=redis.Dial("tcp","localhost:6379")
	if err!=nil{
		println("connect redis failed")
	}
	_,err=conn.Do("set","1","{\"userId\":\"1\",\"userName\":\"kongyl\\n\",\"userPwd\":\"123456\\n\"}")
	if err!=nil{
		fmt.Println("set name=tommy")
	}
	_,err=conn.Do("get","name")
	userinfo,err:=conn.Do("get",userid)
	if userinfo==mes.Data{
		fmt.Printf("%vlogin success",userid)
	}

}
