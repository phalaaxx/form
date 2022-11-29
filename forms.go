package form

import (
	"bytes"
	"context"
	"html/template"
	"strconv"
)

/* ValidatorFunc defines a function for FormField data validation */
type ValidatorFunc func(FormField, context.Context) error

/* ValidatorsList defines a list of ValidatorFunc */
type ValidatorsList []ValidatorFunc

/* A general purpose form  field struct */
type FormField struct {
	Name        string
	Error       []error
	Value       string
	Label       string
	Class       string
	Type        string
	Placeholder string
	Help        string
	Required    bool
	AutoFocus   bool
	Validators  *ValidatorsList
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

/* SetRequired marks FormField as mandatory */
func (f FormField) SetRequired() FormField {
	f.Required = true
	return f
}

/* SetAutoFocus gives focus to current FormField on page load  */
func (f FormField) SetAutoFocus() FormField {
	f.AutoFocus = true
	return f
}

/* SetType specifies input field type */
func (f FormField) SetType(t string) FormField {
	f.Type = t
	return f
}

/* SetPlaceholder specified a placeholder property in Formfield */
func (f FormField) SetPlaceholder(placeholder string) FormField {
	f.Placeholder = placeholder
	return f
}

/* SetHelp specified a help message for current Formfield */
func (f FormField) SetHelp(help string) FormField {
	f.Help = help
	return f
}

/* GetString returns FormField.Value as string */
func (f FormField) GetString() string {
	return f.Value
}

/* GetInt returns FormField.Value as int */
func (f FormField) GetInt() (v int, err error) {
	v, err = strconv.Atoi(f.Value)
	if err != nil {
		err = EInvalidIntValue
	}
	return
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

/* formFieldTemplate is a template to render FormField element in HTML format */
const formFieldTemplate = `
	{{ if .Label }}<label class="form-label" for="{{ .Name }}">{{ .Label }}</label>{{ end }}
	<input type="{{ .Type }}" id="{{ .Name }}" name="{{ .Name }}"
		{{- if .Class }} class="{{ .Class }}"{{ end }}
		{{- if .Value }} value="{{ .Value }}"{{ end }}
		{{- if .Placeholder}} placeholder="{{ .Placeholder }}"{{ end }}
		{{- if .Help }} aria-describedby="{{ .Name }}Help"{{ end }}
		{{- if .Required }} required{{ end }}
		{{- if .AutoFocus }} autofocus{{ end }}>
		{{ if .Help }}<div id="{{ .Name }}Help" class="form-text">{{ .Help }}</div>{{ end }}
	{{ if .Error }}{{ range $e := .Error }}<div class="text-danger">{{ $e }}</div>{{ end }}{{ end }}
`

/* formTemplate is compiled template to render FormField element in HTML format */
var formTemplate = template.Must(template.New("FormField").Parse(formFieldTemplate))

/* HTML renders FormField element in html format */
func (f FormField) HTML() template.HTML {
	var buffer bytes.Buffer
	if err := formTemplate.Execute(&buffer, f); err == nil {
		return template.HTML(buffer.String())
	}
	return template.HTML("")
}
