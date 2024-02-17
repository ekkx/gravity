package gravity

import "fmt"

type CommonService struct {
	g *Gravity
}

func newCommonService(g *Gravity) *CommonService {
	return &CommonService{g: g}
}

func (s *CommonService) SaySomething(msg string) {
	g := s.g
	fmt.Println(g.State.credentials.identifier, msg)
}
