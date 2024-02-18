package gravity

import (
	"testing"
)

func TestMain(t *testing.T) {
	g := New("hello@example.com", "password")
	g.User.IsRegisteredEmail(g.State.credentials.identifier)
}
