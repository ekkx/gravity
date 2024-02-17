package gravity

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	g := New("hello@example.com", "password")
	fmt.Println(g.State.credentials.identifier, g.State.credentials.password)
}
