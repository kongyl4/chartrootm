package main

import (
	"bufio"
	"client/process"
	"fmt"
	"os"
	_"strings"
)


func main() {
	var loop bool = true
	for loop{
		println("1 login the room")
		println("2 register new user")
		println("3 exit system")
		var key int
		fmt.Scanf("%d\n",&key)
		switch key{
		case 1:
			var username string
			var password string
			println("please enter the username")
			fmt.Scanf("%v",&username)
			//username,_=bufio.NewReader(os.Stdin).ReadString('\n')
			println("please enter the password")
			fmt.Scanf("%v",&password)
			//password,_=bufio.NewReader(os.Stdin).ReadString('\n')
			up:=&process.UserProcess{}
			up.Login(username,password)
			loop = false
		case 2:
			println("please enter the username")
			username,_:=bufio.NewReader(os.Stdin).ReadString('\n')
			println("please enter the password")
			password,_:=bufio.NewReader(os.Stdin).ReadString('\n')
			up:=&process.UserProcess{}
			up.Login(username,password)
			loop = false
		case 3:
			println("I will logout! byebye")
			os.Exit(0)

		}

	}




	//for{
	//	reader:=bufio.NewReader(os.Stdin)
	//	line,error:=reader.ReadString('\n')
	//	if error!=nil{
	//		println("写入信息错误")
	//		return
	//	}
	//	line = strings.Trim(line,"\r\n")
	//	conn.Write([]byte(line+"\n"))
	//	if line == "byebye"{
	//		println("client exit")
	//		conn.Write([]byte(line+"\n"))
	//		break
	//	}
	//}

}
