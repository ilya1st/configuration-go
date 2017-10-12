package configuration

// IConfig is basic interface for all optional configuration structures
type IConfig interface {
	// Set default config load settings. In case of file loader - filename
	// if first argument is map object -build config from this map object
	SetDefaultLoadSetting(sl ...interface{}) (err error)
	// checks external configuration file and it's contents - for example for reload or restart purposes
	CheckExternalConfig() (err error)
	// (re)loads internal map
	ReloadInternalMap() (err error)
	// get variant type value
	GetValue(path ...string) (i interface{}, err error)
	// functions below will try make from this variant value typed value
	// returns integer value by path
	GetIntValue(path ...string) (i int, err error)
	// returns string value or error by path
	GetStringValue(path ...string) (s string, err error)
	// returns boolean value
	GetBooleanValue(path ...string) (b bool, err error)
	// returns config interface or nil + error
	GetSubconfig(path ...string) (c IConfig, err error)
}

// this map is intended for GetConfigInstance
var intConfigHash map[string]IConfig

func init() {
	intConfigHash = map[string]IConfig{}
}

// GetConfigInstance get instans of config depending of wanted config format
// for now  supported only HJSON format
// tag is intended not to load configuration files twice and store object hash map inside module
// thats done specially not to create config instance twice + get already created instance
// How to call: GetConfigInstance(tag, format, list_of_settings...)
// tag used not to load config twice. If you need reload them just use CheckExternalConfig() and ReloadInternalMap()
// format: for now HJSON|hjson|JSON|json
// this parameter added for future features(may be we woould add other configuration formats there)
// if it was loaded from file.
// settings must contain config type as first argument and other perameners SetDefaultLoadSetting need
// For example.
// First call when we want load configuration file is:
// GetConfigInstance("mainconfig", "HSON", "/etc/file.hjson")
// When you need get instance again, you just must call:
// GetConfigInstance("mainconfig") and you will get them from internal hash
// this all is intended not to put variable withyou configuration everywhere.
// if you do not want tagging and every time get new object, use nil tag
// e.g. GetConfigInstance(nil, "HSON", "/etc/file.hjson")
func GetConfigInstance(settings ...interface{}) (config IConfig, err error) {
	if len(settings) == 0 {
		return nil, NewConfigUsageError("Wrong usage")
	}
	hasTag := false
	tag := ""
	tmp := interface{}(nil)
	if len(settings) >= 1 {
		tmp = settings[0]
		if tmp == nil {
			hasTag = false
		} else {
			switch v := tmp.(type) {
			case string:
				hasTag = true
				tag = v
			default:
				return nil, NewConfigUsageError("Tag argument must be nil or strig type")
			}
		}
	}
	if hasTag {
		ic, ok := intConfigHash[tag]
		if ok {
			return ic, nil
		}
	}
	if len(settings) < 2 {
		return nil, NewConfigUsageError("While first tagged extraction we do need type of file")
	}
	tmp = settings[1]
	cType := ""
	switch v := tmp.(type) {
	case string:
		cType = v
	default:
		return nil, NewConfigUsageError("Wrong format of type of config")
	}
	switch cType {
	case "HJSON":
		fallthrough
	case "hjson":
		fallthrough
	case "JSON":
		fallthrough
	case "json":
		// all other arguments to constructor
		config, err = NewHJSONConfig(settings[2:]...)
		if err != nil {
			return nil, err
		}
		if hasTag {
			intConfigHash[tag] = config
		}
		return config, nil
	default:
		return nil, NewConfigUsageError("Unknown configuration format:" + cType)
	}
}
