package errors

import (
	"net/http"

	"github.com/unrolled/render"
)

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code    *string `json:"code"`
	Message *string `json:"message,omitempty"`
}

const (
	ResponseCodeBadRequest    = "bad-request"
	ResponseCodeNotFound      = "not-found"
	ResponseCodeConflict      = "conflict"
	ResponseCodeInternalError = "internal-error"
	ResponseCodeForbidden     = "forbidden"
	ResponseCodeUnauthorized  = "unauthorized"
)

func Render(render *render.Render, w http.ResponseWriter, httpStatus int, code, message string) error {
	return render.JSON(w, httpStatus, ErrorResponse{Error: Error{
		Code:    &code,
		Message: &message,
	}})
}
