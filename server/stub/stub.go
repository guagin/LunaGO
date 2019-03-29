package stub

import (
	"LunaGO/server/conn"
	"errors"
	"fmt"
	"log"
)

type Stub struct {
	Packets    chan ([]byte)
	id         int32
	connection *conn.Connection
	handlers   map[int32]func([]byte) []byte
	Process    func(chan<- bool)
}

func New(index int32) *Stub {
	instance := &Stub{
		Packets:  make(chan []byte, 300),
		id:       index,
		handlers: make(map[int32]func([]byte) []byte),
	}

	return instance
}

func (stub *Stub) Handle(code int32, handler func(packet []byte) []byte) error {

	if stub.handlers[code] != nil {
		return errors.New(fmt.Sprintf("code:%d has been added.", code))
	}

	stub.handlers[code] = handler
	return nil
}

func (stub *Stub) GetHandler(code int32) func([]byte) []byte {
	return stub.handlers[code]
}

func (stub *Stub) SetConnection(c *conn.Connection) {
	stub.connection = c
}

func (stub *Stub) Start() {
	quit := make(chan bool)
	go stub.connection.StartReceiving(stub.Packets)
	go stub.processPacket(quit)
	if <-quit {
		log.Println("client send close event.")
	}
}

func (stub *Stub) processPacket(quit chan<- bool) {
	for {
		stub.Process(quit)
		// packet := <-stub.packets
		// message, err := messages.Unmarshal(packet)
		// if err != nil {
		// 	log.Println("message unmarshal error:", err)
		// 	return
		// }
		// log.Println("code: ", message.Code)
		// if message.Code == -1 {
		// 	quit <- true
		// 	return
		// }
		// resData := stub.getHandler(message.Code)(message.Data, stub.connection)

		// resInBytes, err := messages.Marshal(message.Code, resData)
		// if err != nil {
		// 	log.Println("message marshal error:", err)
		// 	return
		// }
		// stub.connection.Write(resInBytes)

	}
}
