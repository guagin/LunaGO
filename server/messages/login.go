package messages

import "github.com/vmihailenco/msgpack"

type Login struct {
	ID string
}

func (login *Login) Marshal() ([]byte, error) {
	b, err := msgpack.Marshal(&Login{ID: login.ID})
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UnmarshalLogin(data []byte) (*Login, error) {
	var login Login
	err := msgpack.Unmarshal(data, &login)
	if err != nil {
		return nil, err
	}
	return &login, err
}
