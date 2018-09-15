package converttype

//Complex64 convert complex64 value to complex64 pointer
func Complex64(val complex64) *complex64 {
	return &val
}

//Complex128 convert complex128 value to complex128 pointer
func Complex128(val complex128) *complex128 {
	return &val
}

//Complex64 return value of complex64 pointer or 0 if pointer is nil
func Complex64Value(val *complex64) complex64 {
	if val != nil {
		return *val
	}
	return 0
}

//Complex128 return value of complex128 pointer or 0 if pointer is nil
func Complex128Value(val *complex128) complex128 {
	if val != nil {
		return *val
	}
	return 0
}
