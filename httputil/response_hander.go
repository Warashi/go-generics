package httputil

import (
	"context"
	"net/http"

	"github.com/Warashi/go-generics/zero"
)

type (
	handler[T any]          func(ctx context.Context, statusCode int, response *http.Response) (T, error)
	Option[T any]           func(responseHandler *ResponseHandler[T])
	HandlerInterface[T any] interface {
		Handle2xx(context.Context, int, *http.Response) (T, error)
		Handle4xx(context.Context, int, *http.Response) (T, error)
		Handle5xx(context.Context, int, *http.Response) (T, error)
		HandleOther(context.Context, int, *http.Response) (T, error)
	}
)

func noop[T any](_ context.Context, _ int, _ *http.Response) (T, error) { return zero.New[T](), nil }

func WithHandle2xx[T any](h handler[T]) Option[T] {
	return func(responseHandler *ResponseHandler[T]) {
		responseHandler.handle2xx = h
	}
}

func WithHandle4xx[T any](h handler[T]) Option[T] {
	return func(responseHandler *ResponseHandler[T]) {
		responseHandler.handle4xx = h
	}
}

func WithHandle5xx[T any](h handler[T]) Option[T] {
	return func(responseHandler *ResponseHandler[T]) {
		responseHandler.handle5xx = h
	}
}

func WithHandleOthers[T any](h handler[T]) Option[T] {
	return func(responseHandler *ResponseHandler[T]) {
		responseHandler.handleOthers = h
	}
}

func WithInterface[T any](i HandlerInterface[T]) Option[T] {
	return func(responseHandler *ResponseHandler[T]) {
		responseHandler.handle2xx = i.Handle2xx
		responseHandler.handle4xx = i.Handle4xx
		responseHandler.handle5xx = i.Handle5xx
		responseHandler.handleOthers = i.HandleOther
	}
}

func NewResponseHandler[T any](opts ...Option[T]) *ResponseHandler[T] {
	h := &ResponseHandler[T]{
		handle2xx:    noop[T],
		handle4xx:    noop[T],
		handle5xx:    noop[T],
		handleOthers: noop[T],
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

type ResponseHandler[T any] struct {
	handle2xx    handler[T]
	handle4xx    handler[T]
	handle5xx    handler[T]
	handleOthers handler[T]
}

func (h *ResponseHandler[T]) Handle(ctx context.Context, response *http.Response) (T, error) {
	switch {
	case 200 <= response.StatusCode && response.StatusCode < 300:
		return h.handle2xx(ctx, response.StatusCode, response)
	case 400 <= response.StatusCode && response.StatusCode < 500:
		return h.handle4xx(ctx, response.StatusCode, response)
	case 500 <= response.StatusCode && response.StatusCode < 600:
		return h.handle5xx(ctx, response.StatusCode, response)
	default:
		return h.handleOthers(ctx, response.StatusCode, response)
	}
}
