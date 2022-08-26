package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidationError returns an INVALID_ARGUMENT error plus a message.
func ValidationError(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

// NotYetImplementedError returns an UNIMPLEMNTED error plus a message.
func NotYetImplementedError(msg string) error {
	return status.Error(codes.Unimplemented, msg)
}

// UnauthenticatedError returns an UNAUTHENTICATED error.
func UnauthenticatedError() error {
	return status.Error(codes.Unauthenticated, "Not authenticated")
}
