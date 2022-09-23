package main

import (
	"fmt"
	"nats/server"
	"nats/storage/cashe"
	"nats/storage/db"
	"time"
)

func main() {
	db.Connect()
	db.Recover()

	//pub.Pub()
	//sub.Subscribe()

	time.Sleep(time.Second * 3)
	fmt.Println(cashe.GetOrder("b563feb7b2b84b6test"))

	server.Run()

}
