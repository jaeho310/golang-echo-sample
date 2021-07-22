package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
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

func ReturnApiFail(c echo.Context, httpStatus int, apiError ApiError, err error, v ...interface{}) error {
	return c.JSON(httpStatus, ApiResult{
		Success: false,
		Error: ApiError{
			Code:    apiError.Code,
			Message: fmt.Sprintf(apiError.Message, v...),
			Details: err.Error(),
		},
	})
}

func ReturnApiSuccess(c echo.Context, status int, result interface{}) error {
	return c.JSON(status, ApiResult{
		Success: true,
		Result:  result,
	})
}
