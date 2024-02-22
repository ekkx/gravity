package gravity

import "fmt"

type CommonService struct {
	g *Gravity
}

func newCommonService(g *Gravity) *CommonService {
	return &CommonService{
		g: g,
	}
}

func (s *CommonService) SaySomething(msg string) {
	fmt.Println(s.g.State.cred.Identifier, msg)
}
