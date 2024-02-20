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

	Storage *StorageService

	Common *CommonService
	User   *UserService
}

func New(identifier string, password string) (g *Gravity, err error) {
	g = &Gravity{
		State:                  NewState(identifier, password, getIDType(identifier)),
		ShouldRetryOnRateLimit: true,
		MaxRestRetries:         3,
		Client:                 &http.Client{Timeout: 20 * time.Second},
	}

	err = g.init()
	if err != nil {
		return nil, err
	}

	return g, nil
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
		// Login with email address
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
