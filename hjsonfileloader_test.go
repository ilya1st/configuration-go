package configuration

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestHJSONConfig_LoadFileContents(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type args struct {
		filename string
	}
	type teststruct struct {
		name        string
		fields      fields
		args        args
		wantCnt     []byte
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{
		{
			name:        "wrong file arg",
			fields:      fields{filename: "", hjsonMap: nil},
			args:        args{filename: ""},
			wantCnt:     nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigNotConfiguredError",
		},
		{
			name:        "file reading error",
			fields:      fields{filename: "", hjsonMap: nil},
			args:        args{filename: "/hjson_not_existing_awful_filename"},
			wantCnt:     nil,
			wantErr:     true,
			wantErrType: "",
		},
		{
			name:   "normal file open and read",
			fields: fields{filename: "", hjsonMap: nil},
			args:   args{filename: "./test.hjson"},
			wantCnt: []byte(`{
# this file it intended for testing purposes. Do not change them
"CONFIG_FILE":"./configuration/config.hjson"
}`),
			wantErr:     false,
			wantErrType: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			gotCnt, err := fl.LoadFileContents(tt.args.filename)
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.LoadFileContents() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.LoadFileContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && ("" != tt.wantErrType) && (tt.wantErrType != reflect.TypeOf(err).String()) {
				t.Errorf("HJSONConfig.LoadFileContents() error type = %v, wantErrType %v", reflect.TypeOf(err), tt.wantErrType)
				return
			}

			if false {
				if (gotCnt == nil && tt.wantCnt != nil) || (gotCnt != nil && tt.wantCnt == nil) || (0 != bytes.Compare(gotCnt, tt.wantCnt)) {
					t.Errorf("HJSONConfig.LoadFileContents() = %v, want %v", gotCnt, tt.wantCnt)
				}
			}
		})
	}
}

func TestHJSONConfig_ParseStringContents(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type args struct {
		cnt []byte
	}
	type teststruct struct {
		name        string
		fields      fields
		args        args
		wantM       map[string]interface{}
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{
		{
			name:   "Wrong broken hjson",
			fields: fields{filename: "", hjsonMap: nil},
			// broken hjson
			args:        args{cnt: []byte("{aas:}")},
			wantM:       nil,
			wantErr:     true,
			wantErrType: "",
		},
		{
			name:   "Simple json",
			fields: fields{filename: "", hjsonMap: nil},
			// broken hjson
			args:        args{cnt: []byte(`{"test field":"test text"}`)},
			wantM:       map[string]interface{}{"test field": interface{}("test text")},
			wantErr:     false,
			wantErrType: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			gotM, err := fl.ParseStringContents(tt.args.cnt)
			if false {
				fmt.Println("*********************", gotM, reflect.TypeOf(err))
			}
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.ParseStringContents() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.ParseStringContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && ("" != tt.wantErrType) && (tt.wantErrType != reflect.TypeOf(err).String()) {
				t.Errorf("HJSONConfig.ParseStringContents() error type = %v, wantErrType %v", reflect.TypeOf(err), tt.wantErrType)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("HJSONConfig.ParseStringContents() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestHJSONConfig_SetDefaultLoadSetting(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type args struct {
		sl []interface{}
	}
	type checkerfunc func(fl *HJSONConfig) bool
	type teststruct struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantErrType string
		checker     checkerfunc
	}

	tests := []teststruct{
		{
			name:   "No arguments",
			fields: fields{},
			args: args{
				sl: []interface{}{},
			},
			wantErr:     true,
			wantErrType: "*configuration.HJSONConfigError",
			checker:     nil,
		},
		{
			name:   "nil argument",
			fields: fields{},
			args: args{
				sl: []interface{}{nil},
			},
			wantErr:     true,
			wantErrType: "*configuration.HJSONConfigError",
			checker:     nil,
		},
		{
			name:   "wrong filename argument",
			fields: fields{},
			args: args{
				// compromised not existing filename
				sl: []interface{}{"test.hjson~~~~~~~"},
			},
			wantErr:     true,
			wantErrType: "*configuration.HJSONConfigError",
			checker:     nil,
		},
		{
			name:   "right filename argument",
			fields: fields{},
			args: args{
				sl: []interface{}{"test.hjson"},
			},
			wantErr:     false,
			wantErrType: "",
			checker: func(fl *HJSONConfig) bool {
				var testmap = map[string]interface{}{"CONFIG_FILE": interface{}("./configuration/config.hjson")}
				if fl.hjsonMap == nil {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): no internal object hasmap formed")
					return false
				}
				if !reflect.DeepEqual("test.hjson", fl.filename) {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): internal object filename %v differs from expected", fl.filename)
					return false
				}
				if !reflect.DeepEqual(testmap, fl.hjsonMap) {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): internal object map %v differs from expected map %v", fl.hjsonMap, testmap)
					return false
				}
				return true
			},
		},
		{
			name:   "HJSON argument - try parse HJSON bytes",
			fields: fields{},
			args: args{
				sl: []interface{}{[]byte(`{"test field":"test text"}`)},
			},
			wantErr:     false,
			wantErrType: "",
			checker: func(fl *HJSONConfig) bool {
				var testmap = map[string]interface{}{"test field": interface{}("test text")}
				if fl.hjsonMap == nil {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): no internal object hasmap formed")
					return false
				}
				if !reflect.DeepEqual("", fl.filename) {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): internal object filename %v must be an empty when try init object with []byte]", fl.filename)
					return false
				}
				if !reflect.DeepEqual(testmap, fl.hjsonMap) {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): internal object map %v differs from expected map %v", fl.hjsonMap, testmap)
					return false
				}
				return true
			},
		},
		{
			name:   "Wrong HJSON argument - try parse string",
			fields: fields{},
			args: args{
				sl: []interface{}{[]byte(`{"test field":"test text"`)},
			},
			wantErr:     true,
			wantErrType: "",
			checker:     nil,
		},
		{
			name: "Hashmap argument",
			// must setup hash map for work and not fail
			fields: fields{},
			args: args{
				sl: []interface{}{map[string]interface{}{"test field": interface{}("test text")}},
			},
			wantErr:     false,
			wantErrType: "",
			checker: func(fl *HJSONConfig) bool {
				var testmap = map[string]interface{}{"test field": interface{}("test text")}
				if fl.hjsonMap == nil {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): no internal object hasmap formed")
					return false
				}
				if !reflect.DeepEqual("", fl.filename) {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): internal object filename %v must be an empty when try init object with []byte]", fl.filename)
					return false
				}
				if !reflect.DeepEqual(testmap, fl.hjsonMap) {
					t.Errorf("HJSONConfig.LoadFileConSetDefaultLoadSettingtents(): internal object map %v differs from expected map %v", fl.hjsonMap, testmap)
					return false
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			err := fl.SetDefaultLoadSetting(tt.args.sl...)
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.ParseStringContents() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.SetDefaultLoadSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && ("" != tt.wantErrType) && (tt.wantErrType != reflect.TypeOf(err).String()) {
				t.Errorf("HJSONConfig.SetDefaultLoadSetting() error type = %v, wantErrType %v", reflect.TypeOf(err), tt.wantErrType)
				return
			}
			if nil != tt.checker {
				if !tt.checker(fl) {
					return
				}
			}
		})
	}
}

func TestHJSONConfig_CheckExternalConfig(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type teststruct struct {
		name        string
		fields      fields
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{
		{
			name: "no such file found",
			fields: fields{
				filename: "test.hjson~~~~~~~~~~~~",
				hjsonMap: nil,
			},
			wantErr: true,
			// yes here would go error from above functions
			wantErrType: "*configuration.ConfigNotConfiguredError",
		},
		{
			name: "Structure builded by hashmap",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "File exists and it is correct",
			fields: fields{
				filename: "test.hjson",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			wantErr:     false,
			wantErrType: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			err := fl.CheckExternalConfig()
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.CheckExternalConfig() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.CheckExternalConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHJSONConfig_ReloadInternalMap(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type teststruct struct {
		name        string
		fields      fields
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{
		{
			name: "no such file found",
			fields: fields{
				filename: "test.hjson~~~~~~~~~~~~",
				hjsonMap: nil,
			},
			wantErr: true,
			// yes here would go error from above functions
			wantErrType: "*configuration.ConfigNotConfiguredError",
		},
		{
			name: "Structure builded by hashmap",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "File exists and it is correct",
			fields: fields{
				filename: "test.hjson",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			wantErr:     false,
			wantErrType: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			err := fl.ReloadInternalMap()
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.ReloadInternalMap() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.ReloadInternalMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHJSONConfig_GetValue(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type args struct {
		path []string
	}
	type teststruct struct {
		name        string
		fields      fields
		args        args
		wantI       interface{}
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{
		{
			name:   "no hash map yet",
			fields: fields{filename: "", hjsonMap: nil},
			args: args{
				path: []string{"somepath there"},
			},
			wantI:       nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "empty argument",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{},
			},
			wantI:       nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "correct path and lost item",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{"nothing"},
			},
			wantI:       nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigItemNotFound",
		},
		{
			name: "normal item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"item1": interface{}("test text")},
			},
			args: args{
				path: []string{"item1"},
			},
			wantI:       interface{}("test text"),
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "recursive item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": interface{}("test text"),
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantI:       interface{}("test text"),
			wantErr:     false,
			wantErrType: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			gotI, err := fl.GetValue(tt.args.path...)
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.GetValue() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("HJSONConfig.GetValue() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestHJSONConfig_GetIntValue(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type args struct {
		path []string
	}
	type teststruct struct {
		name        string
		fields      fields
		args        args
		wantI       int
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{
		{
			name:   "no hash map yet",
			fields: fields{filename: "", hjsonMap: nil},
			args: args{
				path: []string{"somepath there"},
			},
			wantI:       0,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "empty argument",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{},
			},
			wantI:       0,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "correct path and lost item",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{"nothing"},
			},
			wantI:       0,
			wantErr:     true,
			wantErrType: "*configuration.ConfigItemNotFound",
		},
		{
			name: "normal item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"item1": interface{}(42)},
			},
			args: args{
				path: []string{"item1"},
			},
			wantI:       42,
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "recursive item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": interface{}(42),
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantI:       42,
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "type mismatch recursive  case",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": interface{}("bullshit here"),
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantI:       0,
			wantErr:     true,
			wantErrType: "*configuration.ConfigTypeMismatchError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			gotI, err := fl.GetIntValue(tt.args.path...)
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.GetIntValue() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.GetIntValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("HJSONConfig.GetIntValue() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestHJSONConfig_GetStringValue(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type args struct {
		path []string
	}
	type teststruct struct {
		name        string
		fields      fields
		args        args
		wantS       string
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{
		{
			name:   "no hash map yet",
			fields: fields{filename: "", hjsonMap: nil},
			args: args{
				path: []string{"somepath there"},
			},
			wantS:       "",
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "empty argument",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{},
			},
			wantS:       "",
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "correct path and lost item",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{"nothing"},
			},
			wantS:       "",
			wantErr:     true,
			wantErrType: "*configuration.ConfigItemNotFound",
		},
		{
			name: "normal item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"item1": interface{}("42")},
			},
			args: args{
				path: []string{"item1"},
			},
			wantS:       "42",
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "recursive item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": interface{}("42s"),
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantS:       "42s",
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "type mismatch recursive  case",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": interface{}(121212),
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantS:       "",
			wantErr:     true,
			wantErrType: "*configuration.ConfigTypeMismatchError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			gotS, err := fl.GetStringValue(tt.args.path...)
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.GetStringValue() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.GetStringValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("HJSONConfig.GetStringValue() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestHJSONConfig_GetBooleanValue(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type args struct {
		path []string
	}
	type teststruct struct {
		name        string
		fields      fields
		args        args
		wantB       bool
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{
		{
			name:   "no hash map yet",
			fields: fields{filename: "", hjsonMap: nil},
			args: args{
				path: []string{"somepath there"},
			},
			wantB:       false,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "empty argument",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{},
			},
			wantB:       false,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "correct path and lost item",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{"nothing"},
			},
			wantB:       false,
			wantErr:     true,
			wantErrType: "*configuration.ConfigItemNotFound",
		},
		{
			name: "normal item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"item1": interface{}(true)},
			},
			args: args{
				path: []string{"item1"},
			},
			wantB:       true,
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "recursive item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": interface{}(false),
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantB:       false,
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "type mismatch recursive  case",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": interface{}(121212),
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantB:       false,
			wantErr:     true,
			wantErrType: "*configuration.ConfigTypeMismatchError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			gotB, err := fl.GetBooleanValue(tt.args.path...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.GetBooleanValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.GetBooleanValue() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("HJSONConfig.GetBooleanValue() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
func TestHJSONConfig_GetSubconfig(t *testing.T) {
	type fields struct {
		filename string
		hjsonMap map[string]interface{}
	}
	type args struct {
		path []string
	}
	type teststruct struct {
		name        string
		fields      fields
		args        args
		wantC       IConfig
		wantErr     bool
		wantErrType string
	}
	tests := []teststruct{

		// TODO: Add test cases.
		{
			name:   "no hash map yet",
			fields: fields{filename: "", hjsonMap: nil},
			args: args{
				path: []string{"somepath there"},
			},
			wantC:       nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "empty argument",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{},
			},
			wantC:       nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigUsageError",
		},
		{
			name: "correct path and lost item",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			args: args{
				path: []string{"nothing"},
			},
			wantC:       nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigItemNotFound",
		},
		{
			name: "normal item extraction with type mismatch",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{"item1": interface{}("42")},
			},
			args: args{
				path: []string{"item1"},
			},
			wantC:       nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigTypeMismatchError",
		},
		{
			name: "normal item extraction with type mismatch",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{"item1": interface{}("42")},
				},
			},
			args: args{
				path: []string{"item1"},
			},
			wantC:       &HJSONConfig{filename: "", hjsonMap: map[string]interface{}{"item1": interface{}("42")}},
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "recursive item extraction",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": map[string]interface{}{"item1": interface{}("42")},
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantC:       &HJSONConfig{filename: "", hjsonMap: map[string]interface{}{"item1": interface{}("42")}},
			wantErr:     false,
			wantErrType: "",
		},
		{
			name: "type mismatch recursive  case",
			fields: fields{
				filename: "",
				hjsonMap: map[string]interface{}{
					"item1": map[string]interface{}{
						"item2": interface{}(121212),
					},
				},
			},
			args: args{
				path: []string{"item1", "item2"},
			},
			wantC:       nil,
			wantErr:     true,
			wantErrType: "*configuration.ConfigTypeMismatchError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &HJSONConfig{
				filename: tt.fields.filename,
				hjsonMap: tt.fields.hjsonMap,
			}
			gotC, err := fl.GetSubconfig(tt.args.path...)
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("HJSONConfig.GetSubconfig() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HJSONConfig.GetSubconfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("HJSONConfig.GetSubconfig() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestNewHJSONConfig(t *testing.T) {
	type args struct {
		sl []interface{}
	}
	type checkerfunc func(fl *HJSONConfig) bool
	type teststruct struct {
		name        string
		args        args
		wantFl      *HJSONConfig
		wantErr     bool
		wantErrType string
		checker     checkerfunc
	}
	tests := []teststruct{
		{
			name: "No arguments",
			args: args{
				sl: []interface{}{},
			},
			wantFl:      nil,
			wantErr:     true,
			wantErrType: "*configuration.HJSONConfigError",
			checker:     nil,
		},
		{
			name: "nil argument",
			args: args{
				sl: []interface{}{nil},
			},
			wantFl:      nil,
			wantErr:     true,
			wantErrType: "*configuration.HJSONConfigError",
			checker:     nil,
		},
		{
			name: "wrong filename argument",
			args: args{
				// compromised not existing filename
				sl: []interface{}{"test.hjson~~~~~~~"},
			},
			wantFl:      nil,
			wantErr:     true,
			wantErrType: "*configuration.HJSONConfigError",
			checker:     nil,
		},
		{
			name: "right filename argument",
			args: args{
				sl: []interface{}{"test.hjson"},
			},
			wantFl: &HJSONConfig{
				filename: "test.hjson",
				hjsonMap: map[string]interface{}{"CONFIG_FILE": interface{}("./configuration/config.hjson")},
			},
			wantErr:     false,
			wantErrType: "",
			checker:     nil,
		},
		{
			name: "HJSON argument - try parse HJSON bytes",
			args: args{
				sl: []interface{}{[]byte(`{"test field":"test text"}`)},
			},
			wantFl: &HJSONConfig{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			wantErr:     false,
			wantErrType: "",
			checker:     nil,
		},
		{
			name: "Wrong HJSON argument - try parse string",
			args: args{
				sl: []interface{}{[]byte(`{"test field":"test text"`)},
			},
			wantFl:      nil,
			wantErr:     true,
			wantErrType: "",
			checker:     nil,
		},
		{
			name: "Wrong HJSON another variant delimiters",
			args: args{
				sl: []interface{}{[]byte(`{"test field":"test text"}, `)},
			},
			wantFl:      nil,
			wantErr:     true,
			wantErrType: "",
			checker:     nil,
		},
		{
			name: "Hashmap argument",
			// must setup hash map for work and not fail
			args: args{
				sl: []interface{}{map[string]interface{}{"test field": interface{}("test text")}},
			},
			wantFl: &HJSONConfig{
				filename: "",
				hjsonMap: map[string]interface{}{"test field": interface{}("test text")},
			},
			wantErr:     false,
			wantErrType: "",
			checker:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFl, err := NewHJSONConfig(tt.args.sl...)
			if (err != nil) && (reflect.TypeOf(err).String() == "*configuration.ConfigNotImplementedError") {
				t.Errorf("NewHJSONConfig() error = %v, so this part of function is not correctly implemented yet", err)
				return
			}
			if tt.wantErr && ("" != tt.wantErrType) && (tt.wantErrType != reflect.TypeOf(err).String()) {
				t.Errorf("HJSONConfig.LoadFileContents() error type = %v, wantErrType %v", reflect.TypeOf(err), tt.wantErrType)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHJSONConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFl, tt.wantFl) {
				t.Errorf("NewHJSONConfig() = %v, want %v", gotFl, tt.wantFl)
			}
		})
	}
}
