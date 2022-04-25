package milog

import (
	"context"

	"github.com/sirupsen/logrus"
)

type contextKey string

const (
	// ContextKeyReqID is the context key for RequestID
	ContextKeyRequestID     = contextKey("requestID")
	ContextKeyXForwardedFor = contextKey("xForwardedFor")
	ContextKeyClientIP      = contextKey("clientIP")

	// HTTPHeaderNameRequestID has the name of the header for request ID
	HTTPHeaderNameRequestID = "X-Request-ID"
)

// https://blog.friendsofgo.tech/posts/context-en-golang/
func getFieldsFromContext(ctx context.Context) *logrus.Entry {
	xForwardedFor, _ := ctx.Value(ContextKeyXForwardedFor).(string)
	requestID, _ := ctx.Value(ContextKeyRequestID).(string)
	clientIP, _ := ctx.Value(ContextKeyClientIP).(string)
	entry := logrus.NewEntry(logrus.New())
	if xForwardedFor != "" {
		entry = entry.Logger.WithField("x-forwarder-for", xForwardedFor)
	}
	if requestID != "" {
		entry = entry.Logger.WithField("request-id", requestID)
	}
	if clientIP != "" {
		entry = entry.Logger.WithField("client-ip", xForwardedFor)
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
