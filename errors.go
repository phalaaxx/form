package form

import (
	"errors"
)

// pre-defined form errors
var (
	EInvalidMethod     = errors.New("Invalid method")
	EInvalidIntValue   = errors.New("Field value must be integer.")
	EInvalidFloatValue = errors.New("Field value must be float.")
)
