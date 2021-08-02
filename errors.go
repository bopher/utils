package utils

// HasError return true if error not nil, otherwise return false
func HasError(err error) bool {
	if err == nil {
		return false
	}
	return true
}

// PanicOnError generate panic if error is not null
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// UIntOrPanic get function result (uint, error)
// if error is not null generate panic
// otherwise return result
func UIntOrPanic(res uint, err error) uint {
	if err != nil {
		panic(err)
	}
	return res
}

// UInt8OrPanic get function result (uint8, error)
// if error is not null generate panic
// otherwise return result
func UInt8OrPanic(res uint8, err error) uint8 {
	if err != nil {
		panic(err)
	}
	return res
}

// UInt16OrPanic get function result (uint16, error)
// if error is not null generate panic
// otherwise return result
func UInt16OrPanic(res uint16, err error) uint16 {
	if err != nil {
		panic(err)
	}
	return res
}

// UInt32OrPanic get function result (uint32, error)
// if error is not null generate panic
// otherwise return result
func UInt32OrPanic(res uint32, err error) uint32 {
	if err != nil {
		panic(err)
	}
	return res
}

// UInt64OrPanic get function result (uint64, error)
// if error is not null generate panic
// otherwise return result
func UInt64OrPanic(res uint64, err error) uint64 {
	if err != nil {
		panic(err)
	}
	return res
}

// IntOrPanic get function result (int, error)
// if error is not null generate panic
// otherwise return result
func IntOrPanic(res int, err error) int {
	if err != nil {
		panic(err)
	}
	return res
}

// Int8OrPanic get function result (int8, error)
// if error is not null generate panic
// otherwise return result
func Int8OrPanic(res int8, err error) int8 {
	if err != nil {
		panic(err)
	}
	return res
}

// Int16OrPanic get function result (int16, error)
// if error is not null generate panic
// otherwise return result
func Int16OrPanic(res int16, err error) int16 {
	if err != nil {
		panic(err)
	}
	return res
}

// Int32OrPanic get function result (int32, error)
// if error is not null generate panic
// otherwise return result
func Int32OrPanic(res int32, err error) int32 {
	if err != nil {
		panic(err)
	}
	return res
}

// Int64OrPanic get function result (int64, error)
// if error is not null generate panic
// otherwise return result
func Int64OrPanic(res int64, err error) int64 {
	if err != nil {
		panic(err)
	}
	return res
}

// Float32OrPanic get function result (float32, error)
// if error is not null generate panic
// otherwise return result
func Float32OrPanic(res float32, err error) float32 {
	if err != nil {
		panic(err)
	}
	return res
}

// Float64OrPanic get function result (float64, error)
// if error is not null generate panic
// otherwise return result
func Float64OrPanic(res float64, err error) float64 {
	if err != nil {
		panic(err)
	}
	return res
}

// StringOrPanic get function result (string, error)
// if error is not null generate panic
// otherwise return result
func StringOrPanic(res string, err error) string {
	if err != nil {
		panic(err)
	}
	return res
}

// BoolOrPanic get function result (bool, error)
// if error is not null generate panic
// otherwise return result
func BoolOrPanic(res bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return res
}

// InterfaceOrPanic get function result (interface{}, error)
// if error is not null generate panic
// otherwise return result
func InterfaceOrPanic(res interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return res
}
