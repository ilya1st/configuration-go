package configuration

/*
In this file we would store different error for sonfiguration implementation
*/

// ConfigNotImplementedError is for cases if comesthing is not implemented
// we implement standard error interface
type ConfigNotImplementedError struct {
	str string
}

// NewConfigNotImplementedError generates error object
func NewConfigNotImplementedError(s string) *ConfigNotImplementedError {
	return &ConfigNotImplementedError{str: s}
}

// Error is standard error interface metod
func (e *ConfigNotImplementedError) Error() string {
	return e.str
}

// ConfigNotConfiguredError is for cases if comesthing is not implemented
// we implement standard error interface
type ConfigNotConfiguredError struct {
	str string
}

// NewConfigNotConfiguredError generates error object
func NewConfigNotConfiguredError(s string) *ConfigNotConfiguredError {
	return &ConfigNotConfiguredError{str: s}
}

// Error is standard error interface metod
func (e *ConfigNotConfiguredError) Error() string {
	return e.str
}

// HJSONConfigError inform when hson error occured there
type HJSONConfigError struct {
	str string
}

// NewHJSONConfigError generates error object
func NewHJSONConfigError(s string) *HJSONConfigError {
	return &HJSONConfigError{str: s}
}

// Error is standard error interface metod
func (e *HJSONConfigError) Error() string {
	return e.str
}

// ConfigUsageError is intended for misconfiguration cases
type ConfigUsageError struct {
	str string
}

// NewConfigUsageError generates error object
func NewConfigUsageError(s string) *ConfigUsageError {
	return &ConfigUsageError{str: s}
}

// Error is standard error interface metod
func (e *ConfigUsageError) Error() string {
	return e.str
}

// ConfigItemNotFound is intended for misconfiguration cases
type ConfigItemNotFound struct {
	str string
}

// NewConfigItemNotFound generates error object
func NewConfigItemNotFound(s string) *ConfigItemNotFound {
	return &ConfigItemNotFound{str: s}
}

// Error is standard error interface metod
func (e *ConfigItemNotFound) Error() string {
	return e.str
}

// ConfigTypeMismatshErrochis intended for case we get value from object of wrong type
type ConfigTypeMismatchError struct {
	str string
}

// NewConfigTypeMismatchError generates error object
func NewConfigTypeMismatchError(s string) *ConfigTypeMismatchError {
	return &ConfigTypeMismatchError{str: s}
}

// Error is standard error interface metod
func (e *ConfigTypeMismatchError) Error() string {
	return e.str
}
