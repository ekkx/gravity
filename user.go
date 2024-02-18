package gravity

import "fmt"

type UserService struct {
	g *Gravity
}

func newUserService(g *Gravity) *UserService {
	return &UserService{
		g: g,
	}
}

func (s *UserService) IsRegisteredEmail(email string) bool {
	fmt.Println("IsRegisteredEmail", email)
	return false
}
