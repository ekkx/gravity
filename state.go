package gravity

type State struct {
	credentials *credentials
}

type credentials struct {
	identifier string
	password   string
}

func NewState(identifier string, password string) *State {
	return &State{
		credentials: &credentials{
			identifier: identifier,
			password:   password,
		},
	}
}
