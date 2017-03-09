package form

import (
	"context"
	"strconv"
)

// ValidatorFunc defines a function for FormField data validation
type ValidatorFunc func(FormField, context.Context) error

// ValidatorsList defines a list of ValidatorFunc
type ValidatorsList []ValidatorFunc

// A general purpose form  field struct
type FormField struct {
	Name       string
	Value      string
	Error      error
	Validators *ValidatorsList
}

// GetString returns FormField.Value as string
func (f FormField) GetString() string {
	return f.Value
}

// GetInt returns FormField.Value as int
func (f FormField) GetInt() (int, error) {
	return strconv.Atoi(f.Value)
}

// GetFloat returns FormField.Value as float
func (f FormField) GetFloat() (float64, error) {
	return strconv.ParseFloat(f.Value, 64)
}

// GetBool returns boolean value for checkbox fields
func (f FormField) GetBool() (bool, error) {
	// placeholder
	return false, nil
}

// GetChecked returns true if checkbox has been selected
// only works if checkbox value is "on" when selected
func (f FormField) GetChecked() bool {
	return f.Value == "on"
}
