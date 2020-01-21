package main

import (
	"fmt"
	"net"
	_ "strings"
)

//func process(conn net.Conn)  {
//	defer conn.Close()
//	buf:=make([]byte,1024)
//	for{
//		n,error:=conn.Read(buf)
//		if error!=nil{
//			println("read error")
//			break
//		}else {
//			line:=string(buf[:n])
//			if strings.Trim(line,"\r\n") == "byebye"{
//				fmt.Printf("%v tuichule", conn.RemoteAddr())
//				fmt.Println()
//				break
//			}
//			println(line)
//
//		}
//	}
//
//}
func main() {
	ln,error := net.Listen("tcp","localhost:8080")
	if error!=nil{
		fmt.Println("can not listen port 8080")
		return
	}
	fmt.Println("start listening")
	for{
		conn,error := ln.Accept()
		if error != nil{
			println("连接错误")
			return
		}else {
			process:=&Processor{
				Conn: conn,
				Buf:  [8096]byte{},
			}
			process.ProcessHandle()
			//println(conn.RemoteAddr())
			//go process(conn)
		}
	}
}
