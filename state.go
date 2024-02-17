package gravity

type credentials struct {
	identifier string
	password   string
}

type State struct {
	credentials *credentials
	token       string
}

func NewState(identifier string, password string) *State {
	return &State{
		credentials: &credentials{
			identifier: identifier,
			password:   password,
		},
		token: "",
	}
}
