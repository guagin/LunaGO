package stub

import (
	"LunaGO/server/messages"
	"errors"
	"fmt"
	"log"
	"net"
)

type Stub struct {
	// packets chan ([]byte)
	id         int32
	connection net.Conn
	handlers   map[int32]func([]byte)
}

func New(index int32) *Stub {
	instance := &Stub{
		// packets: make([]byte, 0, 300),
		id:       index,
		handlers: make(map[int32]func([]byte)),
	}

	return instance
}

func (stub *Stub) SetConnection(conn net.Conn) {
	stub.connection = conn
}

func (stub *Stub) Handle(code int32, handler func(packet []byte)) error {

	if stub.handlers[code] != nil {
		return errors.New(fmt.Sprintf("code:%d has been added.", code))
	}

	stub.handlers[code] = handler
	return nil
}

func (stub *Stub) getHandler(code int32) func([]byte) {
	return stub.handlers[code]
}

func (stub *Stub) Start() {
	c := stub.connection
	cIndex := stub.id
	for {
		var buf = make([]byte, 1024)
		log.Printf("start to read from conn[%d]\n", cIndex)
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("conn[%d] read %d bytes, content is %s\n", cIndex, n, buf[:n])
		message, err := messages.Unmarshal(buf[:n])
		if err != nil {
			log.Println("message unmarshal error:", err)
		}
		log.Println("code: ", message.Code)
		stub.getHandler(message.Code)(message.Data)
	}
}
