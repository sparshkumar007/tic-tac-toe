package helpers

type CustomError struct {
	Code int
	Msg  string
}

func (e *CustomError) Error() string {
	return e.Msg
}

func (e *CustomError) StatusCode() int {
	return e.Code
}

func SystemError(message string) error {
	return &CustomError{
		Msg:  message,
		Code: 500,
	}
}

func BadRequest(message string) error {
	return &CustomError{
		Msg:  message,
		Code: 400,
	}
}

func UnAuthorized(message string) error {
	return &CustomError{
		Msg:  message,
		Code: 401,
	}
}
