package errors

import (
	"context"
	"net/http"

	"github.com/misalud-ai/go-nurse/milog"
	"github.com/unrolled/render"
)

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code    *string `json:"code"`
	Message *string `json:"message,omitempty"`
}

func Render(ctx context.Context, render *render.Render, w http.ResponseWriter, httpStatus int, code, message string) {
	err := render.JSON(w, httpStatus, ErrorResponse{Error: Error{
		Code:    &code,
		Message: &message,
	}})

	if err != nil {
		milog.Warnf(ctx, "rendering error %w", err)
	}
}
