package gravity

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestWriteReadAndDeleteStorage(t *testing.T) {
	filename := "test.gob"

	cred := &Credentials{
		IdentifierType: 0,
		Identifier:     "hello@example.com",
		Password:       "notpwd",
		GAID:           "",
		UUID:           "",
	}

	err := writeStorage(filename, cred)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		t.Fatal("failed to write storage")
	}

	readCred, err := readStorage(filename)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(cred, readCred) {
		t.Fatal("failed to read storage")
	}

	err = deleteStorage(filename)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filename); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("failed to delete storage")
	}
}
