package messages

import (
	"log"

	"github.com/vmihailenco/msgpack"
)

type Message struct {
	Code int32
	Data []byte
}

func Marshal(code int32, data []byte) ([]byte, error) {
	log.Printf("try to marshal: %d,%b\n", code, data)
	b, err := msgpack.Marshal(&Message{Code: code, Data: data})
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Unmarshal(data []byte) (*Message, error) {
	var message Message
	err := msgpack.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	return &message, err
}
