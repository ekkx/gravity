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
func (s *StorageService) Load() (err error) {
	c, err := readStorage(s.g.storageFilename)
	if err != nil {
		return
	}

	if !(c.Identifier == s.g.state.cred.Identifier && c.Password == s.g.state.cred.Password) {
		return ErrStorageDoesNotMatch
	}

	s.g.state.cred = c

	return
}

// save() exports current Gravity.State as local storage data
func (s *StorageService) Save() error {
	// Check login type just in case.
	ltype := getLoginType(s.g.state.cred.Identifier)
	if ltype == -1 {
		return ErrInvalidIdentifier
	}

	return writeStorage(s.g.storageFilename, s.g.state.cred)
}

func (s *StorageService) CreateOneAndSave() error {
	gaid, _ := generateUUID()
	uuid, _ := generateUUID()

	s.g.state.cred.GAID = gaid
	s.g.state.cred.UUID = strings.ToUpper(uuid)

	return s.Save()
}

func (s *StorageService) Remove() error {
	return deleteStorage(s.g.storageFilename)
}
