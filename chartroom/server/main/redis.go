package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func RedisCase()  {

	conn,err:=redis.Dial("tcp","localhost:6379")
	if err!=nil{
		println("connect redis failed")
	}
	_,err=conn.Do("set","name","tommy")
	if err!=nil{
		fmt.Println("set name=tommy")
	}
	_,err=conn.Do("get","name")
}
