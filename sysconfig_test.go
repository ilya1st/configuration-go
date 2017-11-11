package configuration

import (
	"reflect"
	"testing"
)

func TestGetConfigInstance(t *testing.T) {
	type args struct {
		settings []interface{}
	}
	type teststruct struct {
		name        string
		args        args
		wantConfig  IConfig
		wantErr     bool
		wantErrType string
		checker     func(fl IConfig) bool
	}
	tests := []teststruct{
		teststruct{
			name: "No arguments fail",
			args: args{
				settings: []interface{}{},
			},
			wantConfig:  nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
			checker:     nil,
		},
		teststruct{
			name: "nil first arg anf not existing config class second argument",
			args: args{
				settings: []interface{}{nil, "non_existing_format"},
			},
			wantConfig:  nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
			checker:     nil,
		},
		teststruct{
			name: "nil first, correct HJSON second format, but no other arguments - fail init HJSONFileLoader class",
			args: args{
				settings: []interface{}{nil, "HJSON"},
			},
			wantConfig: nil,
			wantErr:    true,
			// want from here with no args func (fl *HJSONFileLoader) SetDefaultLoadSetting(sl ...interface{}) (err error)
			wantErrType: "*configuration.HJSONConfigError",
			checker:     nil,
		},
		// here we just test that load works
		teststruct{
			name: "nil first, HJSON, correct second arg - hashmap. must get correct HJSONFileLoader",
			args: args{
				settings: []interface{}{nil, "HJSON", map[string]interface{}{"test": "test"}},
			},
			wantConfig:  &HJSONConfig{filename: "", hjsonMap: map[string]interface{}{"test": "test"}},
			wantErr:     false,
			wantErrType: "",
			checker:     nil,
		},
		teststruct{
			name: "check how tags do their work",
			args: args{
				settings: []interface{}{"testtag", "HJSON", map[string]interface{}{"test": "test"}},
			},
			wantConfig:  &HJSONConfig{filename: "", hjsonMap: map[string]interface{}{"test": "test"}},
			wantErr:     false,
			wantErrType: "",
			checker: func(fl IConfig) bool {
				fl1, err := GetConfigInstance("testtag")
				if err != nil {
					t.Errorf(
						"GetConfigInstance(): with tag returned error %v %v while try retrieve cached instance",
						reflect.TypeOf(err),
						err,
					)
					return false
				}
				if fl1 != fl {
					t.Errorf(
						"GetConfigInstance(): cached by tag function returns different instance",
					)
					return false
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfig, err := GetConfigInstance(tt.args.settings...)
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("GetConfigInstance()() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfigInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && ("" != tt.wantErrType) && (tt.wantErrType != reflect.TypeOf(err).String()) {
				t.Errorf("HJSONConfig.LoadFileContents() error type = %v, wantErrType %v", reflect.TypeOf(err), tt.wantErrType)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("GetConfigInstance() = %v, want %v", gotConfig, tt.wantConfig)
				return
			}
			if nil != tt.checker {
				tt.checker(gotConfig)
			}
		})
	}
}
