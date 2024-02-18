package gravity

type Credentials struct {
	identifier string
	password   string
}

type DeviceInfo struct {
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
	credentials *Credentials
	deviceInfo  *DeviceInfo
	address     string
	pnum        string
	idfa        string
	uwd         string
	token       string
}

func NewState(identifier string, password string) *State {
	return &State{
		credentials: &Credentials{
			identifier: identifier,
			password:   password,
		},
		deviceInfo: &DeviceInfo{
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
		address: "",
		pnum:    "",
		idfa:    "",
		uwd:     "",
		token:   "",
	}
}
