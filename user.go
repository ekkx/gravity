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

func (s *UserService) isEmailRegistered(email string) (isreg bool, err error) {
	address := encrypt(email)

	resp, err := s.g.requestWithForm("POST", EndpointUserEmailIsRegistered, &IsEmailRegisteredParams{Address: address})
	if err != nil {
		if resp.ErrNo != ErrNoEmailAddressOrPassword {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

type IsPhoneNumberRegisteredParams struct {
	Number string `json:"pnum"`
}

func (s *UserService) isPhoneNumberRegistered(number string) bool {
	pnum := encrypt(number)

	// isEmailRegistered のように修正する
	_, err := s.g.requestWithForm("POST", EndpointUserMobileIsRegistered, &IsPhoneNumberRegisteredParams{Number: pnum})

	return err == nil
}

type LoginWithEmailParams struct {
	Address  string `json:"address"`
	Password string `json:"pwd"`
}

func (s *UserService) loginWithEmail(email, password string) (st *LoginData, err error) {
	address := encrypt(email)
	pwd := encrypt(password)

	resp, err := s.g.requestWithForm("POST", EndpointUserEmailLogin, &LoginWithEmailParams{Address: address, Password: pwd})
	if err != nil {
		return
	}

	err = unmarshal(resp, &st)

	return
}
