package gravity

type Credentials struct {
	LoginType  int
	Identifier string
	Password   string
	GAID       string
	UUID       string
	Token      string
}

type Device struct {
	SimCountry     string `json:"sim_country"`
	Device         string `json:"device"`
	SystemVersion  string `json:"system_version"`
	Sign           string `json:"sign"`
	Referrer       string `json:"referrer"`
	Zone           string `json:"zone"`
	IDFA           string `json:"idfa"`
	AppVersionCode string `json:"app_version_code"`
	Timestamp      string `json:"ts"`
	SystemLanguage string `json:"sys_lang"`
	AppVersion     string `json:"app_version"`
	LanguageV2     string `json:"languageV2"`
	UWD            string `json:"uwd"`
	Country        string `json:"country"`
	Brand          string `json:"brand"`
	SDKVersion     string `json:"sdk_version"`
	UserCountry    string `json:"user_country"`
	Package        string `json:"pkg"`
	Product        string `json:"product"`
	Model          string `json:"model"`
}

type State struct {
	cred   *Credentials
	device *Device
}

func NewState(identifier string, password string, loginType int) *State {
	return &State{
		cred: &Credentials{
			LoginType:  loginType,
			Identifier: identifier,
			Password:   password,
			GAID:       "",
			UUID:       "",
			Token:      "",
		},
		device: &Device{
			SimCountry:     "JP",
			Device:         "android",
			SystemVersion:  "7.1.2",
			Sign:           "",
			Referrer:       "Organic",
			Zone:           "9",
			IDFA:           "",
			AppVersionCode: "375",
			Timestamp:      "",
			SystemLanguage: "ja",
			AppVersion:     "9.2.0",
			LanguageV2:     "ja",
			UWD:            "",
			Country:        "JP",
			Brand:          "samsung",
			SDKVersion:     "25",
			UserCountry:    "JP",
			Package:        "anonymous.sns.community.gravity",
			Product:        "gravity",
			Model:          "SM-G965N",
		},
	}
}
