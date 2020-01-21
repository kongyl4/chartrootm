package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Transfer)WritePkg(data []byte) (err error) {

	dataLen:= len(data)
	databyte:=make([]byte,1024)
	binary.BigEndian.PutUint32(databyte,uint32(dataLen))

	fmt.Println("client n=",dataLen)
	n ,err := this.Conn.Write(databyte)

	if err!=nil{
		println("send data length failed",n)
		return
	}

	fmt.Println("client data=",string(data))
	n,err = this.Conn.Write(data)
	if err!=nil{
		fmt.Println("send data failed")
		return
	}
	return
}
func (this *Transfer)ReadPkg() (mes message.Message, err error) {
	this.Conn.Read(this.Buf[:4])
	pkgLen:=binary.BigEndian.Uint32(this.Buf[:4])

	fmt.Println(pkgLen)
	n,err:=this.Conn.Read(this.Buf[:pkgLen])
	fmt.Println(n)
	if err!=nil{
		fmt.Println("read error")
		return
	}

	if n!=int(pkgLen) || err!=nil{
		err = json.Unmarshal(this.Buf[:pkgLen],&mes)//why it's address
		fmt.Println("unmarshal failed")
		return
	}
	fmt.Println("mes is",&mes)
	return
}

