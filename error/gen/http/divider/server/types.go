// Code generated by goa v3.0.0, DO NOT EDIT.
//
// divider HTTP server types
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package server

import (
	divider "goa.design/examples/error/gen/divider"
	goa "goa.design/goa/v3"
)

// IntegerDivideHasRemainderResponseBody is the type of the "divider" service
// "integer_divide" endpoint HTTP response body for the "has_remainder" error.
type IntegerDivideHasRemainderResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// IntegerDivideDivByZeroResponseBody is the type of the "divider" service
// "integer_divide" endpoint HTTP response body for the "div_by_zero" error.
type IntegerDivideDivByZeroResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// IntegerDivideTimeoutResponseBody is the type of the "divider" service
// "integer_divide" endpoint HTTP response body for the "timeout" error.
type IntegerDivideTimeoutResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DivideDivByZeroResponseBody is the type of the "divider" service "divide"
// endpoint HTTP response body for the "div_by_zero" error.
type DivideDivByZeroResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DivideTimeoutResponseBody is the type of the "divider" service "divide"
// endpoint HTTP response body for the "timeout" error.
type DivideTimeoutResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewIntegerDivideHasRemainderResponseBody builds the HTTP response body from
// the result of the "integer_divide" endpoint of the "divider" service.
func NewIntegerDivideHasRemainderResponseBody(res *goa.ServiceError) *IntegerDivideHasRemainderResponseBody {
	body := &IntegerDivideHasRemainderResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewIntegerDivideDivByZeroResponseBody builds the HTTP response body from the
// result of the "integer_divide" endpoint of the "divider" service.
func NewIntegerDivideDivByZeroResponseBody(res *goa.ServiceError) *IntegerDivideDivByZeroResponseBody {
	body := &IntegerDivideDivByZeroResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewIntegerDivideTimeoutResponseBody builds the HTTP response body from the
// result of the "integer_divide" endpoint of the "divider" service.
func NewIntegerDivideTimeoutResponseBody(res *goa.ServiceError) *IntegerDivideTimeoutResponseBody {
	body := &IntegerDivideTimeoutResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDivideDivByZeroResponseBody builds the HTTP response body from the result
// of the "divide" endpoint of the "divider" service.
func NewDivideDivByZeroResponseBody(res *goa.ServiceError) *DivideDivByZeroResponseBody {
	body := &DivideDivByZeroResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDivideTimeoutResponseBody builds the HTTP response body from the result
// of the "divide" endpoint of the "divider" service.
func NewDivideTimeoutResponseBody(res *goa.ServiceError) *DivideTimeoutResponseBody {
	body := &DivideTimeoutResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewIntegerDivideIntOperands builds a divider service integer_divide endpoint
// payload.
func NewIntegerDivideIntOperands(a int, b int) *divider.IntOperands {
	return &divider.IntOperands{
		A: a,
		B: b,
	}
}

// NewDivideFloatOperands builds a divider service divide endpoint payload.
func NewDivideFloatOperands(a float64, b float64) *divider.FloatOperands {
	return &divider.FloatOperands{
		A: a,
		B: b,
	}
}
