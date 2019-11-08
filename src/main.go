package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

func main() {
	conn,err := redis.Dial("tcp","192.168.83.8:6379")
	if err != nil {
		fmt.Println("connect redis error :",err)
		return
	}else{
		defer conn.Close()
		_,err = conn.Do("AUTH","123456")
		if err != nil {
			fmt.Println("auth failed:",err)
		}

		const iLoopNum = 12
		for i:=0;i<iLoopNum;i++{
			ukey:= uuid.Must(uuid.NewV4())
			uvalue:= uuid.Must(uuid.NewV4())
			_, err = conn.Do("SET", ukey.String(), uvalue.String())
			if err != nil {
				fmt.Println("redis set error:", err)
			}else{
				fmt.Println(ukey.String(),"=======>", uvalue.String())
			}
		}

		name, err := redis.String(conn.Do("GET", "name"))
		if err != nil {
			fmt.Println("redis get error:", err)
		} else {
			fmt.Printf("Got name: %s \n", name)
		}
	}

}
