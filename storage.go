package gravity

import (
	"encoding/gob"
	"os"
	"strings"
)

type StorageService struct {
	g *Gravity
}

func newStorageService(g *Gravity) *StorageService {
	return &StorageService{
		g: g,
	}
}

func readStorage(filename string) (c *Credentials, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	c = &Credentials{}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		return nil, err
	}

	return
}

func writeStorage(filename string, cred *Credentials) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(cred)
	if err != nil {
		return
	}

	return
}

func deleteStorage(filename string) (err error) {
	err = os.Remove(filename)
	return err
}

// load() transports local storage data into Gravity.State.
// If Gravity.State doesn't match the local straage, returns an error.
func (s *StorageService) load(filename string) (err error) {
	c, err := readStorage(filename)
	if err != nil {
		return
	}

	if !(c.Identifier == s.g.State.cred.Identifier && c.Password == s.g.State.cred.Password) {
		return ErrStorageDoesNotMatch
	}

	s.g.State.cred = c

	return
}

// save() exports current Gravity.State as local storage data
func (s *StorageService) save(filename string) error {
	// Check idtype just in case.
	idtype := getIDType(s.g.State.cred.Identifier)
	if idtype == -1 {
		return ErrInvalidIdentifier
	}

	return writeStorage(filename, s.g.State.cred)
}

func (s *StorageService) createOneAndSave(filename string) error {
	gaid, _ := generateUUID()
	uuid, _ := generateUUID()

	s.g.State.cred.GAID = gaid
	s.g.State.cred.GAID = strings.ToUpper(uuid)

	return s.save(filename)
}
