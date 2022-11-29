package form

import (
	"context"
	"strconv"
)

/* ValidatorFunc defines a function for FormField data validation */
type ValidatorFunc func(FormField, context.Context) error

/* ValidatorsList defines a list of ValidatorFunc */
type ValidatorsList []ValidatorFunc

/* A general purpose form  field struct */
type FormField struct {
	Name       string
	Label      string
	Value      string
	Class      string
	Error      error
	Validators *ValidatorsList
}

/* SetValidators configures validators list in form field */
func (f FormField) SetValidators(validators *ValidatorsList) FormField {
	f.Validators = validators
	return f
}

/* SetLabel configures form label */
func (f FormField) SetLabel(label string) FormField {
	f.Label = label
	return f
}

/* SetClass configures class name in form field */
func (f FormField) SetClass(class string) FormField {
	f.Class = class
	return f
}

/* GetString returns FormField.Value as string */
func (f FormField) GetString() string {
	return f.Value
}

/* GetInt returns FormField.Value as int */
func (f FormField) GetInt() (int, error) {
	return strconv.Atoi(f.Value)
}

/* Int converts FormField.Value to integer value and ignores errors */
func (f FormField) Int() int {
	if result, err := strconv.Atoi(f.Value); err == nil {
		return result
	}
	return 0
}

/* GetFloat returns FormField.Value as float */
func (f FormField) GetFloat() (float64, error) {
	return strconv.ParseFloat(f.Value, 64)
}

/* Float converts FormField.Value to float and ignores errors */
func (f FormField) Float() float64 {
	if result, err := strconv.ParseFloat(f.Value, 64); err == nil {
		return result
	}
	return 0.0
}

/* GetBool returns boolean value for checkbox fields */
func (f FormField) GetBool() (bool, error) {
	/* placeholder */
	return false, nil
}

/* GetChecked returns true if checkbox has been checked and its value is "on" */
func (f FormField) GetChecked() bool {
	return f.Value == "on"
}
