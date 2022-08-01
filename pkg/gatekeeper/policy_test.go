package gatekeeper

import "testing"

const (
	ReadArticle   = "article.read"
	EditArticle   = "article.edit"
	DeleteArticle = "article.delete"
)

var (
	UserRole   = defineUserRole()
	AuthorRole = defineAuthorRole()
	AdminRole  = defineAdminRole()
)

func defineUserRole() *Role {
	role := &Role{}
	role.Allow(ReadArticle)
	return role
}

func defineAuthorRole() *Role {
	role := &Role{}
	role.Allow(ReadArticle)
	role.Allow(EditArticle)
	role.Deny(DeleteArticle)
	return role
}

func defineAdminRole() *Role {
	role := &Role{}
	role.Allow(ReadArticle)
	role.Allow(EditArticle)
	role.Allow(DeleteArticle)
	return role
}

func TestRole_Evaluate(t *testing.T) {
	if allow := UserRole.Evaluate(ReadArticle); !allow {
		t.Error("Expected user to be able to read article")
	}

	if allow := AuthorRole.Evaluate(DeleteArticle); allow {
		t.Error("Expected author to not be able to delete article")
	}

	if allow := AdminRole.Evaluate(DeleteArticle); !allow {
		t.Error("Expected admin to be able to delete article")
	}
}
