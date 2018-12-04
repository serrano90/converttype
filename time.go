package converttype

import "time"

//Time convert time value to time pointer
func Time(val time.Time) *time.Time {
	return &val
}

//TimeValue return value of time pointer or 0 if pointer is nil
func TimeValue(val time.Time) *time.Time {
	return &val
}
