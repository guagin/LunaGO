package models

import (
	"LunaGO/server/stub"
	"errors"
	"fmt"
	"log"
	"sync"
)

type stubRepository struct {
	Stubs map[int32]*stub.Stub
}

var newRepositoryOnce sync.Once
var stubRepo *stubRepository

func StubRepository() *stubRepository {
	newRepositoryOnce.Do(newStubRepo) // make sure repository only get one instance.
	return stubRepo
}

func newStubRepo() {
	stubRepo = &stubRepository{
		Stubs: make(map[int32]*stub.Stub),
	}
	log.Println("new a stub repository ")
}

func (stubRepo *stubRepository) Register(ID int32, stub *stub.Stub) {
	//TODO: add mux lock
	stubRepo.Stubs[ID] = stub
}

func (stubRepo *stubRepository) UnRegister(ID int32) {
	//TODO: add mux lock
	delete(stubRepo.Stubs, ID)
}

func (stubRepo *stubRepository) Get(ID int32) (*stub.Stub, error) {
	stub := stubRepo.Stubs[ID]
	if stub == nil {
		return nil, errors.New(fmt.Sprintf("stub(%d) is not exist", ID))
	}
	return stub, nil
}
