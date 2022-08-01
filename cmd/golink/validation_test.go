package golink

import (
	pb "github.com/nthnluu/aether/pb/out"
	"testing"
)

func TestValidateLookupRequest(t *testing.T) {
	validRequest := &pb.LookupRequest{Slug: "2d2ed2"}
	invalidRequest := &pb.LookupRequest{Slug: ""}

	err := ValidateLookupRequest(validRequest)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	err = ValidateLookupRequest(invalidRequest)
	if err == nil {
		t.Errorf("Unexpected success. Expected a validation error for %s", invalidRequest)
	}
}
