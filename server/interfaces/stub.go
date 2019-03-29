package interfaces

import "net"

//Stub handle the message from client.
type Stub interface {
	New() Stub
	Handle(int32, func([]byte))
	GetHandler(code int32) func([]byte, net.Conn) []byte
	Start()
}
