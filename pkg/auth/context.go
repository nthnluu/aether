package auth

import (
	"context"
)

// GetUserIdFromContext extracts the user ID of the requester from the request context.
func GetUserIdFromContext(ctx context.Context) string {
	userId := ctx.Value("user-id")
	if userId != nil {
		return userId.(string)
	}
	return ""
}

// GetUserRoleFromContext extracts the user role of the requester from the request context.
func GetUserRoleFromContext(ctx context.Context) string {
	userRole := ctx.Value("user-role")
	if userRole != nil {
		return userRole.(string)
	}
	return ""
}
