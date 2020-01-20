package main

import (
	"github.com/common/message"
	"github.com/utils"
	"net"
	"server/process"
)

type Processor struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Processor)processor()  {
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	mes,err:=tf.ReadPkg()
	if err!=nil{
		println("read login mes failed")
		return
	}else {
		switch mes.Data {
		case message.LoginMesType:
			up:=&process.UserProcess{
				Conn: this.Conn,
			}
			up.ServerLogin(&mes)
		case message.LoginResMesType:
			return
		default:
			println("no mestype,byebye")

		}
	}

}