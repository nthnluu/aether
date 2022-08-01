package golink

import "github.com/nthnluu/aether/pkg/gatekeeper"

const (
	LookupLinkPermission = "link.lookup"
	CreateLinkPermission = "link.create"
	ReadLinkPermission   = "link.read"
	UpdateLinkPermission = "link.update"
)

var AnonRole = defineAnonRole()
var UserRole = defineUserRole()

func defineAnonRole() *gatekeeper.Role {
	role := &gatekeeper.Role{}
	role.Name = gatekeeper.AnonymousRole
	role.Allow(LookupLinkPermission)
	return role
}

func defineUserRole() *gatekeeper.Role {
	role := &gatekeeper.Role{}
	role.Name = "user"
	role.Allow(LookupLinkPermission)
	return role
}
