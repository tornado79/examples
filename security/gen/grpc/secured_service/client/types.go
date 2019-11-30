// Code generated by goa v2.0.8, DO NOT EDIT.
//
// secured_service gRPC client types
//
// Command:
// $ goa gen goa.design/examples/security/design -o
// $(GOPATH)/src/goa.design/examples/security

package client

import (
	secured_servicepb "goa.design/examples/security/gen/grpc/secured_service/pb"
	securedservice "goa.design/examples/security/gen/secured_service"
)

// NewSigninRequest builds the gRPC request type from the payload of the
// "signin" endpoint of the "secured_service" service.
func NewSigninRequest() *secured_servicepb.SigninRequest {
	message := &secured_servicepb.SigninRequest{}
	return message
}

// NewSigninResult builds the result type of the "signin" endpoint of the
// "secured_service" service from the gRPC response type.
func NewSigninResult(message *secured_servicepb.SigninResponse) *securedservice.Creds {
	result := &securedservice.Creds{
		JWT:        message.Jwt,
		APIKey:     message.ApiKey,
		OauthToken: message.OauthToken,
	}
	return result
}

// NewSecureRequest builds the gRPC request type from the payload of the
// "secure" endpoint of the "secured_service" service.
func NewSecureRequest(payload *securedservice.SecurePayload) *secured_servicepb.SecureRequest {
	message := &secured_servicepb.SecureRequest{}
	if payload.Fail != nil {
		message.Fail = *payload.Fail
	}
	return message
}

// NewSecureResult builds the result type of the "secure" endpoint of the
// "secured_service" service from the gRPC response type.
func NewSecureResult(message *secured_servicepb.SecureResponse) string {
	result := message.Field
	return result
}

// NewDoublySecureRequest builds the gRPC request type from the payload of the
// "doubly_secure" endpoint of the "secured_service" service.
func NewDoublySecureRequest(payload *securedservice.DoublySecurePayload) *secured_servicepb.DoublySecureRequest {
	message := &secured_servicepb.DoublySecureRequest{
		Key: payload.Key,
	}
	return message
}

// NewDoublySecureResult builds the result type of the "doubly_secure" endpoint
// of the "secured_service" service from the gRPC response type.
func NewDoublySecureResult(message *secured_servicepb.DoublySecureResponse) string {
	result := message.Field
	return result
}

// NewAlsoDoublySecureRequest builds the gRPC request type from the payload of
// the "also_doubly_secure" endpoint of the "secured_service" service.
func NewAlsoDoublySecureRequest(payload *securedservice.AlsoDoublySecurePayload) *secured_servicepb.AlsoDoublySecureRequest {
	message := &secured_servicepb.AlsoDoublySecureRequest{}
	if payload.Username != nil {
		message.Username = *payload.Username
	}
	if payload.Password != nil {
		message.Password = *payload.Password
	}
	if payload.Key != nil {
		message.Key = *payload.Key
	}
	return message
}

// NewAlsoDoublySecureResult builds the result type of the "also_doubly_secure"
// endpoint of the "secured_service" service from the gRPC response type.
func NewAlsoDoublySecureResult(message *secured_servicepb.AlsoDoublySecureResponse) string {
	result := message.Field
	return result
}
