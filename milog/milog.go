package milog

import (
	"context"
	"runtime"

	"github.com/sirupsen/logrus"
)

type contextKey string

var LogrusLogger *logrus.Logger

const (
	// ContextKeyReqID is the context key for RequestID
	ContextKeyRequestID      = contextKey("requestID")
	ContextKeyXForwardedFor  = contextKey("xForwardedFor")
	ContextKeyClientIP       = contextKey("clientIP")
	ContextKeyAuthorizedUser = contextKey("authorizedUser")

	// HTTPHeaderNameRequestID has the name of the header for request ID
	HTTPHeaderNameRequestID = "X-Request-ID"
)

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(4, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// https://blog.friendsofgo.tech/posts/context-en-golang/
func getFieldsFromContext(ctx context.Context) *logrus.Entry {
	xForwardedFor, _ := ctx.Value(ContextKeyXForwardedFor).(string)
	requestID, _ := ctx.Value(ContextKeyRequestID).(string)
	clientIP, _ := ctx.Value(ContextKeyClientIP).(string)
	var entry *logrus.Entry
	if LogrusLogger == nil {
		entry = logrus.NewEntry(logrus.New())
	} else {
		entry = logrus.NewEntry(LogrusLogger)
	}
	if xForwardedFor != "" {
		entry = entry.WithField("x-forwarder-for", xForwardedFor)
	}
	if requestID != "" {
		entry = entry.WithField("request-id", requestID)
	}
	if clientIP != "" {
		entry = entry.WithField("client-ip", xForwardedFor)
	}
	if funcName := runFuncName(); funcName != "" {
		entry = entry.WithField("func-name", funcName)
	}
	return entry
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	getFieldsFromContext(ctx).Debugf(format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	getFieldsFromContext(ctx).Infof(format, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	getFieldsFromContext(ctx).Warnf(format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	getFieldsFromContext(ctx).Errorf(format, args...)
}

func ErrorWithError(ctx context.Context, err error, format string, args ...interface{}) {
	getFieldsFromContext(ctx).WithError(err).Errorf(format, args...)
}
func WarnWithError(ctx context.Context, err error, format string, args ...interface{}) {
	getFieldsFromContext(ctx).WithError(err).Errorf(format, args...)
}
