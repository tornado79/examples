// Code generated by goa v2.1.1, DO NOT EDIT.
//
// secured_service HTTP server types
//
// Command:
// $ goa gen goa.design/examples/security/design -o
// $(GOPATH)/src/goa.design/examples/security

package server

import (
	securedservice "goa.design/examples/security/gen/secured_service"
)

// SigninResponseBody is the type of the "secured_service" service "signin"
// endpoint HTTP response body.
type SigninResponseBody struct {
	// JWT token
	JWT string `form:"jwt" json:"jwt" xml:"jwt"`
	// API Key
	APIKey string `form:"api_key" json:"api_key" xml:"api_key"`
	// OAuth2 token
	OauthToken string `form:"oauth_token" json:"oauth_token" xml:"oauth_token"`
}

// SigninUnauthorizedResponseBody is the type of the "secured_service" service
// "signin" endpoint HTTP response body for the "unauthorized" error.
type SigninUnauthorizedResponseBody string

// SecureInvalidScopesResponseBody is the type of the "secured_service" service
// "secure" endpoint HTTP response body for the "invalid-scopes" error.
type SecureInvalidScopesResponseBody string

// SecureUnauthorizedResponseBody is the type of the "secured_service" service
// "secure" endpoint HTTP response body for the "unauthorized" error.
type SecureUnauthorizedResponseBody string

// DoublySecureInvalidScopesResponseBody is the type of the "secured_service"
// service "doubly_secure" endpoint HTTP response body for the "invalid-scopes"
// error.
type DoublySecureInvalidScopesResponseBody string

// DoublySecureUnauthorizedResponseBody is the type of the "secured_service"
// service "doubly_secure" endpoint HTTP response body for the "unauthorized"
// error.
type DoublySecureUnauthorizedResponseBody string

// AlsoDoublySecureInvalidScopesResponseBody is the type of the
// "secured_service" service "also_doubly_secure" endpoint HTTP response body
// for the "invalid-scopes" error.
type AlsoDoublySecureInvalidScopesResponseBody string

// AlsoDoublySecureUnauthorizedResponseBody is the type of the
// "secured_service" service "also_doubly_secure" endpoint HTTP response body
// for the "unauthorized" error.
type AlsoDoublySecureUnauthorizedResponseBody string

// NewSigninResponseBody builds the HTTP response body from the result of the
// "signin" endpoint of the "secured_service" service.
func NewSigninResponseBody(res *securedservice.Creds) *SigninResponseBody {
	body := &SigninResponseBody{
		JWT:        res.JWT,
		APIKey:     res.APIKey,
		OauthToken: res.OauthToken,
	}
	return body
}

// NewSigninUnauthorizedResponseBody builds the HTTP response body from the
// result of the "signin" endpoint of the "secured_service" service.
func NewSigninUnauthorizedResponseBody(res securedservice.Unauthorized) SigninUnauthorizedResponseBody {
	body := SigninUnauthorizedResponseBody(res)
	return body
}

// NewSecureInvalidScopesResponseBody builds the HTTP response body from the
// result of the "secure" endpoint of the "secured_service" service.
func NewSecureInvalidScopesResponseBody(res securedservice.InvalidScopes) SecureInvalidScopesResponseBody {
	body := SecureInvalidScopesResponseBody(res)
	return body
}

// NewSecureUnauthorizedResponseBody builds the HTTP response body from the
// result of the "secure" endpoint of the "secured_service" service.
func NewSecureUnauthorizedResponseBody(res securedservice.Unauthorized) SecureUnauthorizedResponseBody {
	body := SecureUnauthorizedResponseBody(res)
	return body
}

// NewDoublySecureInvalidScopesResponseBody builds the HTTP response body from
// the result of the "doubly_secure" endpoint of the "secured_service" service.
func NewDoublySecureInvalidScopesResponseBody(res securedservice.InvalidScopes) DoublySecureInvalidScopesResponseBody {
	body := DoublySecureInvalidScopesResponseBody(res)
	return body
}

// NewDoublySecureUnauthorizedResponseBody builds the HTTP response body from
// the result of the "doubly_secure" endpoint of the "secured_service" service.
func NewDoublySecureUnauthorizedResponseBody(res securedservice.Unauthorized) DoublySecureUnauthorizedResponseBody {
	body := DoublySecureUnauthorizedResponseBody(res)
	return body
}

// NewAlsoDoublySecureInvalidScopesResponseBody builds the HTTP response body
// from the result of the "also_doubly_secure" endpoint of the
// "secured_service" service.
func NewAlsoDoublySecureInvalidScopesResponseBody(res securedservice.InvalidScopes) AlsoDoublySecureInvalidScopesResponseBody {
	body := AlsoDoublySecureInvalidScopesResponseBody(res)
	return body
}

// NewAlsoDoublySecureUnauthorizedResponseBody builds the HTTP response body
// from the result of the "also_doubly_secure" endpoint of the
// "secured_service" service.
func NewAlsoDoublySecureUnauthorizedResponseBody(res securedservice.Unauthorized) AlsoDoublySecureUnauthorizedResponseBody {
	body := AlsoDoublySecureUnauthorizedResponseBody(res)
	return body
}

// NewSigninPayload builds a secured_service service signin endpoint payload.
func NewSigninPayload() *securedservice.SigninPayload {
	v := &securedservice.SigninPayload{}

	return v
}

// NewSecurePayload builds a secured_service service secure endpoint payload.
func NewSecurePayload(fail *bool, token string) *securedservice.SecurePayload {
	v := &securedservice.SecurePayload{}
	v.Fail = fail
	v.Token = token

	return v
}

// NewDoublySecurePayload builds a secured_service service doubly_secure
// endpoint payload.
func NewDoublySecurePayload(key string, token string) *securedservice.DoublySecurePayload {
	v := &securedservice.DoublySecurePayload{}
	v.Key = key
	v.Token = token

	return v
}

// NewAlsoDoublySecurePayload builds a secured_service service
// also_doubly_secure endpoint payload.
func NewAlsoDoublySecurePayload(key *string, oauthToken *string, token *string) *securedservice.AlsoDoublySecurePayload {
	v := &securedservice.AlsoDoublySecurePayload{}
	v.Key = key
	v.OauthToken = oauthToken
	v.Token = token

	return v
}
