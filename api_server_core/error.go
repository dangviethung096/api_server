package core

import "fmt"

type Error interface {
	Code() int32
	Message() string
	Error() string
}

type CommonError struct {
	code    int32
	message string
}

func NewError(code int32, err error) CommonError {
	return CommonError{code: code, message: err.Error()}
}

func NewErrorWithMessage(code int32, message string) CommonError {
	return CommonError{code: code, message: message}
}

func (e CommonError) Code() int32 {
	return e.code
}

func (e CommonError) Message() string {
	return e.message
}

func (e CommonError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.code, e.message)
}
