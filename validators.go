package form

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// validation errors
var (
	EInvalidInteger = errors.New("not a valid integer value")
	EInvalidFloat   = errors.New("not a valid float value")
	ERequired       = errors.New("this field is required")
)

// validLettersGeneric is a validator generator for checking
// for valid letters in field
func validLettersGeneric(Letters string, Error error) ValidatorFunc {
	Callback := func(field FormField, ctx context.Context) error {
		for _, Rune := range field.GetString() {
			if strings.IndexRune(Letters, Rune) == -1 {
				return Error
			}
		}
		return nil
	}
	return Callback
}

// validRequired returns nil if there is text in the field.
// If the field is empty it returns error.
func validRequired(field FormField, ctx context.Context) error {
	if field.GetString() == "" {
		return ERequired
	}
	return nil
}

// validLength returns a string field validator that verifies if a string
// length is between specified min and max values.
func validLength(min, max int) ValidatorFunc {
	var ELength = errors.New(
		fmt.Sprintf("must be a string between %d and %d characters in length", min, max))
	return func(field FormField, ctx context.Context) error {
		if len(field.GetString()) != 0 && (len(field.GetString()) < min || len(field.GetString()) > max) {
			return ELength
		}
		return nil
	}
}

// validFieldIn verifies if item is within the list of items
func validFieldIn(list []string) ValidatorFunc {
	var EInvalidValue = errors.New(
		fmt.Sprintf(
			"field value must be one of: %s",
			strings.Join(list, ","),
		),
	)
	return func(field FormField, ctx context.Context) error {
		for _, item := range list {
			if item == field.GetString() {
				return nil
			}
		}
		return EInvalidValue
	}
}

// validInt returns error if field does not contain a valid integer value
func validInt(field FormField, ctx context.Context) error {
	_, err := field.GetInt()
	if err != nil {
		return EInvalidInteger
	}
	return nil
}

// validBetween is a validator generator that makes sure field is
// integer value within the specified range.
func validBetween(min, max int) ValidatorFunc {
	var EInvalidInterval = errors.New(
		fmt.Sprintf("must be integer between %d and %d", min, max))
	return func(field FormField, ctx context.Context) error {
		value, err := field.GetInt()
		if err != nil {
			return err
		}
		if value < min || value > max {
			return EInvalidInterval
		}
		return nil
	}
}

// validFloat returns error if field does not contain a valid integer value
func validFloat(field FormField, ctx context.Context) error {
	_, err := field.GetFloat()
	if err != nil {
		return EInvalidFloat
	}
	return nil
}

// validBetweenFloat32 is a v alidator generator that makes sure field is
// float64 value within the specified range.
func validBetweenFloat(min, max float64) ValidatorFunc {
	var EInvalidInterval = errors.New(
		fmt.Sprintf("must be float value between %.2f and %.2f", min, max))
	return func(field FormField, ctx context.Context) error {
		value, err := field.GetFloat()
		if err != nil {
			return err
		}
		if value < min || value > max {
			return EInvalidInterval
		}
		return nil
	}
}

// validFieldEqualTo is a general purpose validator that checks
// if two fields have the same value for verification purposes.
func validFieldEqualTo(Other *FormField, err error) ValidatorFunc {
	return func(field FormField, ctx context.Context) error {
		if field.GetString() != Other.GetString() {
			return err
		}
		return nil
	}
}
