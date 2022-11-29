package form

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

/* validation errors */
var (
	EInvalidInteger = errors.New("not a valid integer value")
	EInvalidFloat   = errors.New("not a valid float value")
	ERequired       = errors.New("this field is required")
)

/* ValidLettersGeneric is a validator generator for checking for valid letters in field */
func ValidLettersGeneric(Letters string, Error error) ValidatorFunc {
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

/* ValidRequired makes sure field is not empty. */
func ValidRequired(field FormField, ctx context.Context) error {
	if field.GetString() == "" {
		return ERequired
	}
	return nil
}

/* ValidLength makes sure that a string length is between specified min and max values. */
func ValidLength(min, max int) ValidatorFunc {
	var ELength = errors.New(
		fmt.Sprintf("must be a string between %d and %d characters in length", min, max))
	return func(field FormField, ctx context.Context) error {
		if len(field.GetString()) != 0 && (len(field.GetString()) < min || len(field.GetString()) > max) {
			return ELength
		}
		return nil
	}
}

/* ValidFieldIn verifies if item is within the list of items */
func ValidFieldIn(list []string) ValidatorFunc {
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

/* ValidInt returns error if field does not contain a valid integer value */
func ValidInt(field FormField, ctx context.Context) error {
	_, err := field.GetInt()
	if err != nil {
		return EInvalidInteger
	}
	return nil
}

/* ValidBetween makes sure that field is integer value within the specified range. */
func ValidBetween(min, max int) ValidatorFunc {
	var EInvalidInterval = errors.New(
		fmt.Sprintf("must be integer between %d and %d", min, max))
	return func(field FormField, ctx context.Context) error {
		value, err := field.GetInt()
		if err != nil {
			return EInvalidInteger
		}
		if value < min || value > max {
			return EInvalidInterval
		}
		return nil
	}
}

/* ValidFloat returns error if field does not contain a valid integer value */
func ValidFloat(field FormField, ctx context.Context) error {
	_, err := field.GetFloat()
	if err != nil {
		return EInvalidFloat
	}
	return nil
}

/* ValidBetweenFloat32 makes sure field is float64 value within the specified range. */
func ValidBetweenFloat(min, max float64) ValidatorFunc {
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

/* ValidFieldEqualTo is a validator that checks if two fields have the same value. */
func ValidFieldEqualTo(Other *FormField, err error) ValidatorFunc {
	return func(field FormField, ctx context.Context) error {
		if field.GetString() != Other.GetString() {
			return err
		}
		return nil
	}
}
