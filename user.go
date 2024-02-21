package gravity

type UserService struct {
	g *Gravity
}

func newUserService(g *Gravity) *UserService {
	return &UserService{
		g: g,
	}
}

type IsEmailRegisteredParams struct {
	Address string `json:"address"`
}

func (s *UserService) isEmailRegistered(email string) bool {
	address, _ := encrypt(email)

	_, err := s.g.requestWithForm("POST", EndpointEmailIsRegistered, &IsEmailRegisteredParams{Address: address}, nil)

	return err == nil
}

type IsPhoneNumberRegisteredParams struct {
	Number string `json:"pnum"`
}

func (s *UserService) isPhoneNumberRegistered(number string) bool {
	pnum, _ := encrypt(number)

	_, err := s.g.requestWithForm("POST", EndpointMobileIsRegistered, &IsPhoneNumberRegisteredParams{Number: pnum}, nil)

	return err == nil
}

type LoginWithEmailParams struct {
	Address  string `json:"address"`
	Password string `json:"pwd"`
}

func (s *UserService) loginWithEmail(email, password string) (response interface{}, err error) {
	address, _ := encrypt(email)
	pwd, _ := encrypt(password)

	resp, err := s.g.requestWithForm("POST", EndpointEmailLogin, &LoginWithEmailParams{Address: address, Password: pwd}, nil)
	if err != nil {
		return
	}

	return resp, nil
}
