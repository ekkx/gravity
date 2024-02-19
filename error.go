package gravity

import "errors"

var (
	ErrInvalidIdentifier   = errors.New("not a valid identifier given")
	ErrStorageDoesNotMatch = errors.New("storage does not match given credentials")
)

var (
	ErrNoInvalidParams             = -3
	ErrNoSuccess                   = 0
	ErrNoEmailAddressOrPassword    = 50012
	ErrNoEmailAddressNotRegistered = 50014
)
