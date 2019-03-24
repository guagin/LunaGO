package player

type player struct {
	ID string
}

func New() *player {
	instance := &player{
		ID: "",
	}
	return instance
}
