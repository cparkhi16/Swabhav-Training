package customError

type CustomError struct {
	errorMsg  string
	errorType ErrorType
}

type ErrorType string

const (
	FATAL   ErrorType = "Fatal"
	WARNING ErrorType = "Warning"
)

func New(errorMsg string, errorType ErrorType) *CustomError {
	return &CustomError{
		errorMsg:  errorMsg,
		errorType: errorType,
	}
}

func (e *CustomError) Error() string {
	return "ErrorType--> " + string(e.errorType) + "Error msg--> " + e.errorMsg
}
