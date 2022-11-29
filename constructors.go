package form

///* Generate new CharField field with type text */
//func NewCharField(Name string, Label *string, Value *string, Class *string, Validators *ValidatorsList) FormField {
//	strFromPtr := func(str *string) (s string) {
//		if str != nil {
//			s = *str
//		}
//		return
//	}
//	return FormField{Name, strFromPtr(Label), strFromPtr(Value), strFromPtr(Class), nil, Validators}
//}

func strFromPtr(str *string) (s string) {
	if str != nil {
		s = *str
	}
	return
}

/* Generate new CharField field with type text */
func NewCharField(Name string, Value *string) *FormField {
	return &FormField{Name, "", strFromPtr(Value), "", nil, nil}
}

/* Generate new CharField field with type password */
func NewPasswordField(Name string) *FormField {
	return NewCharField(Name, nil)
}
