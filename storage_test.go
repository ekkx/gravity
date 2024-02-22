package gravity

import (
	"errors"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestWriteReadAndDeleteStorage(t *testing.T) {
	filename := "test.gob"

	cred := &Credentials{
		LoginType:  2,
		Identifier: "hello@example.com",
		Password:   "notpwd",
		GAID:       "",
		UUID:       "",
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

func TestStorageService(t *testing.T) {
	filename := "test.gob"

	g := &Gravity{
		client:           &http.Client{},
		state:            NewState("hello@example.com", "notpwd", 0),
		storageFilename:  filename,
		retryOnRateLimit: true,
		maxRestRetries:   3,
	}

	s := newStorageService(g)

	err := s.CreateOneAndSave()
	if err != nil {
		t.Fatal(err)
	}

	credBefore := g.state.cred

	t.Log(credBefore)

	err = s.Load()
	if err != nil {
		t.Fatalf("failed to load: %v", err)
	}

	if !reflect.DeepEqual(credBefore, g.state.cred) {
		t.Fatal("failed to load storage")
	}

	err = s.Remove()
	if err != nil {
		t.Fatalf("failed to remove: %v", err)
	}

	if _, err := os.Stat(filename); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("failed to delete storage")
	}
}
