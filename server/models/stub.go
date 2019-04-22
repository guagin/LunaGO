package models

import (
	"github.com/guagin/LunaGo/server/stub"
	"errors"
	"fmt"
	"log"
	"sync"
)

type stubRepository struct {
	Stubs map[string]*stub.Stub
	tasks chan func()
}

var newRepositoryOnce sync.Once
var stubRepo *stubRepository

func StubRepository() *stubRepository {
	newRepositoryOnce.Do(func() {
		newStubRepo()
		go stubRepo.process()
	}) // make sure repository only get one instance.
	return stubRepo
}

func newStubRepo() {
	stubRepo = &stubRepository{
		Stubs: make(map[string]*stub.Stub),
		tasks: make(chan func(), 100),
	}
	log.Println("new a stub repository ")
}

func (stubRepo *stubRepository) process() {
	for {
		task, ok := <-stubRepo.tasks
		if !ok {
			log.Println("stubRepo tasks channel has been closed.")
			return
		}
		task()
	}
}

func (stubRepo *stubRepository) Register(stub *stub.Stub) {
	stubRepo.tasks <- func() {
		stubRepo.Stubs[stub.ID()] = stub
	}
}

func (stubRepo *stubRepository) UnRegister(ID string) {
	stubRepo.tasks <- func() {
		delete(stubRepo.Stubs, ID)
	}
}

func (stubRepo *stubRepository) Get(ID string) (*stub.Stub, error) {
	stub := stubRepo.Stubs[ID]
	if stub == nil {
		return nil, errors.New(fmt.Sprintf("stub(%d) is not exist", ID))
	}
	return stub, nil
}
