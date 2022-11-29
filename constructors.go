package form

/* strFromPtr converts string pointer to a string */
func strFromPtr(str *string) (s string) {
	if str != nil {
		s = *str
	}
	return
}

/* Generate new CharField field with type text */
func NewCharField(Name string, Value *string) *FormField {
	return &FormField{
		Name,
		nil,
		"",
		strFromPtr(Value),
		"form-control",
		"text",
		"",
		"",
		false,
		false,
		nil,
	}
}

/* Generate new CharField field with type password */
func NewPasswordField(Name string) *FormField {
	field := NewCharField(Name, nil).SetType("password")
	return &field
}
