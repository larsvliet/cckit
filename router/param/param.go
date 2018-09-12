package param

import (
	"github.com/s7techlab/cckit/router"
)

// String creates middleware for converting to string chaincode method parameter
func String(name string, argPoss ...int) router.MiddlewareFunc {
	return Param(name, ``, argPoss...)
}

// Strings creates middleware for converting to string array chaincode method parameter
func Strings(name string, argPoss ...int) router.MiddlewareFunc {
	return Param(name, []string{}, argPoss...)
}

// Int creates middleware for converting to integer chaincode method parameter
func Int(name string, argPoss ...int) router.MiddlewareFunc {
	return Param(name, 1, argPoss...)
}

// Uint64 creates middleware for converting to integer chaincode method parameter
func Uint64(name string, argPoss ...int) router.MiddlewareFunc {
	var t uint64
	return Param(name, &t, argPoss...)
}

// Bool creates middleware for converting to bool chaincode method parameter
func Bool(name string, argPoss ...int) router.MiddlewareFunc {
	return Param(name, true, argPoss...)
}

// Struct creates middleware for converting to struct chaincode method parameter
func Struct(name string, target interface{}, argPoss ...int) router.MiddlewareFunc {
	return Param(name, target, argPoss...)
}

// Bytes creates middleware for converting to []byte chaincode method parameter
func Bytes(name string, argPoss ...int) router.MiddlewareFunc {
	return Param(name, []byte{}, argPoss...)
}
