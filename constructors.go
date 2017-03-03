package form

// Generate new CharField field with type text
func NewCharField(Name, Label string, Value *string, Validators *ValidatorsList) FormField {
	if Value == nil {
		return FormField{Name, "text", Label, "", nil, Validators}
	}
	return FormField{Name, "text", Label, *Value, nil, Validators}
}

// Generate new CharField field with type password
func NewPasswordField(Name, Label string, Validators *ValidatorsList) FormField {
	return FormField{Name, "password", Label, "", nil, Validators}
}
