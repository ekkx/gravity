package gravity

import (
	"fmt"
	"testing"
)

func TestIsRegisteredEmail(t *testing.T) {
	g, _ := New("example@gmail.com", "password")

	isreg := g.User.isEmailRegistered("example@gmail.com")
	fmt.Println(isreg)
}
