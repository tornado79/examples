// Code generated by goa v3.0.0, DO NOT EDIT.
//
// chatter HTTP client types
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package client

import (
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	goa "goa.design/goa/v3"
)

// SummaryResponseBody is the type of the "chatter" service "summary" endpoint
// HTTP response body.
type SummaryResponseBody []*ChatSummaryResponse

// SubscribeResponseBody is the type of the "chatter" service "subscribe"
// endpoint HTTP response body.
type SubscribeResponseBody struct {
	// Message sent to the server
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	Action  *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// Time at which the message was added
	AddedAt *string `form:"added_at,omitempty" json:"added_at,omitempty" xml:"added_at,omitempty"`
}

// HistoryResponseBody is the type of the "chatter" service "history" endpoint
// HTTP response body.
type HistoryResponseBody struct {
	// Message sent to the server
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt *string `form:"sent_at,omitempty" json:"sent_at,omitempty" xml:"sent_at,omitempty"`
}

// LoginUnauthorizedResponseBody is the type of the "chatter" service "login"
// endpoint HTTP response body for the "unauthorized" error.
type LoginUnauthorizedResponseBody string

// EchoerInvalidScopesResponseBody is the type of the "chatter" service
// "echoer" endpoint HTTP response body for the "invalid-scopes" error.
type EchoerInvalidScopesResponseBody string

// EchoerUnauthorizedResponseBody is the type of the "chatter" service "echoer"
// endpoint HTTP response body for the "unauthorized" error.
type EchoerUnauthorizedResponseBody string

// ListenerInvalidScopesResponseBody is the type of the "chatter" service
// "listener" endpoint HTTP response body for the "invalid-scopes" error.
type ListenerInvalidScopesResponseBody string

// ListenerUnauthorizedResponseBody is the type of the "chatter" service
// "listener" endpoint HTTP response body for the "unauthorized" error.
type ListenerUnauthorizedResponseBody string

// SummaryInvalidScopesResponseBody is the type of the "chatter" service
// "summary" endpoint HTTP response body for the "invalid-scopes" error.
type SummaryInvalidScopesResponseBody string

// SummaryUnauthorizedResponseBody is the type of the "chatter" service
// "summary" endpoint HTTP response body for the "unauthorized" error.
type SummaryUnauthorizedResponseBody string

// SubscribeInvalidScopesResponseBody is the type of the "chatter" service
// "subscribe" endpoint HTTP response body for the "invalid-scopes" error.
type SubscribeInvalidScopesResponseBody string

// SubscribeUnauthorizedResponseBody is the type of the "chatter" service
// "subscribe" endpoint HTTP response body for the "unauthorized" error.
type SubscribeUnauthorizedResponseBody string

// HistoryInvalidScopesResponseBody is the type of the "chatter" service
// "history" endpoint HTTP response body for the "invalid-scopes" error.
type HistoryInvalidScopesResponseBody string

// HistoryUnauthorizedResponseBody is the type of the "chatter" service
// "history" endpoint HTTP response body for the "unauthorized" error.
type HistoryUnauthorizedResponseBody string

// ChatSummaryResponse is used to define fields on response body types.
type ChatSummaryResponse struct {
	// Message sent to the server
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt *string `form:"sent_at,omitempty" json:"sent_at,omitempty" xml:"sent_at,omitempty"`
}

// NewLoginUnauthorized builds a chatter service login endpoint unauthorized
// error.
func NewLoginUnauthorized(body LoginUnauthorizedResponseBody) chatter.Unauthorized {
	v := chatter.Unauthorized(body)
	return v
}

// NewEchoerInvalidScopes builds a chatter service echoer endpoint
// invalid-scopes error.
func NewEchoerInvalidScopes(body EchoerInvalidScopesResponseBody) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)
	return v
}

// NewEchoerUnauthorized builds a chatter service echoer endpoint unauthorized
// error.
func NewEchoerUnauthorized(body EchoerUnauthorizedResponseBody) chatter.Unauthorized {
	v := chatter.Unauthorized(body)
	return v
}

// NewListenerInvalidScopes builds a chatter service listener endpoint
// invalid-scopes error.
func NewListenerInvalidScopes(body ListenerInvalidScopesResponseBody) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)
	return v
}

// NewListenerUnauthorized builds a chatter service listener endpoint
// unauthorized error.
func NewListenerUnauthorized(body ListenerUnauthorizedResponseBody) chatter.Unauthorized {
	v := chatter.Unauthorized(body)
	return v
}

// NewSummaryChatSummaryCollectionOK builds a "chatter" service "summary"
// endpoint result from a HTTP "OK" response.
func NewSummaryChatSummaryCollectionOK(body SummaryResponseBody) chatterviews.ChatSummaryCollectionView {
	v := make([]*chatterviews.ChatSummaryView, len(body))
	for i, val := range body {
		v[i] = &chatterviews.ChatSummaryView{
			Message: val.Message,
			Length:  val.Length,
			SentAt:  val.SentAt,
		}
	}
	return v
}

// NewSummaryInvalidScopes builds a chatter service summary endpoint
// invalid-scopes error.
func NewSummaryInvalidScopes(body SummaryInvalidScopesResponseBody) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)
	return v
}

// NewSummaryUnauthorized builds a chatter service summary endpoint
// unauthorized error.
func NewSummaryUnauthorized(body SummaryUnauthorizedResponseBody) chatter.Unauthorized {
	v := chatter.Unauthorized(body)
	return v
}

// NewSubscribeEventOK builds a "chatter" service "subscribe" endpoint result
// from a HTTP "OK" response.
func NewSubscribeEventOK(body *SubscribeResponseBody) *chatter.Event {
	v := &chatter.Event{
		Message: *body.Message,
		Action:  *body.Action,
		AddedAt: *body.AddedAt,
	}
	return v
}

// NewSubscribeInvalidScopes builds a chatter service subscribe endpoint
// invalid-scopes error.
func NewSubscribeInvalidScopes(body SubscribeInvalidScopesResponseBody) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)
	return v
}

// NewSubscribeUnauthorized builds a chatter service subscribe endpoint
// unauthorized error.
func NewSubscribeUnauthorized(body SubscribeUnauthorizedResponseBody) chatter.Unauthorized {
	v := chatter.Unauthorized(body)
	return v
}

// NewHistoryChatSummaryOK builds a "chatter" service "history" endpoint result
// from a HTTP "OK" response.
func NewHistoryChatSummaryOK(body *HistoryResponseBody) *chatterviews.ChatSummaryView {
	v := &chatterviews.ChatSummaryView{
		Message: body.Message,
		Length:  body.Length,
		SentAt:  body.SentAt,
	}
	return v
}

// NewHistoryInvalidScopes builds a chatter service history endpoint
// invalid-scopes error.
func NewHistoryInvalidScopes(body HistoryInvalidScopesResponseBody) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)
	return v
}

// NewHistoryUnauthorized builds a chatter service history endpoint
// unauthorized error.
func NewHistoryUnauthorized(body HistoryUnauthorizedResponseBody) chatter.Unauthorized {
	v := chatter.Unauthorized(body)
	return v
}

// ValidateSubscribeResponseBody runs the validations defined on
// SubscribeResponseBody
func ValidateSubscribeResponseBody(body *SubscribeResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Action == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action", "body"))
	}
	if body.AddedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("added_at", "body"))
	}
	if body.Action != nil {
		if !(*body.Action == "added") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.action", *body.Action, []interface{}{"added"}))
		}
	}
	if body.AddedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.added_at", *body.AddedAt, goa.FormatDateTime))
	}
	return
}

// ValidateChatSummaryResponse runs the validations defined on
// ChatSummaryResponse
func ValidateChatSummaryResponse(body *ChatSummaryResponse) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.SentAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sent_at", "body"))
	}
	if body.SentAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.sent_at", *body.SentAt, goa.FormatDateTime))
	}
	return
}
