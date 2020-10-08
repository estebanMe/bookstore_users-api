package errors

import "net/http"

//RestErr model error common
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

//NewBadRequestError return BadRequestError with param message
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}


//NewNotFoundError return NewNotFoundError with param message
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
