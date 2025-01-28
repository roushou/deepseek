package http_client

import "errors"

var (
	ErrInvalidFormat        = errors.New("invalid request format")
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrInsufficientBalance  = errors.New("insufficient balance")
	ErrInvalidParameters    = errors.New("invalid parameters")
	ErrRateLimitExceeded    = errors.New("rate limit exceeded")
	ErrServer               = errors.New("server error")
	ErrServiceUnavailable   = errors.New("service unavailable")
)
