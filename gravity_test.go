package gravity

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

var (
	envEmail    string
	envPassword string
)

func init() {
	readEnv()
	envEmail = os.Getenv("TEST_EMAIL")
	envPassword = os.Getenv("TEST_PASSWORD")
}

func TestGravity(t *testing.T) {
	g, err := New(envEmail, envPassword, WithClient(&http.Client{}))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(g.state.cred)
}
