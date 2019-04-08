package stub

import (
	"LunaGO/server/conn"
	"log"
)

// Stub to keep the infomation like connection..etc.
type Stub struct {
	inPackets  chan ([]byte)
	id         string
	connection *conn.Connection
	handlers   map[int32]func([]byte) []byte
	process    func([]byte)
	outPackets chan []byte
}

func New(id string) *Stub {
	instance := &Stub{
		inPackets:  make(chan []byte, 300),
		id:         id,
		handlers:   make(map[int32]func([]byte) []byte),
		outPackets: make(chan []byte, 300),
	}

	return instance
}

func (stub *Stub) ID() string {
	return stub.id
}

func (stub *Stub) Handle(code int32, handler func(packet []byte) []byte) {

	if stub.handlers[code] != nil {
		// return
		return
	}

	stub.handlers[code] = handler
	return
}

func (stub *Stub) GetHandler(code int32) func([]byte) []byte {
	return stub.handlers[code]
}

func (stub *Stub) SetConnection(c *conn.Connection) {
	stub.connection = c
}

func (stub *Stub) SetProcess(process func([]byte)) {
	stub.process = process
}

func (stub *Stub) Start() {
	quit := make(chan bool)
	go stub.connection.StartReceiving(stub.inPackets)
	go stub.processPacket(quit)
	go stub.processResponses()
	if <-quit {
		log.Println("client send close event.")
	}
}

func (stub *Stub) processPacket(quit chan<- bool) {
	for {
		if stub.process == nil {
			log.Println("u have set process method.")
			return
		}
		packet, ok := <-stub.inPackets
		if !ok {
			log.Println("packet channel closed.")
			close(stub.outPackets)
			quit <- true
			return
		}
		stub.process(packet)

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

func (stub *Stub) processResponses() {
	for {
		response, ok := <-stub.outPackets
		if !ok {
			log.Println("response channel closed.")
			return
		}
		stub.connection.SendBytes(response)
	}
}

func (stub *Stub) Send(packet []byte) {
	if packet != nil {
		stub.outPackets <- packet
	}
}
