package server

import (
	"LunaGO/server/interfaces"
	"LunaGO/server/stub"
	"errors"
	"net"
)

// Server handle the connection in coming. and keep the dependancies
type Server struct {
	interfaces.Server
	port              int32
	subs              map[int32]stub.Stub
	dependancy        int32
	connectionHandler func(int32, net.Conn)
}

// New initialize a server struct
func New() *Server {

	instance := &Server{
		port:       55550,
		dependancy: 1,
	}

	return instance
}

// get dependancy
func (server Server) Dependancy() int32 {
	return server.dependancy

}

// HandleNewConnection handle the connection for new incoming
func (server *Server) HandleNewConnection(cIndex int32, c net.Conn) {
	server.connectionHandler(cIndex, c)
}

// setup how to deal with new connection coming.
func (server *Server) SetConnectionHandler(handler func(int32, net.Conn)) error {
	if handler == nil {
		return errors.New("handler is nil")
	}
	server.connectionHandler = handler
	return nil
}
