package requests

import "fmt"

type NoUrlError struct{}

func (n *NoUrlError) Error() string {
	return "no RUrl found"
}

type NoMethodError struct{}

func (n *NoMethodError) Error() string {
	return "no RMethod found"
}

type StatusCodeError struct {
	statusCode int
}

func NewStatusCodeError(statusCode int) *StatusCodeError {
	return &StatusCodeError{statusCode: statusCode}
}

func (s *StatusCodeError) Error() string {
	return fmt.Sprintf("wrong status, server status code is: %d", s.statusCode)
}
