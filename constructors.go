package form

/* Generate new CharField field with type text */
func NewCharField(Name string, Value *string, Validators *ValidatorsList) FormField {
	if Value == nil {
		return FormField{Name, "", nil, Validators}
	}
	return FormField{Name, *Value, nil, Validators}
}

/* Generate new CharField field with type password */
func NewPasswordField(Name string, Validators *ValidatorsList) FormField {
	return FormField{Name, "", nil, Validators}
}
