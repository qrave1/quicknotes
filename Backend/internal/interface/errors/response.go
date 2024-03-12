package errors

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Status  int `json:"-"`
	Message any `json:"message"`
}

func (e ResponseError) Error() string {
	res, _ := json.Marshal(e.Message)
	return string(res)
}

var (
	ForbiddenError = ResponseError{
		Status:  http.StatusForbidden,
		Message: "forbidden resource",
	}
)
