package gravity

import (
	"net/http"
	"time"
)

// VERSION of Go Gravity
const VERSION = "0.0.0"

type Gravity struct {
	State                  *State
	ShouldRetryOnRateLimit bool
	MaxRestRetries         int
	Client                 *http.Client

	Common *CommonService
	User   *UserService
}

func New(identifier string, password string) *Gravity {
	g := &Gravity{
		State:                  NewState(identifier, password),
		ShouldRetryOnRateLimit: true,
		MaxRestRetries:         3,
		Client:                 &http.Client{Timeout: 20 * time.Second},
	}
	g.init()

	return g
}

func (g *Gravity) init() {
	g.Common = newCommonService(g)
	g.User = newUserService(g)
}
