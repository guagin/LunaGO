package server

import (
	"LunaGO/server/handlers"
	"LunaGO/server/interfaces"
	"LunaGO/server/stub"
	"net"
)

type server struct {
	interfaces.Server
	port       int32
	subs       map[int32]stub.Stub
	dependancy int32
}

func New() *server {

	instance := &server{
		port:       55550,
		dependancy: 1,
	}

	return instance
}

func (server server) Dependancy() int32 {
	return server.dependancy

}

func (server *server) HandleNewConnection(cIndex int32, c net.Conn) {
	defer c.Close()
	stub := stub.New(cIndex)
	stub.SetConnection(c)
	stub.Handle(0, handlers.HandlerLogin(server))
	stub.Start()
}
