package gravity

type Credentials struct {
	idtype     int
	identifier string
	pwd        string
	gaid       string
	uuid       string
}

type Device struct {
	simCountry  string
	device      string
	sysVer      string
	referrer    string
	zone        string
	appVerCode  string
	sysLang     string
	appVer      string
	langV2      string
	country     string
	brand       string
	sdkVer      string
	userCountry string
	pkg         string
	product     string
	model       string
}

type State struct {
	cred   *Credentials
	device *Device
	token  string
}

func NewState(identifier string, password string, idtype int) *State {
	return &State{
		cred: &Credentials{
			idtype:     idtype,
			identifier: identifier,
			pwd:        password,
			gaid:       "",
			uuid:       "",
		},
		device: &Device{
			simCountry:  "JP",
			device:      "android",
			sysVer:      "7.1.2",
			referrer:    "Organic",
			zone:        "9",
			appVerCode:  "375",
			sysLang:     "ja",
			appVer:      "9.2.0",
			langV2:      "ja",
			country:     "JP",
			brand:       "samsung",
			sdkVer:      "25",
			userCountry: "JP",
			pkg:         "anonymous.sns.community.gravity",
			product:     "gravity",
			model:       "SM-G965N",
		},
		token: "",
	}
}
