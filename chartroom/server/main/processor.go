package main

import (
	"fmt"
	"github.com/common/message"
	"github.com/utils"
	"net"
	"server/process"
)

type Processor struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Processor)ProcessHandle()  {
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	mes,err:=tf.ReadPkg()
	fmt.Println("mes.type=",mes.Type)
	fmt.Println(message.LoginResMesType)
	if err!=nil{
		println("read login mes failed",err)
		return
	}else {
		switch mes.Type {
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