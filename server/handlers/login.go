package handlers

import (
	"LunaFramework/server/interfaces"
	"log"
)

func HandlerLogin(server interfaces.Server) func([]byte) {
	dependancy := server.Dependancy()
	return func(packet []byte) {

		log.Printf("got dependancy:%d", dependancy)
		login(packet)
	}
}

func login(packet []byte) {
	playerID := packet[0]
	log.Printf("player(%d) login", playerID)
}
