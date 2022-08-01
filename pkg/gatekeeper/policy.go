package gatekeeper

import "fmt"

// AnonymousRole is a special role name for a role that is applied to unauthenticated requests.
const AnonymousRole = "anon"

func Permission(action string, subject string) string {
	return fmt.Sprintf("%s.%s", action, subject)
}

type Role struct {
	Name   string
	can    []string
	cannot []string
}

// Allow explicitly grants a permission to a Role
func (r *Role) Allow(permission string) {
	r.can = append(r.can, permission)
}

// Deny explicitly forbids a permission for a Role. Deny statements always take precedence over Allow statements.
func (r *Role) Deny(permission string) {
	r.cannot = append(r.can, permission)
}

// Evaluate checks a Role for the given permission. Roles default to DENY-ALL, meaning that a permission will be denied
// unless it was explicitly allowed. Allow rules are checked by logical OR. Deny rules are checked by logical AND.
//
// For example:
// User.Allow(DefinePermission('manage', 'all'))
// User.Deny(DefinePermission('delete', 'all'))
//
// In this example, User can do anything on all entities, but cannot delete any of them. When defining direct and
// inverted rules for the same pair of action and subject, the order of rule definition matters. Deny statements should
// follow Allow statements, otherwise they will be overridden.
func (r *Role) Evaluate(permission string) (allow bool) {
	for _, perm := range r.can {
		if perm == permission {
			allow = true
		}
	}

	for _, perm := range r.cannot {
		if perm == permission {
			return false
		}
	}

	return true
}
