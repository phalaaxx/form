package form

import (
	"net/http"
	"reflect"
)

// ValidateForm parses a POST form into pre-defined struct
func ValidateForm(r *http.Request, p interface{}) error {
	// only support POST methods
	if r.Method != "POST" {
		return EInvalidMethod
	}

	// parse POST data into form
	if err := r.ParseForm(); err != nil {
		return err
	}

	var FormError error
	// Parse form data into interface
	formStruct := reflect.ValueOf(p).Elem()

	// populate FormField value
	for HttpFormField, HttpFormValue := range r.Form {
		for n := 0; n < formStruct.NumField(); n++ {
			fieldt := formStruct.Type().Field(n)
			// only proceed if field name or tag matches that of form field
			if fieldt.Name != HttpFormField && fieldt.Tag.Get("form") != HttpFormField {
				continue
			}
			// get n-th field
			fieldn := formStruct.Field(n)
			// set form data to field
			// equivalent of form.Value = HttpFormValue[0]
			fieldn.Field(1).Set(reflect.ValueOf(HttpFormValue[0]))
		}
	}

	// run form field validators
	for n := 0; n < formStruct.NumField(); n++ {
		fieldn := formStruct.Field(n)
		field := fieldn.Interface().(FormField)
		if field.Validators == nil {
			continue
		}
		for _, validator := range *field.Validators {
			if err := validator(field, r.Context()); err != nil {
				fieldn.Field(2).Set(reflect.ValueOf(err))
				FormError = err
				break
			}
		}
	}

	// return status
	return FormError
}
