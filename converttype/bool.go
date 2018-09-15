package converttype

//Bool convert bool value to bool pointer
func Bool(val bool) *bool {
	return &val
}

//BoolValue return value of bool pointer or 0 if pointer is nil
func BoolValue(val *bool) bool {
	if val != nil {
		return *val
	}
	return false
}
