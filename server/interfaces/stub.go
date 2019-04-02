package interfaces

import "net"

//Stub handle the message from client.
type Stub interface {
	New() Stub
	ID() int32
	Handle(int32, func([]byte))
	GetHandler(code int32) func([]byte, net.Conn) []byte
	Start()
}
