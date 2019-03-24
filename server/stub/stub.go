package stub

import (
	"errors"
	"fmt"
	"net"
)

type Stub struct {
	// packets chan ([]byte)
	handlers map[int32]func([]byte)
}

func New() *Stub {
	instance := &Stub{
		// packets: make([]byte, 0, 300),
		handlers: make(map[int32]func([]byte)),
	}

	return instance
}

func (stub *Stub) SetConnection(conn *net.Conn) {

}

func (stub *Stub) Handle(code int32, handler func(packet []byte)) error {

	if stub.handlers[code] != nil {
		return errors.New(fmt.Sprintf("code:%d has been added.", code))
	}

	stub.handlers[code] = handler
	return nil
}

func (stub *Stub) Start() {

}
