package gravity

import (
	"net/http"
	"time"
)

// VERSION of Go Gravity
const VERSION = "0.0.0"

type Gravity struct {
	State  *State
	Client *http.Client
}

func New(identifier string, password string) *Gravity {
	g := &Gravity{
		State:  NewState(identifier, password),
		Client: &http.Client{Timeout: 20 * time.Second},
	}
	return g
}

func (g *Gravity) SetHTTPClient(client *http.Client) {
	g.Client = client
}
