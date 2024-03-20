package gravity

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

var (
	envEmail    = os.Getenv("TEST_GRAVITY_EMAIL")
	envPassword = os.Getenv("TEST_GRAVITY_PASSWORD")
)

func TestGravity(t *testing.T) {
	g, err := New(envEmail, envPassword, WithClient(&http.Client{}))
	if err != nil {
		fmt.Println(err)
		return
	}

	feed, err := g.Feed.RecommendFeedList(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(feed)
}
