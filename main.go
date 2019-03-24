package main

import (
	"LunaFramework/server"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":55555")
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	server_1 := server.New()
	var connIndex int32 = 0
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept err:", err)
			break
		}
		go server_1.Start(connIndex, c)
		connIndex++
	}

}
