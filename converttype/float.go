package converttype

//Float32 convert float32 value to float32 pointer
func Float32(val float32) *float32 {
	return &val
}

//Float64 convert float64 value to float64 pointer
func Float64(val float64) *float64 {
	return &val
}

//Float32 return value of float32 pointer or 0 if pointer is nil
func Float32Value(val *float32) float32 {
	if val != nil {
		return *val
	}
	return 0
}

//Float64 return value of int64 pointer or 0 if pointer is nil
func Float64Value(val *float64) float64 {
	if val != nil {
		return *val
	}
	return 0
}
