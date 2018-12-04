package converttype

//String convert string value to string pointer
func String(val string) *string {
	return &val
}

//StringValue return value of string pointer or 0 if pointer is nil
func StringValue(val *string) string {
	if val != nil {
		return *val
	}
	return ""
}
