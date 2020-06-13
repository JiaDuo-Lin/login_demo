package main

import (
	"demo/jwt"
	"demo/server"
)


func main() {
	go server.StartServer()
	go jwt.StartServer()
	select {}
}

