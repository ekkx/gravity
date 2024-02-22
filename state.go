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
	SimCountry     string `json:"sim_country,omitempty"`
	Device         string `json:"device,omitempty"`
	SystemVersion  string `json:"system_version,omitempty"`
	Sign           string `json:"sign,omitempty"`
	Referrer       string `json:"referrer,omitempty"`
	Zone           string `json:"zone,omitempty"`
	IDFA           string `json:"idfa,omitempty"`
	AppVersionCode string `json:"app_version_code,omitempty"`
	Timestamp      string `json:"ts,omitempty"`
	SystemLanguage string `json:"sys_lang,omitempty"`
	AppVersion     string `json:"app_version,omitempty"`
	LanguageV2     string `json:"languageV2,omitempty"`
	UWD            string `json:"uwd,omitempty"`
	Country        string `json:"country,omitempty"`
	Brand          string `json:"brand,omitempty"`
	SDKVersion     string `json:"sdk_version,omitempty"`
	UserCountry    string `json:"user_country,omitempty"`
	Package        string `json:"pkg,omitempty"`
	Product        string `json:"product,omitempty"`
	Model          string `json:"model,omitempty"`
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
