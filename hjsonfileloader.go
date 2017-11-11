package configuration

import (
	"io/ioutil"

	hjson "github.com/hjson/hjson-go"
)

// HJSONConfig is configuration loader HJSON interface
type HJSONConfig struct {
	filename string
	hjsonMap map[string]interface{}
}

// LoadFileContents load contents of file. separate function to make tests possible
func (fl *HJSONConfig) LoadFileContents(filename string) (cnt []byte, err error) {
	if filename == "" {
		return nil, NewConfigNotConfiguredError("Cannot load config file with no filename")
	}
	return ioutil.ReadFile(filename)
}

// ParseStringContents parses HJSON - separated to method cause I want test that
func (fl *HJSONConfig) ParseStringContents(cnt []byte) (m map[string]interface{}, err error) {
	m = map[string]interface{}{}
	err = hjson.Unmarshal(cnt, &m)
	if nil != err {
		m = nil
		return
	}
	return
}

// SetDefaultLoadSetting sets default config file for loader
func (fl *HJSONConfig) SetDefaultLoadSetting(sl ...interface{}) (err error) {
	if len(sl) == 0 {
		return NewHJSONConfigError("No arguments given to SetDefaultLoadSetting")
	}
	a0 := sl[0]
	switch v := a0.(type) {
	case string:
		cnt, err := fl.LoadFileContents(v)
		if err != nil {
			return NewHJSONConfigError("Error loading file occurred: " + err.Error())
		}
		m, err := fl.ParseStringContents(cnt)
		if err != nil {
			return err
		}
		fl.filename = v
		fl.hjsonMap = m
	case []byte:
		m, err := fl.ParseStringContents(v)
		fl.filename = ""
		if err != nil {
			return err
		}
		fl.hjsonMap = m
	case map[string]interface{}:
		fl.filename = ""
		fl.hjsonMap = v
	default:
		return NewHJSONConfigError("HJSONConfig.SetDefaultLoadSetting() argument must be string, []byte, or map[string]interface{}")
	}
	// all error cases are solved above
	return nil
}

// CheckExternalConfig checks external configuration file and it's contents - e.g.check file before reload
func (fl *HJSONConfig) CheckExternalConfig() (err error) {
	if "" == fl.filename {
		return NewConfigUsageError("Can not check external file cause it's not configured inside")
	}
	cnt, err := fl.LoadFileContents(fl.filename)
	if err != nil {
		return err
	}
	_, err = fl.ParseStringContents(cnt)
	if nil != err {
		return err
	}
	return nil
}

// ReloadInternalMap (re)loads internal map - if from file. If not - says ConfigUsageError
func (fl *HJSONConfig) ReloadInternalMap() (err error) {
	if "" == fl.filename {
		return NewConfigUsageError("Can not check external file cause it's not configured inside")
	}
	cnt, err := fl.LoadFileContents(fl.filename)
	if err != nil {
		return err
	}
	m, err := fl.ParseStringContents(cnt)
	if nil != err {
		return err
	}
	fl.hjsonMap = m
	return nil
}

// GetValue get any type value on programmer mind own
// usage on initialized object: fl.GetValue("a", "b", "c", "d")
// on this function would be based functions below
func (fl *HJSONConfig) GetValue(path ...string) (i interface{}, err error) {
	if nil == fl.hjsonMap {
		return nil, NewConfigUsageError("No config was initialized yet")
	}
	if 0 == len(path) {
		return nil, NewConfigUsageError("You must get values with path arguments there")
	}
	currentMap := fl.hjsonMap
	for _, key := range path {
		val, ok := currentMap[key]
		if !ok {
			return nil, NewConfigItemNotFound("Item not found")
		}
		switch v := val.(type) {
		case map[string]interface{}:
			currentMap = v
		default:
			return interface{}(v), nil
		}
	}
	// here stays 1 algorithmic variant - currentMap is answer itself
	return interface{}(currentMap), nil
}

// GetIntValue returns integer value by path
func (fl *HJSONConfig) GetIntValue(path ...string) (i int, err error) {
	i1, err1 := fl.GetValue(path...)
	if nil != err1 {
		return 0, err1
	}
	switch v := i1.(type) {
	case float64:
		return int(v), nil
	case int:
		return int(v), nil
	default:
		return 0, NewConfigTypeMismatchError("Wrong value type detected")
	}
}

// GetStringValue returns string value or error by path
func (fl *HJSONConfig) GetStringValue(path ...string) (s string, err error) {
	i1, err1 := fl.GetValue(path...)
	if nil != err1 {
		return "", err1
	}
	switch v := i1.(type) {
	case string:
		return v, nil
	default:
		return "", NewConfigTypeMismatchError("Wrong value type detected")
	}
}

// GetBooleanValue returns boolean value or error by path
func (fl *HJSONConfig) GetBooleanValue(path ...string) (b bool, err error) {
	i1, err1 := fl.GetValue(path...)
	if nil != err1 {
		return false, err1
	}
	switch v := i1.(type) {
	case bool:
		return v, nil
	default:
		return false, NewConfigTypeMismatchError("Wrong value type detected")
	}
}

// GetSubconfig returns config interface or nil + error
func (fl *HJSONConfig) GetSubconfig(path ...string) (c IConfig, err error) {
	i1, err1 := fl.GetValue(path...)
	if nil != err1 {
		return nil, err1
	}
	switch v := i1.(type) {
	case map[string]interface{}:
		return &HJSONConfig{filename: "", hjsonMap: v}, nil
	default:
		return nil, NewConfigTypeMismatchError("Wrong value type detected")
	}
}

// NewHJSONConfig creates new object or gives err0r
// all arguments are same as HJSONConfig.SetDefaultLoadSetting
func NewHJSONConfig(sl ...interface{}) (fl *HJSONConfig, err error) {
	fl = &HJSONConfig{}
	err = fl.SetDefaultLoadSetting(sl...)
	if err != nil {
		return nil, err
	}
	return
}
