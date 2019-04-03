package interfaces

//Stub handle the message from client.
type Stub interface {
	New() Stub
	ID() int32
	Handle(int32, func([]byte) []byte)
	GetHandler(code int32) func([]byte) []byte
	Start()
	Send(packet []byte)
}
