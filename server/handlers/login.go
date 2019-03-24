package handlers

import (
	"LunaGO/server/interfaces"
	"LunaGO/server/messages"
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
	login, err := messages.UnmarshalLogin(packet)
	if err != nil {
		log.Println("login error:", err)
	}
	log.Printf("player(%s) login", login.ID)
}
