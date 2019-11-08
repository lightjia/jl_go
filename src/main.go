package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn,err := redis.Dial("tcp","10.1.210.69:6379")
	if err != nil {
		fmt.Println("connect redis error :",err)
		return
	}
	defer conn.Close()
}
