package main

import (
	"http-gateway/service/user"
	"http-gateway/service/user_cli"
	"time"
)

func main() {
	go user.StartUserServer()
	user_cli.CallSayHello()
	time.Sleep(1 * time.Second)
}
