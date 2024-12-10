package api

import "fmt"

type ApiError struct {
	HttpStatusCode int
	ErrorMessage   string
}

func (ae *ApiError) Error() string {
	return fmt.Sprintf("HttpStatusCode： %d, ErrorMessage： %s", ae.HttpStatusCode, ae.ErrorMessage)
}

func NewApiError(httpStatusCode int, errorMessage string) *ApiError {
	return &ApiError{
		HttpStatusCode: httpStatusCode,
		ErrorMessage:   errorMessage,
	}
}
