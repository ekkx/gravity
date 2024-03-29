package gravity

import (
	"net/http"
	"time"
)

// VERSION of Go Gravity
const VERSION = "0.0.0"

type Gravity struct {
	client *http.Client

	state            *State
	storageFilename  string
	retryOnRateLimit bool
	maxRestRetries   int

	Common  *CommonService
	Feed    *FeedService
	Storage *StorageService
	User    *UserService
}

type GravityConfig struct {
	client          *http.Client
	storageFilename string
}

func newGravityConfig() *GravityConfig {
	return &GravityConfig{
		client:          &http.Client{Timeout: 20 * time.Second},
		storageFilename: "secret.gob",
	}
}

type GravityOption func(cfg *GravityConfig)

// WithClient changes the HTTP client used for the request.
func WithClient(client *http.Client) GravityOption {
	return func(cfg *GravityConfig) {
		if client != nil {
			cfg.client = client
		}
	}
}

// WithStorageFilename changes the filename used for the storage.
func WithStorageFilename(filename string) GravityOption {
	return func(cfg *GravityConfig) {
		cfg.storageFilename = filename
	}
}

func New(identifier, password string, options ...GravityOption) (g *Gravity, err error) {
	cfg := newGravityConfig()
	for _, opt := range options {
		opt(cfg)
	}

	g = &Gravity{
		client:           cfg.client,
		state:            NewState(identifier, password, getLoginType(identifier)),
		storageFilename:  cfg.storageFilename,
		retryOnRateLimit: true,
		maxRestRetries:   3,
	}

	err = g.init()
	if err != nil {
		return nil, err
	}

	return
}

func (g *Gravity) init() error {
	g.Common = newCommonService(g)
	g.Feed = newFeedService(g)
	g.Storage = newStorageService(g)
	g.User = newUserService(g)

	// Initialize the storage
	err := g.Storage.Load()
	if err != nil {
		if err2 := g.Storage.CreateOneAndSave(); err2 != nil {
			return err2
		}
	}

	g.Storage.prepareState()

	if g.state.cred.Token == "" {
		token, err := g.authenticate()
		if err != nil {
			return err
		}
		g.state.cred.Token = token
	}

	return nil
}

// authenticate() authenticates the user with Gravity.State.credentials.
func (g *Gravity) authenticate() (string, error) {
	switch g.state.cred.LoginType {
	case LoginTypeEmail:
		isreg, err := g.User.isEmailRegistered(g.state.cred.Identifier)
		if err != nil {
			return "", err
		}

		if !isreg {
			return "", ErrAuthenticationFailed
		}

		resp, err := g.User.loginWithEmail(g.state.cred.Identifier, g.state.cred.Password)
		if err != nil {
			return "", err
		}

		return resp.Token, nil
	default:
		return "", ErrInvalidIdentifier
	}
}
