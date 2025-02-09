package helper

import "errors"

var (
	ErrNotFound       = errors.New("data not found")
	ErrBadRequest     = errors.New("invalid request")
	ErrEmailExists    = errors.New("email already exists")
	ErrIsbnExists     = errors.New("isbn already exists")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInternalServer = errors.New("internal server error")
)

// MapErrorCode mengembalikan status HTTP yang sesuai berdasarkan jenis error.
func MapErrorCode(err error) int {
	switch err {
	case ErrNotFound:
		return 404 // Not Found
	case ErrBadRequest, ErrEmailExists, ErrUnauthorized:
		return 400 // Bad Request atau Unauthorized
	default:
		return 500 // Internal Server Error
	}
}
