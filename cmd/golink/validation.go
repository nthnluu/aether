package golink

import (
	pb "aether/pb/out"
	"aether/pkg/errors"
)

func ValidateLookupRequest(req *pb.LookupRequest) error {
	req.ProtoReflect().Descriptor().Fields().Get(0).Options()

	if req.GetSlug() == "" {
		return errors.ValidationError("URL slug missing")
	}

	return nil
}

func ValidateCreateLinkRequest(req *pb.CreateLinkRequest) error {
	if req.GetDestinationUrl() == "" {
		return errors.ValidationError("destination url missing")
	}

	if req.GetSlug() == "" {
		return errors.ValidationError("slug missing")
	}

	return nil
}
