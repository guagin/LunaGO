package server

import (
	"LunaFramework/server/handlers"
	"LunaFramework/server/interfaces"
	"LunaFramework/server/stub"
	"log"
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

func (server *server) handleNewConnection(connection *net.Conn) {
	stub := stub.New()
	stub.SetConnection(connection)
	stub.Handle(0, handlers.HandlerLogin(server))
	stub.Start()

}

func (server *server) Start(cIndex int32, c net.Conn) {
	defer c.Close()
	for {
		var buf = make([]byte, 10)
		log.Printf("start to read from conn[%d]\n", cIndex)
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("conn[%d] read %d bytes, content is %s\n", cIndex, n, string(buf[:n]))
	}

}
