package converttype

//Int convert int value to int pointer
func Int(val int) *int {
	return &val
}

//Int8 convert int8 value to int8 pointer
func Int8(val int8) *int8 {
	return &val
}

//Int16 convert int16 value to int16 pointer
func Int16(val int16) *int16 {
	return &val
}

//Int32 convert int32 value to int32 pointer
func Int32(val int32) *int32 {
	return &val
}

//Int64 convert int64 value to int64 pointer
func Int64(val int64) *int64 {
	return &val
}

//Int return value of int pointer or 0 if pointer is nil
func IntValue(val *int) int {
	if val != nil {
		return *val
	}
	return 0
}

//Int8 return value of int8 pointer or 0 if pointer is nil
func Int8Value(val *int8) int8 {
	if val != nil {
		return *val
	}
	return 0
}

//Int16 return value of int16 pointer or 0 if pointer is nil
func Int16Value(val *int16) int16 {
	if val != nil {
		return *val
	}
	return 0
}

//Int32 return value of int32 pointer or 0 if pointer is nil
func Int32Value(val *int32) int32 {
	if val != nil {
		return *val
	}
	return 0
}

//Int64 return value of int64 pointer or 0 if pointer is nil
func Int64Value(val *int64) int64 {
	if val != nil {
		return *val
	}
	return 0
}
