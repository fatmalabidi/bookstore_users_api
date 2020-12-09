package error_handler

import "net/http"

type RestErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error_handler"`
}

func NewBadRequestError(message string) *RestErr {

	return &RestErr{
		Code:    http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}

}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusNotFound,
		Message: message,
		Error:   "not_found",
	}
}
