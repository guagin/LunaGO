package interfaces

import "net"

type Stub interface {
	New() Stub
	SetConnection(*net.Conn)
	Handle(int32, func([]byte))
	Start()
}
