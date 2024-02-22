package gravity

type Credentials struct {
	IdentifierType int
	Identifier     string
	Password       string
	GAID           string
	UUID           string
}

type Device struct {
	simCountry  string `json:"sim_country"`
	device      string `json:"device"`
	sysVer      string `json:"system_version"`
	sign        string `json:"sign"`
	referrer    string `json:"referrer"`
	zone        string `json:"zone"`
	idfa        string `json:"idfa"`
	appVerCode  string `json:"app_version_code"`
	timestamp   string `json:"ts"`
	sysLang     string `json:"sys_lang"`
	appVer      string `json:"app_version"`
	langV2      string `json:"languageV2"`
	uwd         string `json:"uwd"`
	country     string `json:"country"`
	brand       string `json:"brand"`
	sdkVer      string `json:"sdk_version"`
	userCountry string `json:"user_country"`
	pkg         string `json:"pkg"`
	product     string `json:"product"`
	model       string `json:"model"`
}

type State struct {
	cred   *Credentials
	device *Device
	token  string
}

func NewState(identifier string, password string, idtype int) *State {
	return &State{
		cred: &Credentials{
			IdentifierType:     idtype,
			Identifier: identifier,
			Password:        password,
			GAID:       "",
			UUID:       "",
		},
		device: &Device{
			simCountry:  "JP",
			device:      "android",
			sysVer:      "7.1.2",
			sign:        "",
			referrer:    "Organic",
			zone:        "9",
			idfa:        "",
			appVerCode:  "375",
			timestamp:   "",
			sysLang:     "ja",
			appVer:      "9.2.0",
			langV2:      "ja",
			uwd:         "",
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
