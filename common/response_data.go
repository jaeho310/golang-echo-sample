package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResult struct {
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
	Error   ApiError    `json:"error"`
}

type ApiError struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

type ArrayResult struct {
	Items      interface{} `json:"items"`
	TotalCount int64       `json:"totalCount"`
}

var (
	ApiParameterError = ApiError{Code: 601, Message: "failed to parse filter parameters"}
	ApiQueryError     = ApiError{Code: 602, Message: "failed to query"}
)

func ReturnApiFail(w http.ResponseWriter, httpStatus int, apiError ApiError, err error, v ...interface{}) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(ApiResult{
		Success: false,
		Error: ApiError{
			Code:    apiError.Code,
			Message: fmt.Sprintf(apiError.Message, v...),
			Details: err.Error(),
		},
	})
}

func ReturnApiSuccess(w http.ResponseWriter, httpStatus int, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(ApiResult{
		Success: true,
		Result:  result,
	})
}
