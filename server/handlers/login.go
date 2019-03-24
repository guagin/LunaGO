package handlers

import (
	"LunaGO/server/interfaces"
	"LunaGO/server/messages"
	"log"
)

func HandlerLogin(server interfaces.Server) func([]byte) []byte {
	dependancy := server.Dependancy()
	return func(packet []byte) []byte {

		log.Printf("got dependancy:%d", dependancy)
		return login(packet)
	}
}

func login(packet []byte) []byte {
	login, err := messages.UnmarshalLogin(packet)
	if err != nil {
		log.Println("login error:", err)
	}
	log.Printf("player(%s) login", login.ID)
	return nil
}
