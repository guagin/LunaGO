package stub

import (
	"LunaGO/server/messages"
	"errors"
	"fmt"
	"log"
	"net"
)

type Stub struct {
	packets    chan ([]byte)
	id         int32
	connection net.Conn
	handlers   map[int32]func([]byte) []byte
}

func New(index int32) *Stub {
	instance := &Stub{
		packets:  make(chan []byte, 300),
		id:       index,
		handlers: make(map[int32]func([]byte) []byte),
	}

	return instance
}

func (stub *Stub) SetConnection(conn net.Conn) {
	stub.connection = conn
}

func (stub *Stub) Handle(code int32, handler func(packet []byte) []byte) error {

	if stub.handlers[code] != nil {
		return errors.New(fmt.Sprintf("code:%d has been added.", code))
	}

	stub.handlers[code] = handler
	return nil
}

func (stub *Stub) getHandler(code int32) func([]byte) []byte {
	return stub.handlers[code]
}

func (stub *Stub) Start() {
	quit := make(chan bool)
	go stub.startReadFrom(stub.connection)
	go stub.processPacket(quit)
	if <-quit {
		log.Println("client send close event.")
	}
}

func (stub *Stub) startReadFrom(c net.Conn) {
	for {
		var buf = make([]byte, 1024)
		log.Printf("start to read from conn[%d]\n", stub.id)
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("conn[%d] read %d bytes, content is %s\n", stub.id, n, buf[:n])
		stub.packets <- buf[:n]
	}
}

func (stub *Stub) processPacket(quit chan<- bool) {
	for {
		packet := <-stub.packets
		message, err := messages.Unmarshal(packet)
		if err != nil {
			log.Println("message unmarshal error:", err)
			return
		}
		log.Println("code: ", message.Code)
		if message.Code == -1 {
			quit <- true
			return
		}
		resData := stub.getHandler(message.Code)(message.Data)

		resInBytes, err := messages.Marshal(message.Code, resData)
		if err != nil {
			log.Println("message marshal error:", err)
			return
		}
		stub.connection.Write(resInBytes)

	}

}
