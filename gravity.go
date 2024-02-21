package gravity

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// VERSION of Go Gravity
const VERSION = "0.0.0"

type Gravity struct {
	sling *sling.Sling

	State                  *State
	ShouldRetryOnRateLimit bool
	MaxRestRetries         int

	Common  *CommonService
	Storage *StorageService
	User    *UserService
}

type GravityConfig struct {
	Client *http.Client
}

func newGravityConfig() *GravityConfig {
	return &GravityConfig{
		Client: &http.Client{Timeout: 20 * time.Second},
	}
}

type GravityOption func(cfg *GravityConfig)

// WithClient changes the HTTP client used for the request.
func WithClient(client *http.Client) GravityOption {
	return func(cfg *GravityConfig) {
		if client != nil {
			cfg.Client = client
		}
	}
}

func New(identifier string, password string, options ...GravityOption) (g *Gravity, err error) {
	cfg := newGravityConfig()
	for _, opt := range options {
		opt(cfg)
	}

	g = &Gravity{
		sling:                  sling.New().Client(cfg.Client).Base(EndpointRoot),
		State:                  NewState(identifier, password, getIDType(identifier)),
		ShouldRetryOnRateLimit: true,
		MaxRestRetries:         3,
	}

	err = g.init()
	if err != nil {
		return nil, err
	}

	return
}

func (g *Gravity) init() (err error) {
	g.Common = newCommonService(g)
	g.User = newUserService(g)

	g.Storage = newStorageService(g)

	// Initialize the storage
	err = g.Storage.load()
	if err != nil {
		err2 := g.Storage.createOneAndSave()
		if err2 != nil {
			return err2
		}
	}

	token, err := g.authenticate()
	if err != nil {
		return
	}
	g.State.token = token

	return
}

// authenticate() authenticates the user with Gravity.State.credentials.
func (g *Gravity) authenticate() (token string, err error) {
	switch g.State.cred.idtype {
	case 0:
		if !(g.User.isEmailRegistered(g.State.cred.identifier)) {
			return "", ErrAuthenticationFailed
		}
		resp, err := g.User.loginWithEmail(g.State.cred.identifier, g.State.cred.pwd)
		if err != nil {
			return "", err
		}

		return resp.token, nil
	// case 1:
	// 	// Login with phone number
	// 	if !(g.User.isPhoneNumberRegistered(g.State.cred.identifier)) {
	// 		return "", ErrAuthenticationFailed
	// 	}
	// 	resp, err := g.User.loginWithPhoneNumber(g.State.cred.identifier, g.State.cred.pwd)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	return resp.token, nil
	default:
		return "", ErrInvalidIdentifier
	}
}
