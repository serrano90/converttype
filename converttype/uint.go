package converttype

//UInt convert uint value to uint pointer
func UInt(val uint) *uint {
	return &val
}

//UInt8 convert uint8 value to uint8 pointer
func UInt8(val uint8) *uint8 {
	return &val
}

//UInt16 convert uint16 value to uint16 pointer
func UInt16(val uint16) *uint16 {
	return &val
}

//UInt32 convert uint32 value to uint32 pointer
func UInt32(val uint32) *uint32 {
	return &val
}

//UInt64 convert uint64 value to uint64 pointer
func UInt64(val uint64) *uint64 {
	return &val
}

//UInt return value of uint pointer or 0 if pointer is nil
func UIntValue(val *uint) uint {
	if val != nil {
		return *val
	}
	return 0
}

//UInt8 return value of uint8 pointer or 0 if pointer is nil
func UInt8Value(val *uint8) uint8 {
	if val != nil {
		return *val
	}
	return 0
}

//UInt16 return value of uint16 pointer or 0 if pointer is nil
func UInt16Value(val *uint16) uint16 {
	if val != nil {
		return *val
	}
	return 0
}

//UInt32 return value of uint32 pointer or 0 if pointer is nil
func UInt32Value(val *uint32) uint32 {
	if val != nil {
		return *val
	}
	return 0
}

//UInt64 return value of uint64 pointer or 0 if pointer is nil
func UInt64Value(val *uint64) uint64 {
	if val != nil {
		return *val
	}
	return 0
}
