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
	address := encrypt(email)

	_, err := s.g.requestWithForm("POST", EndpointEmailIsRegistered, &IsEmailRegisteredParams{Address: address})

	return err == nil
}

type IsPhoneNumberRegisteredParams struct {
	Number string `json:"pnum"`
}

func (s *UserService) isPhoneNumberRegistered(number string) bool {
	pnum := encrypt(number)

	_, err := s.g.requestWithForm("POST", EndpointMobileIsRegistered, &IsPhoneNumberRegisteredParams{Number: pnum})

	return err == nil
}

type LoginWithEmailParams struct {
	Address  string `json:"address"`
	Password string `json:"pwd"`
}

func (s *UserService) loginWithEmail(email, password string) (st LoginData, err error) {
	address := encrypt(email)
	pwd := encrypt(password)

	response, err := s.g.requestWithForm("POST", EndpointEmailLogin, &LoginWithEmailParams{Address: address, Password: pwd})
	if err != nil {
		return
	}

	err = unmarshal(response, &st)

	return
}
