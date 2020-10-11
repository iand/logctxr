package logctxr

import (
	"context"

	"github.com/go-logr/logr"
)

type contextKey struct{}

// NewLogger returns a new instance of a Logger. This variable should be set to a function that returns an implementation
// of Logger before using FromContext.
var NewLogger func() logr.Logger = func() logr.Logger {
	panic("NewLogger must be set to a function that can create new logger instances")
}

// FromContext returns a logr.Logger constructed from the context or calls NewLogger if no logger details are found in the context.
func FromContext(ctx context.Context) logr.Logger {
	if v, ok := ctx.Value(contextKey{}).(logr.Logger); ok {
		return v
	}

	return NewLogger()
}

// NewContext returns a new context that embeds the logger.
func NewContext(ctx context.Context, l logr.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, l)
}
