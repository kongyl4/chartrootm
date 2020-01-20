package main


import (
"fmt"
)

func writeData(initChanl chan int)  {
	for i:=1;i<=8000;i++{
		initChanl<-i
	}
	close(initChanl)

}
func prime(initChanl chan int,primeChanl chan int,exitChanl chan bool)  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("prime()fashengcuowu=",err)
		}
	}()
	for{
		flag:=true
		v,ok:=<-initChanl
		if !ok{
			break
		}
		for j:=2;j<v;j++{
			if v%j == 0{
				flag = false
			}
		}
		if flag == true{
			primeChanl<-v
		}
	}
	exitChanl<-true

}
func main() {
	initChanl:=make(chan int,8000)
	primeChanl:=make(chan int,8000)
	exitChanl:=make(chan bool,4)
	//start:=time.Now()
	go writeData(initChanl)
	var num int = 4
	for i:=0;i<num;i++{
		go prime(initChanl,primeChanl,exitChanl)
	}
	go func() {
		for i:=0;i<num;i++{
			<-exitChanl
		}
		close(exitChanl)
		close(primeChanl)
	}()

	for {
		res,ok:=<-primeChanl
		if !ok{
			break
		}
		fmt.Println("sushu=%d\n",res)
	}
	//end:=time.Now()
	//fmt.Println("time=%d",)
}
