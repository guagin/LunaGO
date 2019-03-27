package interfaces

import "net"

//Stub handle the message from client.
type Stub interface {
	New() Stub
	SetConnection(*net.Conn)
	Handle(int32, func([]byte))
	Start()
}
