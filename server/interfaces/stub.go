package interfaces

//Stub handle the message from client.
type Stub interface {
	New() Stub
	ID() string
	Handle(int32, func([]byte) []byte)
	GetHandler(code int32) func([]byte) []byte
	Start()
	Send(packet []byte)
}
