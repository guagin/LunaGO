package interfaces

type Server interface {
	GetStub() Stub
	Dependancy() int32
}
