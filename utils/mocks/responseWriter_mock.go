package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type ResponseWriterMock struct {
	// Mock fields
	mock.Mock
}

// Write provides a mock function with given fields: p
func (m *ResponseWriterMock) Write(p []byte) (int, error) {
	args := m.Called(p)
	return args.Get(0).(int), args.Error(1)
}

// Header provides a mock function with given fields:
func (m *ResponseWriterMock) Header() http.Header {
	args := m.Called()
	var req http.Header
	if args.Get(0) != nil {
		req = args.Get(0).(http.Header)
	}
	return req
}

// WriteHeader provides a mock function with given fields: code
func (m *ResponseWriterMock) WriteHeader(code int) {
	m.Called(code)
}

// NewResponseWriterMock returns a new mock instance of ResponseWriterMock
func NewResponseWriterMock() *ResponseWriterMock {
	return &ResponseWriterMock{}
}
