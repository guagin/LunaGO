package interfaces

type Server interface {
	GetStub() *Stub
	Dependancy() int32
	Get(name string) (interface{}, error)
	Register(name string, obj interface{})
}
