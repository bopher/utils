package utils

import (
	"fmt"
	"math"
	"strconv"
)

func asBool(v interface{}) (bool, bool) {
	b, e := strconv.ParseBool(fmt.Sprint(v))
	if e != nil {
		return false, false
	} else {
		return b, true
	}
}

func asInt64(v interface{}) (int64, bool) {
	i, e := strconv.ParseInt(fmt.Sprint(v), 10, 64)
	if e != nil {
		return 0, false
	} else {
		return i, true
	}
}

func asUint64(v interface{}) (uint64, bool) {
	i, e := strconv.ParseUint(fmt.Sprint(v), 10, 64)
	if e != nil {
		return 0, false
	} else {
		return i, true
	}
}

func asFloat64(v interface{}) (float64, bool) {
	i, e := strconv.ParseFloat(fmt.Sprint(v), 64)
	if e != nil {
		return 0, false
	} else {
		return i, true
	}
}

// CastBoolE try to parse value as bool or return error
func CastBoolE(v interface{}) (bool, error) {
	if r, ok := asBool(v); ok {
		return r, nil
	}
	return false, TaggedError([]string{"CastBool"}, "failed cast %v as bool!", v)
}

// CastBool try to parse value as bool or return fallback
func CastBool(v interface{}, fallback bool) bool {
	if v, e := CastBoolE(v); e == nil {
		return v
	}
	return fallback
}

// CastIntE try to parse value as int or return error
func CastIntE(v interface{}) (int, error) {
	if r, ok := asInt64(v); ok &&
		r >= math.MinInt &&
		r <= math.MaxInt {
		return int(r), nil
	}
	return 0, TaggedError([]string{"CastInt"}, "failed cast %v as int!", v)
}

// CastInt try to parse value as int or return fallback
func CastInt(v interface{}, fallback int) int {
	if v, e := CastIntE(v); e == nil {
		return v
	}
	return fallback
}

// CastInt8E try to parse value as int8 or return error
func CastInt8E(v interface{}) (int8, error) {
	if r, ok := asInt64(v); ok &&
		r >= math.MinInt8 &&
		r <= math.MaxInt8 {
		return int8(r), nil
	}
	return 0, TaggedError([]string{"CastInt8"}, "failed cast %v as int8!", v)
}

// CastInt8 try to parse value as int8 or return fallback
func CastInt8(v interface{}, fallback int8) int8 {
	if v, e := CastInt8E(v); e == nil {
		return v
	}
	return fallback
}

// CastInt16E try to parse value as int16 or return error
func CastInt16E(v interface{}) (int16, error) {
	if r, ok := asInt64(v); ok &&
		r >= math.MinInt16 &&
		r <= math.MaxInt16 {
		return int16(r), nil
	}
	return 0, TaggedError([]string{"CastInt16"}, "failed cast %v as int16!", v)
}

// CastInt16 try to parse value as int16 or return fallback
func CastInt16(v interface{}, fallback int16) int16 {
	if v, e := CastInt16E(v); e == nil {
		return v
	}
	return fallback
}

// CastInt32E try to parse value as int32 or return error
func CastInt32E(v interface{}) (int32, error) {
	if r, ok := asInt64(v); ok &&
		r >= math.MinInt32 &&
		r <= math.MaxInt32 {
		return int32(r), nil
	}
	return 0, TaggedError([]string{"CastInt32"}, "failed cast %v as int32!", v)
}

// CastInt32 try to parse value as int32 or return fallback
func CastInt32(v interface{}, fallback int32) int32 {
	if v, e := CastInt32E(v); e == nil {
		return v
	}
	return fallback
}

// CastInt64E try to parse value as int64 or return error
func CastInt64E(v interface{}) (int64, error) {
	if r, ok := asInt64(v); ok {
		return r, nil
	}
	return 0, TaggedError([]string{"CastInt64"}, "failed cast %v as int64!", v)
}

// CastInt64 try to parse value as int64 or return fallback
func CastInt64(v interface{}, fallback int64) int64 {
	if v, e := CastInt64E(v); e == nil {
		return v
	}
	return fallback
}

// CastUIntE try to parse value as uint or return error
func CastUIntE(v interface{}) (uint, error) {
	if r, ok := asUint64(v); ok &&
		r >= 0 &&
		r <= math.MaxUint {
		return uint(r), nil
	}
	return 0, TaggedError([]string{"CastUInt"}, "failed cast %v as uint!", v)
}

// CastUInt try to parse value as uint or return fallback
func CastUInt(v interface{}, fallback uint) uint {
	if v, e := CastUIntE(v); e == nil {
		return v
	}
	return fallback
}

// CastUInt8E try to parse value as uint8 or return error
func CastUInt8E(v interface{}) (uint8, error) {
	if r, ok := asUint64(v); ok &&
		r >= 0 &&
		r <= math.MaxUint8 {
		return uint8(r), nil
	}
	return 0, TaggedError([]string{"CastUInt8"}, "failed cast %v as uint8!", v)
}

// CastUInt8 try to parse value as uint8 or return fallback
func CastUInt8(v interface{}, fallback uint8) uint8 {
	if v, e := CastUInt8E(v); e == nil {
		return v
	}
	return fallback
}

// CastUInt16E try to parse value as uint16 or return error
func CastUInt16E(v interface{}) (uint16, error) {
	if r, ok := asUint64(v); ok &&
		r >= 0 &&
		r <= math.MaxUint16 {
		return uint16(r), nil
	}
	return 0, TaggedError([]string{"CastUInt16"}, "failed cast %v as uint16!", v)
}

// CastUInt16 try to parse value as uint16 or return fallback
func CastUInt16(v interface{}, fallback uint16) uint16 {
	if v, e := CastUInt16E(v); e == nil {
		return v
	}
	return fallback
}

// CastUInt32E try to parse value as uint32 or return error
func CastUInt32E(v interface{}) (uint32, error) {
	if r, ok := asUint64(v); ok &&
		r >= 0 &&
		r <= math.MaxUint32 {
		return uint32(r), nil
	}
	return 0, TaggedError([]string{"CastUInt32"}, "failed cast %v as uint32!", v)
}

// CastUInt32 try to parse value as uint32 or return fallback
func CastUInt32(v interface{}, fallback uint32) uint32 {
	if v, e := CastUInt32E(v); e == nil {
		return v
	}
	return fallback
}

// CastUInt64E try to parse value as uint64 or return error
func CastUInt64E(v interface{}) (uint64, error) {
	if r, ok := asUint64(v); ok {
		return r, nil
	}
	return 0, TaggedError([]string{"CastUInt64"}, "failed cast %v as uint64!", v)
}

// CastUInt64 try to parse value as uint64 or return fallback
func CastUInt64(v interface{}, fallback uint64) uint64 {
	if v, e := CastUInt64E(v); e == nil {
		return v
	}
	return fallback
}

// CastFloat64E try to parse value as float64 or return error
func CastFloat64E(v interface{}) (float64, error) {
	if r, ok := asFloat64(v); ok {
		return r, nil
	}
	return 0, TaggedError([]string{"CastFloat64"}, "failed cast %v as float64!", v)
}

// CastFloat64 try to parse value as float64 or return fallback
func CastFloat64(v interface{}, fallback float64) float64 {
	if v, e := CastFloat64E(v); e == nil {
		return v
	}
	return fallback
}
